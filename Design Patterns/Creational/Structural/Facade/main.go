package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// Facade Design Pattern
// UserFacade hides the complexity
type UserFacade struct{}

func (uf *UserFacade) RegisterUser(username, email string) string {
	userService := &UserService{}
	userID := userService.createUser(username, email)
	notificationService := &NotificationService{}
	notificationService.sendWelcomeEmail(email)
	return fmt.Sprintf("User registered with ID: %d", userID)
}

type UserService struct{}

func (us *UserService) createUser(username, _ string) int {

	userID := rand.Intn(1000)
	fmt.Printf("User %s created with ID: %d\n", username, userID)
	return userID
}

type NotificationService struct{}

func (ns *NotificationService) sendWelcomeEmail(email string) {

	fmt.Printf("Sending welcome email to %s\n", email)
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")

	userFacade := &UserFacade{}
	result := userFacade.RegisterUser(username, email)

	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/user", registerUserHandler)
	fmt.Println("Server is running on port 8888...")
	http.ListenAndServe(":8888", nil)
}
