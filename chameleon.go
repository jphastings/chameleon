package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var serviceName = os.Getenv("SERVICE_NAME")

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Query received", r.URL.Path)
	io.WriteString(w, fmt.Sprintf("%v: %v", serviceName, r.URL.Path))
}

func main() {

	// public views
	http.HandleFunc("/", HandleIndex)
	log.Println("Starting webserver:", serviceName)
	log.Fatal(http.ListenAndServe(":80", nil))
}
