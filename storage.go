package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// global database connection pool pointer
var DB *pgxpool.Pool

// InitStorage hooks the backend to postgres directly
func InitStorage() {
	connStr := "postgres://postgres:Avin14AG07@localhost:5432/nasa_space_project"

	var err error
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v \n", err)
	}

	// verify the connections work
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database ping failed: %v \n", err)
	}

	fmt.Println("Success! Connected to the PostgreSQL.\n")
}

// inserts a new asteroid row into the database
func CreateAsteroid(newAsteroid Asteroid) {
	query := `
			INSERT INTO asteroids (id, name, absolute_magnitude_h, is_potentially_hazardous_asteroid)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO NOTHING;`

	_, err := DB.Exec(context.Background(), query,
		newAsteroid.ID,
		newAsteroid.Name,
		newAsteroid.AbsoluteMagnitudeH,
		newAsteroid.IsPotentiallyHazardousAsteroid,
	)

	if err != nil {
		log.Printf("Failed to insert asteroid %s into the database: %v \n", newAsteroid.Name, err)
	}
}

// GetAsteroids gives all the paginated data
func GetAsteroids(offset int, limit int, filterHazardous bool, hazardousValue bool) []Asteroid {
	var rows pgx.Rows
	var err error

	baseQuery := `SELECT id, name, absolute_magnitude_h, is_potentially_hazardous_asteroid FROM asteroids`

	if filterHazardous {
		sqlQuery := baseQuery + ` WHERE is_potentially_hazardous_asteroid = $1 LIMIT $2 OFFSET $3;`
		rows, err = DB.Query(context.Background(), sqlQuery, hazardousValue, limit, offset)
	} else {
		sqlQuery := baseQuery + ` LIMIT $1 OFFSET $2;`
		rows, err = DB.Query(context.Background(), sqlQuery, limit, offset)
	}

	if err != nil {
		log.Printf("Qe=uery failed: %v \n", err)
	}
	defer rows.Close()

	var asteroids []Asteroid
	for rows.Next() {
		var a Asteroid
		err := rows.Scan(&a.ID, &a.Name, &a.AbsoluteMagnitudeH, &a.IsPotentiallyHazardousAsteroid)
		if err != nil {
			log.Printf("Row scan failed: %v \n", err)
		}
		asteroids = append(asteroids, a)
	}

	return asteroids
}

// DeleteAsteroid drops the row matching the requested ID
func DeleteAsteroid(id string) {
	query := `DELETE FROM asteroids WHERE id = $1;`

	_, err := DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Failed to delete asteroid with ID %s : %v \n", id, err)
	}
}
