package service

import (
	smtpSender "StoriTxChallenge/internal/infrastructure/adapters/smtp"
	fileReader "StoriTxChallenge/internal/infrastructure/ports/input"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type transaction struct {
	id          int
	date        time.Time
	transaction float64
}

func BalanceGenerator(route string) {
	var monthlyBalanceList []smtpSender.MonthlyBalance
	balanceNumber := 0.0
	averageDebitNumber := 0.0
	averageCreditNumber := 0.0
	input := fileReader.ReadFile(route)
	transactionsReader := convertTransactions(input)
	monthMap := make(map[time.Month]int)

	for _, transactionReaderItem := range transactionsReader {

		month := transactionReaderItem.date.Month()
		monthMap[month]++
		balanceNumber = balanceNumber + transactionReaderItem.transaction
		if transactionReaderItem.transaction > 0 {
			averageCreditNumber = averageCreditNumber + transactionReaderItem.transaction
		} else {
			averageDebitNumber = averageDebitNumber + transactionReaderItem.transaction
		}
	}
	for month, quantity := range monthMap {
		monthlyBalanceListItem := smtpSender.MonthlyBalance{Month: month.String(), Quantity: quantity}
		monthlyBalanceList = append(monthlyBalanceList, monthlyBalanceListItem)
	}
	balance := smtpSender.Balance{
		BalanceValue:   fmt.Sprintf("%.2f", balanceNumber),
		AverageDebit:   fmt.Sprintf("%.2f", averageDebitNumber/2),
		AverageCredit:  fmt.Sprintf("%.2f", averageCreditNumber/2),
		MonthlyBalance: monthlyBalanceList,
	}
	smtpSender.SendMail(balance)
}

func convertTransactions(input []fileReader.Transaction) []transaction {
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
