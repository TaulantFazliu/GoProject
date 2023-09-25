package database

import (
	"database/sql"
	"fmt"
)

func ConnectToDatabase() (*sql.DB, error) {
	connStr := "user=postgres dbname=go_project sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}
