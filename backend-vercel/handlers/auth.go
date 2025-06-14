package handlers

import (
	"encoding/json"
	"net/http"
	"sparepart-api/middleware"
	"sparepart-api/models"
	"sparepart-api/storage"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var creds models.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	expectedPassword, ok := storage.Users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
