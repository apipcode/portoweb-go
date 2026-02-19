package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := "portfolio.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if column exists
	// Simplest way: try to add it, if fails, it might exist
	// Or check pragma table_info

	fmt.Println("Migrating database: Adding github_url to projects table...")

	query := `ALTER TABLE projects ADD COLUMN github_url TEXT DEFAULT '';`
	_, err = db.Exec(query)
	if err != nil {
		// Calculate if error is "duplicate column name"
		if err.Error() == "duplicate column name: github_url" {
			fmt.Println("Column github_url already exists. Skipping.")
			return
		}
		// If other error, verify if it is indeed duplicate or something else
		fmt.Printf("Warning/Error: %v\n", err)
	} else {
		fmt.Println("Success: github_url column added.")
	}
}
