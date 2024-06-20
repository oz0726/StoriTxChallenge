package output

import "StoriTxChallenge/internal/application/domain"

type SmtpGmailPort interface {
	SendMail(balance domain.Balance)
}
