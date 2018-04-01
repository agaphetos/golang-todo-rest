package main

import (
	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", handleRoot).Methods("GET")
	router.HandleFunc("/api/task", handleGetTasks).Methods("GET")
	router.HandleFunc("/api/task/{id}", handleGetTask).Methods("GET")
	router.HandleFunc("/api/task/", handlePostTask).Methods("POST")
	router.HandleFunc("/api/task/", handlePutTask).Methods("PUT")
	router.HandleFunc("/api/task/{id}", handleDeleteTask).Methods("DELETE")
}
