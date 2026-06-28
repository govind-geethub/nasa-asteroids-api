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

	fmt.Printf("Fetched %d asteroids from NASA. \n", len(asteroids))
}
