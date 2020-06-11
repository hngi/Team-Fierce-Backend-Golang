package model

//MailerService for all mailing services
type MailerService interface {
	// Common methods
	Send()
	SendWithTemplate()
	SendMultiple()
}
