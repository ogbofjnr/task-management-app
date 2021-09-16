package routes

import (
	"github.com/gorilla/mux"
	"github.com/ogbofjnr/maze/handlers"
	"github.com/ogbofjnr/maze/pkg/logger"
)

func InitRouter(userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get", userHandler.Get).Methods("GET")
	r.Use(logger.RecoveryMiddleware)
	api := r.PathPrefix("/api/v1/").Subrouter()
	RegisterUserRoutes(api.PathPrefix("/user").Subrouter(), userHandler)
	RegisterTaskRoutes(api.PathPrefix("/task").Subrouter(), taskHandler)

	return r
}
