package main

import (
	"log"
	"os"

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
		databaseURL = "postgres://localhost:5432/job_tracker"
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	jobsRepository := repository.NewJobsRepository(db)
	jobsService := service.NewJobsService(jobsRepository)
	router := gin.Default()
	httpx.SetupRoutes(router, jobsService)
	router.Run(":8080")
}
