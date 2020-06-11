package sendgrid

import (
  "fmt"
  "log"
  "os"

  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
  "github.com/hngi/Team-Fierce.Backend-Golang/model"
)

type sendgridstruct struct {
	Sendername string
	Sendermail string
	Sub string
	Recipientname string
	Recipientmail string
	Contents string
	templatehtml string	
}

//Send method from interface
func (sg Mailer) Send(s *sendgridstruct) {
	from := mail.NewEmail(s.Sendername, s.Sendermail)
	subject := s.Sub
	to := mail.NewEmail(s.Recipientname, s.Recipientmail)
	plainTextContent := s.Contents
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
func (sg Mailer) SendWithTemplate(s *sendgridstruct) {
	from := mail.NewEmail(s.Sendermail, s.Sendermail)
	subject := s.Sub
	to := mail.NewEmail(s.Recipientname, s.Recipientmail)
	plainTextContent := s.Contents
	htmlContent := s.templatehtml
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