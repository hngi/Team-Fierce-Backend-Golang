package controller

import (
	"fmt"
	"net/http"
	"os"
)

// Please note that this env variable
//should be declared in this program's execution environment,
//not the environment you access the api with.
var transport = os.Getenv("MAIL_TP")

//SendMailHandler should send mail directly
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
	fmt.Printf("Defined transport is %s\n", transport)
	mService := NewMailerService(transport)
	if mService == nil {
		fmt.Fprintln(w, "Please define a proper TRANSPORT")
	}
}

//SendMailWithTemplateHandler should send mail with a template attached
func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
	fmt.Printf("Defined transport is %s\n", transport)
	mService := NewMailerService(transport)
	if mService == nil {
		fmt.Fprintln(w, "Please define a proper TRANSPORT")
	}
}
