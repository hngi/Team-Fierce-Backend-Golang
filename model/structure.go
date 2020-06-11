package model

import (
	"github.com/hngi/Team-Fierce.Backend-Golang/mailgun"
	"github.com/hngi/Team-Fierce.Backend-Golang/sendgrid"
)

// Transport types
const (
	tpSendgrid = "SENDGRID"
	tpMailgun  = "MAILGUN"
)

//MailerService is an imterface
type MailerService interface {
	Send()
	SendWithTemplate()
	SendMultiple()
}

//NewMailerService returns a MailerService
//depending on what transport the user wants
func NewMailerService(transport string) *MailerService {
	switch transport {
	case tpSendgrid:
		return sendgrid.New()
	case tpMailgun:
		return mailgun.New()
	default:
		return nil
	}
}
