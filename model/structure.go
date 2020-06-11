package model

// Transport types
const (
	sendgrid = "SENDGRID"
	mailgun  = "MAILGUN"
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
	case sendgrid:
		return nil
	case mailgun:
		return nil
	}
}
