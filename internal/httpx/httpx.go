package httpx

import (
	"net/http"
	"strings"

	"github.com/el-jaouhari/Job-Tracker-API/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, jobsService *service.JobsService) {
	router.GET("/jobs", func(c *gin.Context) {
		jobs, err := jobsService.GetJobs()
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, jobs)
	})

	router.POST("/jobs", func(c *gin.Context) {
		job := service.Job{}
		if err := c.ShouldBindJSON(&job); err != nil {
			if strings.Contains(err.Error(), "invalid character") {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Invalid JSON format",
					"details": err.Error(),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Invalid request body",
					"details": err.Error(),
				})
			}
			return
		}
		if err := jobsService.CreateJob(&job); err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Job created successfully"})
	})

	router.PUT("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		status := c.Query("status")
		if err := jobsService.UpdateJobStatus(id, status); err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
	})

	router.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := jobsService.DeleteJob(id); err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
	})
}

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *service.ValidationError:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Message,
			"field": e.Field,
		})
	case *service.StatusError:
		c.JSON(http.StatusBadRequest, gin.H{
			"error":          "Invalid application status: " + e.Status,
			"valid_statuses": e.ValidStatus,
		})
	default:
		if err == service.ErrJobNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Job not found",
			})
		} else if err == service.ErrInvalidID {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid job ID",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}
