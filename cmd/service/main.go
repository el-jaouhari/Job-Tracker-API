package main

import (
	"log"
	"os"

	"github.com/el-jaouhari/Job-Tracker-API/internal/database"
	"github.com/el-jaouhari/Job-Tracker-API/internal/httpx"
	"github.com/el-jaouhari/Job-Tracker-API/internal/repository"
	"github.com/el-jaouhari/Job-Tracker-API/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = os.Getenv("POSTGRES_URL")
	}
	if databaseURL == "" {
		// Default for local development
		databaseURL = "postgres://localhost:5432/job_tracker?sslmode=disable"
		log.Println("Warning: Using default database URL for local development")
	}

	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully")

	if err := database.RunMigrations(db); err != nil {
		log.Printf("Warning: Migration failed (this is okay if schema already exists): %v", err)
	}

	jobsRepository := repository.NewJobsRepository(db)
	jobsService := service.NewJobsService(jobsRepository)
	router := gin.Default()
	httpx.SetupRoutes(router, jobsService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
