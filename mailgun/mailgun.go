package mailgun

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go"
)

// Mail structure to be sent
type Mail struct {
	Sender    string
	subject   string
	Body      string
	HTML      string
	Recipient string
}

var domain string = os.Getenv("DOMAIN")
var key string = os.Getenv("MG-PRIVATE-KEY")

// Send takes in the mail object and sends
func (mg Mailer) Send(m *Mail) (string, error) {
	mg := mailgun.NewMailgun(domain, key)

	sender := m.Sender
	subject := m.subject
	body := m.Body
	recipient := m.Recipient

	ok := vaildate(key, recipient)

	if !ok {
		fmt.Print("recipient email is invalid")
	}
	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, id, err := mg.Send(message)

	return id, err
}

//SendWithTemplate sends email with a space for a HTML input
func (mg Mailer) SendWithTemplate(m *Mail) (string, error) {
	mg := mailgun.NewMailgun(domain, key)

	sender := m.Sender
	subject := m.subject
	body := ""
	recipient := m.Recipient

	message := mg.NewMessage(sender, subject, body, recipient)
	message.SetHtml(m.HTML)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(message)

	return id, err
}
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
