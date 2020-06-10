package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	fmt.Printf("Serving on port %s...", port)
	http.ListenAndServe(":"+port, router)
}
