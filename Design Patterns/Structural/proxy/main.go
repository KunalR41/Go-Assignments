package main

import (
	"fmt"
	"net/http"
)

// proxy design pattern
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Application Status Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created Sucessfully"
	}
	return 404, "Not Found...!!"
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	method := r.Method

	allowed := n.checkRateLimiting(url)
	if !allowed {
		http.Error(w, "Not Allowed...!!", http.StatusForbidden)
		return
	}

	statusCode, message := n.application.handleRequest(url, method)
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s\n", message)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}

func main() {
	nginxServer := newNginxServer()

	http.HandleFunc("/", nginxServer.handleRequest)

	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
