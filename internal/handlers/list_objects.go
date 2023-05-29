package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func ListObjects(c *gin.Context) {
	params := db.ListParams{BucketName: c.Param("bucket"), Marker: c.Query("marker"), Max: c.Query("max-keys"), Prefix: c.Query("prefix")}

	var objects []models.Object

	if err := db.DbClient.ListObjects(&objects, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("listing objects resulted in error: %v", err),
		})
		c.Abort()
		return
	}

	// TODO: format message response accordingly
	c.JSON(http.StatusOK, objects)
}
