package controller

import (
	"fmt"
	"net/http"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail")
}

func SendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail with template")
}
