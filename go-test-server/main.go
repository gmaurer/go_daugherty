package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodPost:
		name, _ := io.ReadAll(r.Body)
		msg := fmt.Sprintf("Hello %s!", name)

		w.WriteHeader(201)
		_, err = w.Write([]byte(msg))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err = w.Write([]byte("Only POST accepted"))
	}

	if err != nil {
		log.Printf("Writing the response failed: %s\n", err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.NewServeMux()

	server.HandleFunc("/post-example", postHandler)

	log.Printf("Starting http server on port [%s]", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), server)
	if err != nil {
		log.Fatalf("Server stopped with error: %s", err.Error())
	}
}
