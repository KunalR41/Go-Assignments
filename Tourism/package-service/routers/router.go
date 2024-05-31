package routers

import (
	"tourism/package-service/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/packages", controllers.GetPackages).Methods("GET")
	router.HandleFunc("/packages", controllers.CreatePackage).Methods("POST")
	router.HandleFunc("/packages/{id}", controllers.GetPackage).Methods("GET")
	router.HandleFunc("/packages/{id}", controllers.UpdatePackage).Methods("PUT")
	router.HandleFunc("/packages/{id}", controllers.DeletePackage).Methods("DELETE")
	return router
}
