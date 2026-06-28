package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAsteroidHandler(w http.ResponseWriter, r *http.Request) {

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
	asteroidData := GetAsteroids(offset, limit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asteroidData)
}
