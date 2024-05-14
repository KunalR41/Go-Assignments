package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Singleton design pattern
func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", getMsg).Methods("GET")
	r.HandleFunc("/menu", GetMenus).Methods("GET")
	r.HandleFunc("/menu/{id}", GetMenu).Methods("GET")
	r.HandleFunc("/menu", CreateMenu).Methods("POST")
	r.HandleFunc("/menu/{id}", UpdateMenu).Methods("PUT")
	r.HandleFunc("/menu/{id}", DeleteMenu).Methods("DELETE")

	fmt.Println("Hotel Server is Running on Port 8888....!!")
	log.Fatal(http.ListenAndServe(":8888", r))
}

func getMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "~~~ Welcome To INDIAN HOTEL.....!!! ~~~")
}
