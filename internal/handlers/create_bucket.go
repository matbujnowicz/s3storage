package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func CreateBucket(c *gin.Context) {
	bucketName := c.Param("bucket")

	bucket := models.Bucket{Name: bucketName}

	if err := db.DbClient.CreateBucket(&bucket); err != nil {
		xmlError(c, http.StatusInternalServerError, "bucket creation failed", err)
		return
	}

	c.Status(http.StatusOK)
}
