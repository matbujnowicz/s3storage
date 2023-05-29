package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func CreateObject(c *gin.Context) {
	bucketName := c.Param("bucket")
	objectKey := c.Param("key")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("reading form file failed with error: %v", err),
		})
		return
	}

	eTag, eTagErr := calculateEtag(file)
	if eTagErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("could not calculate etag, error: %v", err),
		})
		return
	}

	md5, ok := c.Request.Header["Content-Md5"]
	if ok && md5[0] != eTag {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "provided Content-MD5 value is different than calculated ETag",
		})
		return
	}

	object := models.Object{Key: objectKey, Bucket: bucketName, FileName: file.Filename, ETag: eTag}

	if err := db.DbClient.CreateObject(&object); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("object creation resulted in error: %v", err),
		})
		return
	}

	err = c.SaveUploadedFile(file, fmt.Sprintf("uploads/%v/%v/%v", bucketName, objectKey, file.Filename))
	if err != nil {
		// if saving file for an object did not succeed we should remove previously created database entry for the object
		if deletionErr := db.DbClient.Delete(&object); deletionErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("file saving resulted in error: %v and record deletion resulted in error: %v", err, deletionErr),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("file saving resulted in error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ETag": eTag,
	})
}

func calculateEtag(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	etag := hex.EncodeToString(hash.Sum(nil))
	return etag, nil
}
