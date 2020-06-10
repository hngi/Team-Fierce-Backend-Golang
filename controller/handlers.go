package controller

import (
	"fmt"
	"net/http"
)

//SendMailHandler should send mail directly
func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail")
}

//SendMailWithTemplateHandler should send mail with a template attached
func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail with template")
}
