package main

import (
	"fmt"
	"net/http"
)

func sendMailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail")
}

func sendMailWithTemplateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sending mail with template")
}
