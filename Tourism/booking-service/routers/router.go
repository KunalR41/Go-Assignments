package routers

import (
	"tourism/booking-service/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings/{id}", controllers.GetBooking).Methods("GET")
	router.HandleFunc("/bookings/{id}", controllers.UpdateBooking).Methods("PUT")
	router.HandleFunc("/bookings/{id}", controllers.DeleteBooking).Methods("DELETE")
	return router
}
