package smtp

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
)

import "StoriTxChallenge/internal/application/domain"

type GmailAdapter struct{}

func (ga GmailAdapter) SendMail(balance domain.Balance) {
	t, err := template.ParseFiles("./resources/mail_template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	smtpSender := os.Getenv("SMTP_SENDER")
	smtpReceiver := os.Getenv("SMTP_RECEIVER")
	smtpPassword := "vnzdbkvwjvgabopr"

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, balance)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", "oz.odi26@gmail.com")
	mail.SetHeader("To", smtpReceiver)
	mail.SetHeader("Subject", "Stori Balance")
	mail.SetHeader("Content-Type", "text/html; charset=utf-8")
	mail.SetBody("text/html", buffer.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, smtpSender, smtpPassword)
	err = dialer.DialAndSend(mail)
	if err != nil {
		log.Fatalf("Error sending email %v", err)
	}
}
