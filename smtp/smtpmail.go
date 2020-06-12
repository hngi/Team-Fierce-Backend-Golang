package smtp

import (
	"log"
	"net/smtp"
	"os"

	"github.com/hngi/Team-Fierce.Backend-Golang/model"
)

var (
	hostname = os.Getenv("SMTP_HOST")
	port     = os.Getenv("SMTP_PORT")
	password = os.Getenv("PASSWORD")
)

// SMTP implements the mail service interface
type SMTP struct {
	mail model.Mail
}

// New returns a new SMTP interface
func New() *SMTP {
	return &SMTP{}
}

// Send email with smtp
func (s *SMTP) Send() {
	from := s.mail.Sender.Email
	to := s.mail.Recipient.Email
	subject := s.mail.Subject
	body := s.mail.Body
	msg := []byte(
		"To: " + to +
			"Subject: " + subject +
			body)
	auth := smtp.PlainAuth("", from, password, hostname)
	err := smtp.SendMail(hostname+":"+port, auth, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SMTP) SendWithTemplate() {}
func (s *SMTP) SendMultiple()     {}
