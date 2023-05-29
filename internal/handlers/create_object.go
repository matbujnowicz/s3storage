package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func CreateObject(c *gin.Context) {
	// ensure that required headers for create object are present
	contentMD5 := c.GetHeader("Content-MD5")
	if contentMD5 != "Content-MD5" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-MD5 header missing"})
		c.Abort()
		return
	}

	bucketName := c.Param("bucket")
	// TODO: should I check if this bucket exists?
	objectKey := c.Param("key")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("reading form file failed with error: %v", err),
		})
		c.Abort()
		return
	}

	err = c.SaveUploadedFile(file, fmt.Sprintf("uploads/%v/%v", objectKey, file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("file saving resulted in error: %v", err),
		})
		return
	}

	object := models.Object{Key: objectKey, Bucket: bucketName}

	if err := db.DbClient.Create(&object); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("object creation resulted in error: %v", err),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("object %v created", objectKey),
	})
}
