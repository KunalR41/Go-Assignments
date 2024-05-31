package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/destinations").HandlerFunc(destinationHandler)
	router.PathPrefix("/packages").HandlerFunc(packageHandler)
	router.PathPrefix("/bookings").HandlerFunc(bookingHandler)

	log.Println("Starting API Gateway on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func destinationHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8081"+r.RequestURI, http.StatusTemporaryRedirect)
}

func packageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8082"+r.RequestURI, http.StatusTemporaryRedirect)
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8083"+r.RequestURI, http.StatusTemporaryRedirect)
}
