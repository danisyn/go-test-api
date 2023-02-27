package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"api/handlers"
	//"api/landing"
)

func run() {
	r := mux.NewRouter()
	r.HandleFunc("/api", handlers.GetTest).Methods("POST")
	r.HandleFunc("/api/logs", handlers.Logs)
	r.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	run()
}

