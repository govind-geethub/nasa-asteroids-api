package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// enableCORS sets up the required security headers for browser-to-backend communication
func enableCORS(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// If the browser is just asking for permission configurations (Preflight OPTIONS request)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func GetAsteroidHandler(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	} // Handle preflight check

	queryParams := r.URL.Query()

	limit := 10
	offset := 0

	if limitStr := queryParams.Get("limit"); limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil && val > 0 {
			limit = val
		}
	}

	if offsetStr := queryParams.Get("offset"); offsetStr != "" {
		if val, err := strconv.Atoi(offsetStr); err == nil && val > 0 {
			offset = val
		}
	}

	hasFilter := false
	hazardousVal := false
	if hazStr := queryParams.Get("hazardous"); hazStr != "" {
		hasFilter = true

		// parse the string into "true" / "false"
		if val, err := strconv.ParseBool(hazStr); err == nil {
			hazardousVal = val
		}
	}

	asteroidData := GetAsteroids(offset, limit, hasFilter, hazardousVal)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asteroidData)
}

func CreateAsteroidHandler(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	} // Handle preflight check

	// create for post nullifies for other http requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// storing the customer's data into newAsteroid if it fails badRequest alert
	var newAsteroid Asteroid
	err := json.NewDecoder(r.Body).Decode(&newAsteroid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// data parsed success
	CreateAsteroid(newAsteroid)
	w.WriteHeader(http.StatusCreated)
}

func DeleteAsteroidHandler(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	} // Handle preflight check

	// only DELETE method will pass
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	targetID := queryParams.Get("id")

	// extracting the asteroid ID to delete
	if targetID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Missing required 'id' query parameter"}`))
		return
	}

	DeleteAsteroid(targetID)
	w.WriteHeader(http.StatusOK) // successfully deleted
}
