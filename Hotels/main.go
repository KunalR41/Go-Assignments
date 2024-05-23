package main

import (
	"MyHotel/Hotels/config"
	"MyHotel/Hotels/handlers"
	"MyHotel/Hotels/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	repo := &repository.HotelRepository{DB: db}
	handler := &handlers.HotelHandler{Repo: repo}

	r := mux.NewRouter()
	r.HandleFunc("/hotels", handler.GetAllHotels).Methods("GET")
	r.HandleFunc("/hotels/{id}", handler.GetHotel).Methods("GET")
	r.HandleFunc("/hotels", handler.CreateHotel).Methods("POST")
	r.HandleFunc("/hotels/{id}", handler.UpdateHotel).Methods("PUT")
	r.HandleFunc("/hotels/{id}", handler.DeleteHotel).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
