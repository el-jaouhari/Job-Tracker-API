package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

// RunMigrations executes the schema.sql file to set up the database
func RunMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Check if the enum type already exists
	var exists bool
	err = sqlDB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM pg_type WHERE typname = 'application_status_enum')",
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check enum existence: %w", err)
	}

	// If enum exists, assume migration has already run
	if exists {
		log.Println("Database schema already exists, skipping migration")
		return nil
	}

	// Read schema.sql file
	schemaPath := "db/schema.sql"
	// Try to find schema.sql relative to the executable
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		// Try alternative paths
		possiblePaths := []string{
			"./db/schema.sql",
			"/app/db/schema.sql",
			filepath.Join(filepath.Dir(os.Args[0]), "../db/schema.sql"),
		}
		found := false
		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				schemaPath = path
				found = true
				break
			}
		}
		if !found {
			log.Println("Warning: schema.sql not found, skipping migration. Tables may need to be created manually.")
			return nil
		}
	}

	schemaSQL, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	log.Println("Running database migrations...")

	// Execute the schema
	_, err = sqlDB.Exec(string(schemaSQL))
	if err != nil {
		// Check if it's a "already exists" error, which is okay
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			log.Println("Database schema already exists, skipping migration")
			return nil
		}
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}
