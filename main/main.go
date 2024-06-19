package main

import (
	smtpSender "StoriTxChallenge/internal/infrastructure/adapters/smtp"
	fileReader "StoriTxChallenge/internal/infrastructure/ports/input"
	"fmt"
)

func main() {
	fmt.Println(fileReader.ReadFile("./resources/txns.csv"))
	smtpSender.SendMail("hello world")
}
