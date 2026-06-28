package main

import (
	"encoding/json"
	"net/http"
)

type Asteroid struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	AbsoluteMagnitutde float64 `json:"absolute_magnitude_h"`
	IsHazardous        bool    `json:"is_potentially_hazardous_asteroid"`
}

type NeoResponse struct {
	NearEarthObjects []Asteroid `json:"near_earth_objects"`
}

func fetchAsteroid() ([]Asteroid, error) {
	// browser API endpoint
	resp, err := http.Get("https://api.nasa.gov/neo/rest/v1/neo/browse?api_key=DEMO_KEY")
	if err != nil {
		return nil, err
	}

	// prevent network connection leaks
	defer resp.Body.Close()

	// decode the data into NeoResponse
	var result NeoResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.NearEarthObjects, nil
}
