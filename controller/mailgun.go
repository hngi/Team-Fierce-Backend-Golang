package controller

import (
	"os"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go"
)

// Mail structure to be sent
type Mail struct {
	Sender    string
	subject   string
	Body      string
	Recipient string
}

var domain string = os.Getenv("DOMAIN")
var key string = os.Getenv("MG-PRIVATE-KEY")

// SendMail takes in the mail object and sends
func SendMail(m *Mail) {
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
	resp, id, err := mg.Send(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func SendMailWithTemplate(m *Mail) {
	
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