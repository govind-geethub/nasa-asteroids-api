// linke to use on the browser : http://localhost:8080/asteroids
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

	if len(asteroids) == 0 {
		fmt.Println("Warning: NASA returned 0 items. Injecting mock testing data...")
		asteroids = []Asteroid{
			{ID: "2001863", Name: "1863 Antinous (1948 EA)", AbsoluteMagnitudeH: 15.32, IsPotentiallyHazardousAsteroid: false},
			{ID: "2001915", Name: "1915 Quetzalcoatl (1953 EA)", AbsoluteMagnitudeH: 18.38, IsPotentiallyHazardousAsteroid: false},
			{ID: "2001917", Name: "1917 Cuyo (1968 AA)", AbsoluteMagnitudeH: 14.38, IsPotentiallyHazardousAsteroid: false},
		}
	}

	InitStorage(asteroids)
	fmt.Printf("Fetched %d asteroids from NASA. \n", len(asteroids))

	var savedAsteroids = GetAsteroids(0, 100, false, false)
	fmt.Printf("number of saved asteroids are : %d \n", len(savedAsteroids))

	var newAsteroids = GetAsteroids(0, 100, false, false)
	fmt.Printf("after new asteroid add total asteroids are : %d \n", len(newAsteroids))

	// delete test
	// TargetID := newAsteroids[0].ID
	// DeleteAsteroid(TargetID)

	// var updatedAsteroids = GetAsteroids(0, 100, false, false)
	// fmt.Printf("after the 1st asteroid deletion number of asteroids are : %d \n", len(updatedAsteroids))

	// http.HandleFunc("/asteroids", GetAsteroidHandler)
	// http.HandleFunc("/asteroids/create", CreateAsteroidHandler)

	// RESTful routing mappings
	http.HandleFunc("/asteroids", GetAsteroidHandler)
	http.HandleFunc("/asteroids/create", CreateAsteroidHandler)
	http.HandleFunc("/asteroids/delete", DeleteAsteroidHandler)

	log.Println("Server starting on Port : 8080...")
	http.ListenAndServe(":8080", nil)
}
