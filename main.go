package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Email service here")
	})

	// Spin up basic web server with localhost:{port}
	//localhost can be omitted
	infoLogger.Printf("Serving on port %s...\n", port)
	http.ListenAndServe(":"+port, router)
}
