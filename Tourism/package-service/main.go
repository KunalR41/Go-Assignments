package main

import (
	"log"
	"net/http"
	"tourism/package-service/routers"
)

func main() {
	router := routers.InitRoutes()
	log.Println("Starting Package Service on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", router))
}
