package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hngi/Team-Fierce.Backend-Golang/controller"
	"github.com/gorilla/mux"

)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
)
var port string

func init() {
	// Get serve port from environment
	// Allows for easy port change
	port = os.Getenv("EMS_PORT")
	//Default to 9000
	if port == "" {
		port = "9000"
	}
}

func main() {
	// Register routers
	router := mux.NewRouter()
	router.HandleFunc("/v1/sendmail", controller.SendMailHandler)
	router.HandleFunc("/v1/sendmailwithtemplate", controller.SendMailWithTemplateHandler)

	// Spin up basic web server with localhost:{port}
	//localhost can be omitted
	infoLogger.Printf("Serving on port %s...\n", port)
	http.ListenAndServe(":"+port, router)
}
