package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Book represents a book entity
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Clonable interface
type Clonable interface {
	Clone() Clonable
}

type ConcreteBook struct {
	Book
}

func (c *ConcreteBook) Clone() Clonable {
	return &ConcreteBook{
		Book: c.Book,
	}
}

var bookPrototype *ConcreteBook

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := make([]*ConcreteBook, 0)
	for i := 0; i < 3; i++ {
		bookClone := bookPrototype.Clone().(*ConcreteBook)
		bookClone.ID = uuid.New().String()

		books = append(books, bookClone)
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	router := mux.NewRouter()

	bookPrototype = &ConcreteBook{
		Book: Book{
			Title:  "The Book",
			Author: "S.Smith",
		},
	}
	//  "/api/books"
	router.HandleFunc("/api/books", GetBooks).Methods("GET")
	//"http://localhost:8000"
	log.Fatal(http.ListenAndServe(":8000", router))
}
