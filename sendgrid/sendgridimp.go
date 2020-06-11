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
	from := mail.NewEmail(sg.mail.Sender.Name, sg.mail.Sender.Email)
	subject := sg.mail.Subject
	to := mail.NewEmail(sg.mail.Recipient.Name, sg.mail.Recipient.Email)
	plainTextContent := sg.mail.Body
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
	from := mail.NewEmail(sg.mail.Sender.Name, sg.mail.Sender.Email)
	subject := sg.mail.Subject
	to := mail.NewEmail(sg.mail.Recipient.Name, sg.mail.Recipient.Email)
	plainTextContent := sg.mail.Body
	htmlContent := sg.mail.HTMLBody
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

//SendMultiple sends multiple emails
func (sg *Sendgrid) SendMultiple() {}
