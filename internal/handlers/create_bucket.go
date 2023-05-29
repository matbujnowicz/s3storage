package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func CreateBucket(c *gin.Context) {
	bucketName := c.Param("bucket")

	bucket := models.Bucket{Name: bucketName}

	if err := db.DbClient.Create(&bucket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("bucket creation resulted in error: %v", err),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Bucket %v created", bucketName),
	})
}
