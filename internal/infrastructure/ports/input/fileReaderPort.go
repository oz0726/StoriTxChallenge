package input

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Transaction struct {
	Id          int
	Date        string
	Transaction string
}

func ReadFile(route string) []Transaction {
	var transactions []Transaction
	file, err := os.Open(route)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = 3

	_, err = r.Read()
	if err != nil {
		log.Printf("Error reading heading: %v", err)
		return nil
	}

	lineNumber := 1

	for {
		read, err := r.Read()
		if err == io.EOF {
			log.Println("End of file")
			break
		} else if err != nil {
			log.Printf("Error reading line %d: %v", lineNumber, err)
			break
		}
		transaction := Transaction{
			Date:        read[1],
			Transaction: read[2],
		}
		if read[0] != "" {
			id, err := strconv.Atoi(read[0])
			if err != nil {
				log.Printf("Error converting id to int line %d: %v", lineNumber, err)
				break
			}
			transaction.Id = id
		}

		transactions = append(transactions, transaction)
		lineNumber++
	}
	return transactions
}
