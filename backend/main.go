package main

import (
	"log"
	"net/http"
	"sparepart-api/handlers"
	"sparepart-api/middleware"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Izinkan asal dari frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Tangani preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Public
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected
	s := r.PathPrefix("/spareparts").Subrouter()
	s.Use(middleware.JWTMiddleware)
	s.HandleFunc("", handlers.GetSpareparts).Methods("GET")
	s.HandleFunc("/{id}", handlers.GetSparepart).Methods("GET")
	s.HandleFunc("", handlers.CreateSparepart).Methods("POST")
	s.HandleFunc("/{id}", handlers.UpdateSparepart).Methods("PUT")
	s.HandleFunc("/{id}", handlers.DeleteSparepart).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(r)))
}
