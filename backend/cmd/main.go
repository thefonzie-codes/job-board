package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./jobboard.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		log.Fatal("Could not read schema.sql:", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal("Could not apply schema:", err)
	}

	seedSQL, err := os.ReadFile("internal/db/seed.sql")
	if err != nil {
		log.Fatal("Failed to read seed file:", err)
	}

	_, err = db.Exec(string(seedSQL))
	if err != nil {
		log.Fatal("Failed to seed db", err)
	}

	fmt.Println("DB initialized w/ schema and seeded.")

	if err := printJobs(db); err != nil {
		log.Fatal("Failed to print jobs:", err)
	}
}

func printJobs(db *sql.DB) error {
	rows, err := db.Query(`SELECT id, title, company, location, type, tags, apply_url, posted_at FROM jobs`)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	fmt.Println("\nAvailable Jobs:")
	for rows.Next() {
		var (
			id       int
			title    string
			company  string
			location string
			jobType  string
			tags     string
			applyURL string
			postedAt string
		)

		err := rows.Scan(&id, &title, &company, &location, &jobType, &tags, &applyURL, &postedAt)
		if err != nil {
			return fmt.Errorf("row scan failed: %w", err)
		}

		fmt.Printf("🔹 [%d] %s @ %s (%s)\n", id, title, company, location)
		fmt.Printf("    Type: %s | Tags: %s\n", jobType, tags)
		fmt.Printf("    Apply: %s | Posted: %s\n\n", applyURL, postedAt)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration failed: %w", err)
	}

	return nil
}
