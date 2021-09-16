package routes

import (
	"github.com/gorilla/mux"
	"github.com/ogbofjnr/maze/handlers"
)

func RegisterUserRoutes(r *mux.Router, userHandler *handlers.UserHandler) {

	r.HandleFunc("/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/get", userHandler.Get).Methods("POST")
	r.HandleFunc("/delete", userHandler.Delete).Methods("DELETE")
	r.HandleFunc("/search", userHandler.Search).Methods("POST")

}
