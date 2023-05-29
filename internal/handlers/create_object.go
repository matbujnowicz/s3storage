package handlers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

const contentMd5Header = "Content-Md5"

func CreateObject(c *gin.Context) {
	bucketName := c.Param("bucket")
	objectKey := c.Param("key")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		xmlError(c, http.StatusBadRequest, "reading file body failed", err)
		return
	}

	md5List, ok := c.Request.Header[contentMd5Header]
	if ok {
		actualMd5 := calculateMD5Digest(body)
		if md5List[0] != actualMd5 {
			xmlError(c, http.StatusInternalServerError, "provided Content-MD5 value is different than calculated one", nil)
			return
		}
	}

	eTag := calculateEtag(body)
	object := models.Object{Key: objectKey, Bucket: bucketName, ETag: eTag, Size: len(body)}

	if err := db.DbClient.CreateObject(&object); err != nil {
		xmlError(c, http.StatusInternalServerError, "object creation failed", err)
		return
	}

	if err := writeFile(body, bucketName, objectKey); err != nil {
		// if saving file for an object did not succeed we should remove previously created database entry for the object
		if deletionErr := db.DbClient.Delete(&object); deletionErr != nil {
			xmlError(c, http.StatusInternalServerError, "file saving and record deletion failed", deletionErr)
			return
		}
		xmlError(c, http.StatusInternalServerError, "file saving failed", err)
		return
	}

	c.Header("ETag", eTag)
	c.Status(http.StatusOK)
}

func writeFile(body []byte, bucketName string, objectKey string) error {
	path := fmt.Sprintf("uploads/%v/%v", bucketName, objectKey)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, body, 0644)
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
