package main

import (
	"log"
	"net/http"
)

// builder design pattern
func main() {
	builder := NewSimpleAPIBuilder()

	// Create a new handler using the builder
	handler := builder.Build()

	// http://localhost:8000
	log.Println("Server listening on port 8000...")
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
