package model

//MailerService for all mailing services
type MailerService interface {
	Send()
	SendWithTemplate()
	SendMultiple()
	mail *Mail	//Every mailing service holds its mail
}