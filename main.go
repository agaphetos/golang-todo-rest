package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	registerRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
