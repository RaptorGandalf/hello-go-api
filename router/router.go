package router

import (
	"github.com/RaptorGandalf/hello-go-api/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Database - Database object used by http handler functions
var Database repository.Database

// Setup - Creates a database connection and sets up routes
func Setup(r *gin.Engine, connection *gorm.DB) {
	Database = repository.GetDatabaseForConnection(connection)

	api := r.Group("/api")

	api.GET("/samples", GetSamples)
	api.GET("/samples/:id", GetSample)
	api.PUT("/samples/:id", UpdateSample)
	api.POST("/samples", CreateSample)
	api.DELETE("/samples/:id", DeleteSample)
}
