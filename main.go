package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	asteroids, err := fetchAsteroid()
	if err != nil {
		log.Fatalf("failed to fetch data")
	}

	InitStorage(asteroids)
	fmt.Printf("Fetched %d asteroids from NASA. \n", len(asteroids))

	var savedAsteroids = GetAsteroids()
	fmt.Printf("number of saved asteroids are : %d \n", len(savedAsteroids))

	// create test
	var customAsteroid Asteroid
	CreateAsteroid((customAsteroid))

	var newAsteroids = GetAsteroids()
	fmt.Printf("after new asteroid add total asteroids are : %d \n", len(newAsteroids))

	// delete test
	TargetID := newAsteroids[0].ID
	DeleteAsteroid(TargetID)

	var updatedAsteroids = GetAsteroids()
	fmt.Printf("after the 1st asteroid deletion number of asteroids are : %d", len(updatedAsteroids))

	http.HandleFunc("/asteroids", GetAsteroidHandler)
	http.ListenAndServe(":8080", nil)
}
