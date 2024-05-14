package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Iterator Design Pattern
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection struct {
	Items []interface{} `json:"items"`
}

type CollectionIterator struct {
	collection *Collection
	index      int
}

func NewCollection(items ...interface{}) *Collection {
	return &Collection{Items: items}
}
func (c *Collection) Iterator() Iterator {
	return &CollectionIterator{collection: c, index: 0}
}

func (ci *CollectionIterator) HasNext() bool {
	return ci.index < len(ci.collection.Items)
}

func (ci *CollectionIterator) Next() interface{} {
	item := ci.collection.Items[ci.index]
	ci.index++
	return item
}

func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var items []interface{}
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := NewCollection(items...)
	response, err := json.Marshal(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func IterateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection Collection
	if err := json.NewDecoder(r.Body).Decode(&collection); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	iterator := collection.Iterator()
	var result []interface{}
	for iterator.HasNext() {
		result = append(result, iterator.Next())
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
func main() {
	http.HandleFunc("/create", CreateCollectionHandler)
	http.HandleFunc("/iterate", IterateCollectionHandler)
	//http://localhost:8080/
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
