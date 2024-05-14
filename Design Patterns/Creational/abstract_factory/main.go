package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/area", areaHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Type   string             `json:"type"`
		Params map[string]float64 `json:"params"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var factory Factory

	switch requestData.Type {
	case "rectangle":
		factory = RectangleFactory{}
	case "circle":
		factory = CircleFactory{}
	default:
		http.Error(w, "Invalid shape type", http.StatusBadRequest)
		return
	}

	shape := factory.CreateShape(requestData.Params)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shape)
}

func areaHandler(w http.ResponseWriter, r *http.Request) {
	var shape Shape

	err := json.NewDecoder(r.Body).Decode(&shape)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	area := shape.Area()

	fmt.Fprintf(w, "Area: %f", area)
}
