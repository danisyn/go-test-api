package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"api/test"
	//"api/landing"
)

func run() {
	r := mux.NewRouter()
	r.HandleFunc("/api", test.GetTest).Methods("POST")
	r.HandleFunc("/api/logs", test.Logs)
	r.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	run()
}

