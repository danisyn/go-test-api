package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"api/handlers"
)

func run() {
	r := mux.NewRouter()
	r.HandleFunc("/api", handlers.GetTest).Methods("POST")
	r.HandleFunc("/api/config", handlers.GetConfig).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	run()
	
}

