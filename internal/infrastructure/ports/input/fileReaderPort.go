package input

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Transaction struct {
	id          int
	date        string
	transaction string
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
			date:        read[1],
			transaction: read[2],
		}
		if read[0] != "" {
			id, err := strconv.Atoi(read[0])
			if err != nil {
				log.Printf("Error converting id to int line %d: %v", lineNumber, err)
				break
			}
			transaction.id = id
		}

		transactions = append(transactions, transaction)
		lineNumber++
	}
	return transactions
}
