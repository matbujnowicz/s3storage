package handlers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func CreateObject(c *gin.Context) {
	bucketName := c.Param("bucket")
	objectKey := c.Param("key")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("reading form file failed with error: %v", err),
		})
		return
	}

	md5List, ok := c.Request.Header["Content-Md5"]
	if ok {
		actualMd5 := calculateMD5Digest(body)
		if md5List[0] != actualMd5 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "provided Content-MD5 value is different than calculated one",
			})
			return
		}
	}

	eTag := calculateEtag(body)
	object := models.Object{Key: objectKey, Bucket: bucketName, ETag: eTag}

	if err := db.DbClient.CreateObject(&object); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("object creation resulted in error: %v", err),
		})
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("uploads/%v/%v", bucketName, objectKey), body, 0644)
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

func calculateEtag(body []byte) string {
	hash := md5.Sum(body)
	etag := hex.EncodeToString(hash[:])
	return etag
}

func calculateMD5Digest(data []byte) string {
	hash := md5.Sum(data)
	base64Digest := base64.StdEncoding.EncodeToString(hash[:])
	return base64Digest
}
