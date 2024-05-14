package main

import (
	"net/http"
)

type APIBuilder interface {
	Build() http.Handler
}

type SimpleAPIBuilder struct{}

func NewSimpleAPIBuilder() *SimpleAPIBuilder {
	return &SimpleAPIBuilder{}
}

func (s *SimpleAPIBuilder) Build() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/name", nameHandler)

	return mux
}
