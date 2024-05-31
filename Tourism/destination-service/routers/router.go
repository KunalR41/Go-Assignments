package routers

import (
	"tourism/destination-service/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/destinations", controllers.GetDestinations).Methods("GET")
	router.HandleFunc("/destinations", controllers.CreateDestination).Methods("POST")
	router.HandleFunc("/destinations/{id}", controllers.GetDestination).Methods("GET")
	router.HandleFunc("/destinations/{id}", controllers.UpdateDestination).Methods("PUT")
	router.HandleFunc("/destinations/{id}", controllers.DeleteDestination).Methods("DELETE")
	return router
}
