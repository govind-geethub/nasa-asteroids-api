package main

import (
	"fmt"
)

var AsteroidDB []Asteroid

func InitStorage(data []Asteroid) {
	AsteroidDB = data
	fmt.Printf("Success! the data has been extracted. \n")
}
