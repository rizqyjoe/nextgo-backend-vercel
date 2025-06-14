package main

import (
	"net/http"
	"sparepart-api/handlers"
	"sparepart-api/middleware"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/api/login", handlers.Login).Methods("POST")

	s := rtr.PathPrefix("/api/spareparts").Subrouter()
	s.Use(middleware.JWTMiddleware)
	s.HandleFunc("", handlers.GetSpareparts).Methods("GET")
	s.HandleFunc("/{id}", handlers.GetSparepart).Methods("GET")
	s.HandleFunc("", handlers.CreateSparepart).Methods("POST")
	s.HandleFunc("/{id}", handlers.UpdateSparepart).Methods("PUT")
	s.HandleFunc("/{id}", handlers.DeleteSparepart).Methods("DELETE")

	rtr.Use(enableCORS)

	rtr.ServeHTTP(w, r)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
