package main

import (
	"log"
	"net/http"
	"ocr-pdf-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.OCRHandler).Methods("POST")

	log.Println("Starting Server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
