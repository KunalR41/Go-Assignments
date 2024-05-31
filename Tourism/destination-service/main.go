package main

import (
	"log"
	"net/http"
	"tourism/destination-service/routers"
)

func main() {
	router := routers.InitRoutes()
	log.Println("Starting Destination Service on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
