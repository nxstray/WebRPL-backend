package main

import (
	"log"
	"net/http"

	"backend/kost/transport"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kosts", transport.GetKostsHandler).Methods("GET")

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}