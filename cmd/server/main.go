package main

import (
	"log"
	"net/http"

	"backend/database"
	"backend/kost/transport"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()


	r := mux.NewRouter()


	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)


	r.HandleFunc("/kosts", transport.GetKostsHandler).Methods("GET")

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}


func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}