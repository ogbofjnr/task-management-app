package routes

import (
	"github.com/gorilla/mux"
	"github.com/ogbofjnr/maze/handlers"
)

func RegisterTaskRoutes(r *mux.Router, taskHandler *handlers.TaskHandler) {

	r.HandleFunc("/create", taskHandler.Create).Methods("POST")
	r.HandleFunc("/get", taskHandler.Get).Methods("POST")
	r.HandleFunc("/delete", taskHandler.Delete).Methods("DELETE")
	r.HandleFunc("/search", taskHandler.Search).Methods("POST")
	r.HandleFunc("/setStatus", taskHandler.SetStatus).Methods("POST")
	r.HandleFunc("/setReminder", taskHandler.SetReminder).Methods("POST")
	r.HandleFunc("/addWorklog", taskHandler.AddWorklog).Methods("POST")

}
