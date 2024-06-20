package service

import (
	fileReader "StoriTxChallenge/internal/infrastructure/adapter/filesystem"
	smtpSender "StoriTxChallenge/internal/infrastructure/adapter/smtp"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)
import "StoriTxChallenge/internal/application/domain"

type transaction struct {
	id          int
	date        time.Time
	transaction float64
}

func BalanceGenerator(route string) {
	var (
		monthlyBalanceList   []domain.MonthlyBalance
		balanceNumber        float64
		averageDebitNumber   float64
		averageCreditNumber  float64
		averageDebitCounter  float64
		averageCreditCounter float64
	)

	gmailAdapter := smtpSender.GmailAdapter{}
	fileReaderAdapter := fileReader.FileReaderAdapter{}
	monthMap := make(map[time.Month]int)

	input := fileReaderAdapter.ReadFile(route)
	transactionsReader := convertTransactions(input)

	for _, transaction := range transactionsReader {
		month := transaction.date.Month()
		monthMap[month]++
		balanceNumber += transaction.transaction

		if transaction.transaction > 0 {
			averageCreditNumber += transaction.transaction
			averageCreditCounter++
		} else {
			averageDebitNumber += transaction.transaction
			averageDebitCounter++
		}
	}

	for month, quantity := range monthMap {
		monthlyBalanceList = append(monthlyBalanceList, domain.MonthlyBalance{
			Month:    month.String(),
			Quantity: quantity,
		})
	}

	averageCreditString := "0.00"
	if averageCreditCounter > 0 {
		averageCreditString = fmt.Sprintf("%.2f", averageCreditNumber/averageCreditCounter)
	}

	averageDebitString := "0.00"
	if averageDebitCounter > 0 {
		averageDebitString = fmt.Sprintf("%.2f", averageDebitNumber/averageDebitCounter)
	}

	balance := domain.Balance{
		BalanceValue:   fmt.Sprintf("%.2f", balanceNumber),
		AverageDebit:   averageDebitString,
		AverageCredit:  averageCreditString,
		MonthlyBalance: monthlyBalanceList,
	}

	gmailAdapter.SendMail(balance)
}

func convertTransactions(input []domain.Transaction) []transaction {
	var transactionsReader []transaction
	year := time.Now().Year()
	for _, tx := range input {
		var transactionReaderItem transaction
		transactionReaderItem.id = tx.Id
		if tx.Date != "" {
			strFullDate := fmt.Sprintf("%d/%s", year, tx.Date)
			date, err := time.Parse("2006/1/2", strFullDate)
			if err != nil {
				log.Printf("Error converting date Id %d: %v", tx.Id, err)
			}
			transactionReaderItem.date = date
		}
		if tx.Transaction != "" {
			strCleanTx := strings.Replace(tx.Transaction, "+", "", -1)
			txValue, err := strconv.ParseFloat(strCleanTx, 32)
			if err != nil {
				log.Printf("Error converting transaction value Id %d: %v", tx.Id, err)
			}
			transactionReaderItem.transaction = txValue
		}
		transactionsReader = append(transactionsReader, transactionReaderItem)
	}
	return transactionsReader
}
