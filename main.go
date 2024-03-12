package main

import (
	//	"blogs/kube-jobs/handlers"
	"blogs/kube-jobs/handlers"
	"log"
	"net/http"
	// "log"
	// "net/http"
)

func main() {
	handlers.RegisterHandlers()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
