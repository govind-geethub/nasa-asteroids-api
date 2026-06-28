package main

import (
	"encoding/json"
	"net/http"
)

func GetAsteroidHandler(w http.ResponseWriter, r *http.Request) {
	asteroidData := GetAsteroids()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asteroidData)
}
