package filesystem

import (
	"encoding/csv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"os"
	"strconv"
)

import "StoriTxChallenge/internal/application/domain"

type FileReaderAdapter struct{}

func (a FileReaderAdapter) ReadFile(route string) []domain.Transaction {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	bucket := os.Getenv("AWS_BUCKET")
	downloader := s3manager.NewDownloader(sess)

	file, err := os.CreateTemp("", "local_*.csv")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
		return nil
	}
	defer os.Remove(file.Name())

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(route),
	}

	_, err = downloader.Download(file, params)
	if err != nil {
		log.Fatalf("Error downloading file: %v", err)
		return nil
	}

	r := csv.NewReader(file)
	r.FieldsPerRecord = 3

	_, err = r.Read()
	if err != nil {
		log.Fatalf("Error reading heading: %v", err)
		return nil
	}

	var transactions []domain.Transaction
	lineNumber := 1

	for {
		read, err := r.Read()
		if err == io.EOF {
			log.Println("End of file")
			break
		} else if err != nil {
			log.Fatalf("Error reading line %d: %v", lineNumber, err)
			return nil
		}

		transaction := domain.Transaction{
			Date:        read[1],
			Transaction: read[2],
		}

		if read[0] != "" {
			id, err := strconv.Atoi(read[0])
			if err != nil {
				log.Fatalf("Error converting id to int line %d: %v", lineNumber, err)
				return nil
			}
			transaction.Id = id
		}

		transactions = append(transactions, transaction)
		lineNumber++
	}

	return transactions
}
