package router

import (
	"commerce/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes(userController controllers.UserController) *mux.Router {
	router := mux.NewRouter()

	// User management routes
	router.HandleFunc("/user/create", userController.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", userController.GetUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", userController.DeleteUser).Methods("DELETE")

	return router
}
