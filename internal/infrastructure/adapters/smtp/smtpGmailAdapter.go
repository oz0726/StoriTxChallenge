package smtp

import (
	"bytes"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
)

type Balance struct {
	AverageDebit   string
	AverageCredit  string
	BalanceValue   string
	MonthlyBalance []MonthlyBalance
}

type MonthlyBalance struct {
	Month    string
	Quantity int
}

func SendMail(balance Balance) {
	t, err := template.ParseFiles("./resources/mail_template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	smtpSender := os.Getenv("SMTP_SENDER")
	smtpReceiver := os.Getenv("SMTP_RECEIVER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, balance)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", "oz.odi26@gmail.com")
	mail.SetHeader("To", smtpReceiver)
	mail.SetHeader("Subject", "Hello World Gmail")
	mail.SetHeader("Content-Type", "text/html; charset=utf-8")
	mail.SetBody("text/html", buffer.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, smtpSender, smtpPassword)
	err = dialer.DialAndSend(mail)
	if err != nil {
		log.Fatalf("Error sending email %v", err)
	}
}
