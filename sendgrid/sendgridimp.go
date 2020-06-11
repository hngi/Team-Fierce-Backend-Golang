package sendgrid

import (
	"fmt"
	"log"
	"os"

	"github.com/hngi/Team-Fierce.Backend-Golang/model"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Sendgrid implements the MailService interface
type Sendgrid struct {
	mail model.Mail
}

//New return a new Sendgrid instance
func New() *Sendgrid {
	return &Sendgrid{}
}

//Send method from interface
func (sg *Sendgrid) Send() {
	from := mail.NewEmail(sg.mail.sender.name, sg.mail.sender.email)
	subject := sg.mail.subject
	to := mail.NewEmail(sg.mail.recipient.name, sg.mail.recipient.email)
	plainTextContent := sg.body
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//SendWithTemplate method
func (sg *Sendgrid) SendWithTemplate() {
	from := mail.NewEmail(sg.mail.sender.name, sg.mail.sender.email)
	subject := sg.mail.subject
	to := mail.NewEmail(sg.mail.recipient.name, sg.mail.recipient.email)
	plainTextContent := s.Contents
	htmlContent := sg.mail.htmlBody
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func (sg *Sendgrid) SendMultiple() {}
