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
func (mg *Mailgun) Send() {
	mgClient := mailgun.NewMailgun(domain, key)

	sender := mg.mail.Sender.Email
	subject := mg.mail.Subject
	body := mg.mail.Body
	recipient := mg.mail.Recipient.Email

	ok := vaildate(key, recipient)

	if !ok {
		fmt.Print("recipient email is invalid")
	}
	// The message object allows you to add attachments and Bcc recipients
	message := mgClient.NewMessage(sender, subject, body, recipient)

	// Send the message with a 10 second timeout
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//TODO: Return error if any
	mgClient.Send(message)
}

//SendWithTemplate sends email with a space for a HTML input
func (mg *Mailgun) SendWithTemplate() {
	mgClient := mailgun.NewMailgun(domain, key)

	sender := mg.mail.Sender.Email
	subject := mg.mail.Subject
	body := ""
	recipient := mg.mail.Recipient.Email

	message := mgClient.NewMessage(sender, subject, body, recipient)
	message.SetHtml(mg.mail.HTMLBody)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//TODO: Return error if any
	mgClient.Send(message)
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
