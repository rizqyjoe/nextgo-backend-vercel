package handlers

import (
	"encoding/json"
	"net/http"
	"sparepart-api/models"
	"sparepart-api/storage"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSpareparts(w http.ResponseWriter, r *http.Request) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	json.NewEncoder(w).Encode(storage.Spareparts)
}

func GetSparepart(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	for _, sp := range storage.Spareparts {
		if sp.ID == id {
			json.NewEncoder(w).Encode(sp)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateSparepart(w http.ResponseWriter, r *http.Request) {
	var sp models.Sparepart
	if err := json.NewDecoder(r.Body).Decode(&sp); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	storage.LastID++
	sp.ID = storage.LastID
	storage.Spareparts = append(storage.Spareparts, sp)

	json.NewEncoder(w).Encode(sp)
}

func UpdateSparepart(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var updated models.Sparepart
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	for i, sp := range storage.Spareparts {
		if sp.ID == id {
			updated.ID = id
			storage.Spareparts[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteSparepart(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	for i, sp := range storage.Spareparts {
		if sp.ID == id {
			storage.Spareparts = append(storage.Spareparts[:i], storage.Spareparts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
