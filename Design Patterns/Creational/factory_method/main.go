package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Factory Method
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64 `json:"radius"`
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type ShapeFactory struct{}

func (f ShapeFactory) CreateShape(shapeType string, params map[string]float64) Shape {
	switch shapeType {
	case "rectangle":
		return Rectangle{
			Width:  params["width"],
			Height: params["height"],
		}
	case "circle":
		return Circle{
			Radius: params["radius"],
		}
	default:
		return nil
	}
}

func main() {
	shapeFactory := ShapeFactory{}
	// "/create"
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		var requestData struct {
			Type   string             `json:"type"`
			Params map[string]float64 `json:"params"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		shape := shapeFactory.CreateShape(requestData.Type, requestData.Params)
		if shape == nil {
			http.Error(w, "Invalid shape type", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shape)
	})
	//"/area"
	http.HandleFunc("/area", func(w http.ResponseWriter, r *http.Request) {
		var shape Shape

		err := json.NewDecoder(r.Body).Decode(&shape)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		area := shape.Area()

		fmt.Fprintf(w, "Area: %f", area)
	})

	//http://localhost:8080
	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
	fmt.Println("Server listening on port:", port)
}
