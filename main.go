package main

import (
	"fmt"
	"log"
)

func main() {

	asteroids, err := fetchAsteroid()
	if err != nil {
		log.Fatalf("failed to fetch data")
	}

	InitStorage(asteroids)
	fmt.Printf("Fetched %d asteroids from NASA. \n", len(asteroids))
}
