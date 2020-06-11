package controller

import (
	"fmt"
	"net/http"
	"os"
)

var transport = os.Getenv("TRANSPORT")

//SendMailHandler should send mail directly
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail")
	// Get mail service for transport
	mService := NewMailerService(transport)
	if mService == nil {
		fmt.Fprintln("Please define a TRANSPORT")
	}
}

//SendMailWithTemplateHandler should send mail with a template attached
func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail with template")
	// Get mail service for transport
	mService := NewMailerService(transport)
	if mService == nil {
		fmt.Fprintln("Please define a TRANSPORT")
	}
}
