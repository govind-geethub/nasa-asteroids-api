package main

import (
	"fmt"
)

var AsteroidDB []Asteroid

func InitStorage(data []Asteroid) {
	AsteroidDB = data
	fmt.Printf("Success! the data has been stored. \n")
}

func GetAsteroids() []Asteroid {
	return AsteroidDB
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
