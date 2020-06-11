package mailgun

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hngi/Team-Fierce.Backend-Golang/model"
	"github.com/mailgun/mailgun-go"
)

var (
	domain string = os.Getenv("DOMAIN")
	key    string = os.Getenv("MG-PRIVATE-KEY")
)

//Mailgun implements the MailService interface
type Mailgun struct {
	mail model.Mail
}

//New returns a new Mailgun instance
func New() *Mailgun {
	return &Mailgun{}
}

// Send takes in the mail object and sends
func (mg *Mailgun) Send() (string, error) {
	mgClient := mailgun.NewMailgun(domain, key)

	sender := mg.mail.sender.email
	subject := mg.mail.subject
	body := mg.mail.body
	recipient := mg.mail.recipient.email

	ok := vaildate(key, recipient)

	if !ok {
		fmt.Print("recipient email is invalid")
	}
	// The message object allows you to add attachments and Bcc recipients
	message := mgClient.NewMessage(sender, subject, body, recipient)

	// Send the message with a 10 second timeout
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mgClient.Send(message)

	return id, err
}

//SendWithTemplate sends email with a space for a HTML input
func (mg *Mailgun) SendWithTemplate() (string, error) {
	mgClient := mailgun.NewMailgun(domain, key)

	sender := mg.mail.sender
	subject := mg.mail.subject
	body := ""
	recipient := mg.mail.recipient

	message := mgClient.NewMessage(sender, subject, body, recipient)
	message.SetHtml(mg.mail.htmlBody)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mgClient.Send(message)

	return id, err
}

//SendMultiple sends multiple emails
func (mg *Mailgun) SendMultiple() {}

func vaildate(key, email string) bool {
	v := mailgun.NewEmailValidator(key)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := v.ValidateEmail(email, false)
	if err != nil {
		panic(err)
	}
	return true
}
