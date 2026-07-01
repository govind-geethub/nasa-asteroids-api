// linke to use on the browser : http://localhost:8080/asteroids
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No .env file found in the system")
	}
	// initialising postgres connection pool first
	InitStorage()
	defer DB.Close() // safely closes database connection when project stops

	// fetch data from the NASA API
	asteroids, err := fetchAsteroid()
	if err != nil {
		log.Fatalf("failed to fetch the data \n")
	}

	// save the the data direct into the database
	for _, asteroid := range asteroids {
		CreateAsteroid(asteroid)
	}
	fmt.Printf("Processed %d asteroids into PostgreSQL database \n", len(asteroids))

	var savedAsteroids = GetAsteroids(0, 100, false, false)
	fmt.Printf("number of saved asteroids are : %d \n", len(savedAsteroids))

	// RESTful routing mappings
	http.HandleFunc("/asteroids", GetAsteroidHandler)
	http.HandleFunc("/asteroids/create", CreateAsteroidHandler)
	http.HandleFunc("/asteroids/delete", DeleteAsteroidHandler)

	// spin up the network engine
	log.Println("Server starting on Port : 8080...")
	http.ListenAndServe(":8080", nil)
}
