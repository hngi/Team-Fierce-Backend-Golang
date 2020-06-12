package controller

import (
	"log"
	"net/http"
	"os"
)

var infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

// Please note that this env variable
//should be declared in this program's execution environment,
//not the environment you access the api from.
var transport = os.Getenv("MAIL_TP")

//SendMailHandler should send mail directly
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
	NewMailerService(transport)
}

//SendMailWithTemplateHandler should send mail with a template attached
func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	// Get mail service for defined transport
}
