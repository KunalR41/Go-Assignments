package main

import (
	"log"
	"net/http"
	"tourism/booking-service/routers"
)

func main() {
	router := routers.InitRoutes()
	log.Println("Starting Booking Service on port 8083...")
	log.Fatal(http.ListenAndServe(":8083", router))
}
