package model

//MailerService is an imterface
type MailerService interface {
	Send()
	SendWithTemplate()
	SendMultiple()
}

//Mailer struct
type Mailer struct {
	client MailerService
}

//NewMailer function
func NewMailer(client MailerService) *Mailer {

}
