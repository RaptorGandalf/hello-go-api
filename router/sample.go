package router

import (
	"github.com/RaptorGandalf/hello-go-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetSamples - Returns all samples in database
func GetSamples(c *gin.Context) {
	samples, err := Database.SampleRepo.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"samples": samples,
	})
}

// GetSample - Returns a single Sample with the specified ID
func GetSample(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": err})
		return
	}

	result, err := Database.SampleRepo.Get(uid)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	if result == nil {
		c.AbortWithStatusJSON(404, gin.H{"errors": "Not found."})
		return
	}

	c.JSON(200, result)
}

// CreateSample - Creates a new Sample
func CreateSample(c *gin.Context) {
	reqBody := new(model.Sample)

	err := c.BindJSON(reqBody)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid sample"})
		return
	}

	err = Database.SampleRepo.Create(reqBody)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "failed to create sample"})
		return
	}

	c.JSON(201, reqBody)
}

// UpdateSample - Updates an existing Sample
func UpdateSample(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "id param required"})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid uuid"})
		return
	}

	sample, err := Database.SampleRepo.Get(uid)
	if err != nil || sample == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "not found"})
		return
	}

	err = c.BindJSON(&sample)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err})
		return
	}

	err = Database.SampleRepo.Update(sample)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, sample)
}

// DeleteSample - Deletes a sample from the database
func DeleteSample(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "id param required"})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid uuid"})
		return
	}

	err = Database.SampleRepo.Delete(uid)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}

	c.Status(200)
}
