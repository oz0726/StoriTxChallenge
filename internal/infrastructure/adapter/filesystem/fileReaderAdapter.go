package filesystem

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

import "StoriTxChallenge/internal/application/domain"

type FileReaderAdapter struct{}

func (a FileReaderAdapter) ReadFile(route string) []domain.Transaction {
	file, err := os.Open(route)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return nil
	}
	defer file.Close()

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
