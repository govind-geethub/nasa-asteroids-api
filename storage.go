package main

import (
	"fmt"
)

var AsteroidDB []Asteroid

func InitStorage(data []Asteroid) {
	AsteroidDB = data
	fmt.Printf("Success! the data has been stored. \n")
}

func GetAsteroids(offset int, limit int, filterHazardous bool, hazardousValue bool) []Asteroid {

	if filterHazardous {
		sourceData := FilterAsteroids(hazardousValue)

		// if offset has passed the last asteroid
		if offset >= len(sourceData) {
			return []Asteroid{} // empty slice
		}

		end := offset + limit
		if end > len(sourceData) {
			end = len(sourceData)
		}

		// return the slice window
		return sourceData[offset:end]
	}

	if offset >= len(AsteroidDB) {
		return []Asteroid{}
	}

	end := offset + limit
	if end > len(AsteroidDB) {
		end = len(AsteroidDB)
	}

	return AsteroidDB[offset:end]
}

func CreateAsteroid(newAsteroid Asteroid) {
	AsteroidDB = append(AsteroidDB, newAsteroid)
}

func DeleteAsteroid(id string) {
	for i := 0; i < len(AsteroidDB); i++ {
		if AsteroidDB[i].ID == id {
			AsteroidDB = append(AsteroidDB[:i], AsteroidDB[i+1:]...)
			return
		}
	}
}

func FilterAsteroids(hazardous bool) []Asteroid {
	var filtered []Asteroid
	for _, asteroid := range AsteroidDB {
		if asteroid.IsPotentiallyHazardousAsteroid == hazardous {
			filtered = append(filtered, asteroid)
		}
	}
	return filtered
}
