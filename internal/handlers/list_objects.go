package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/models"
)

func ListObjects(c *gin.Context) {
	params := db.ListParams{
		BucketName: c.Param("bucket"),
		Marker:     c.Query("marker"),
		Max:        c.Query("max-keys"),
		Prefix:     c.Query("prefix"),
	}

	var objects []models.Object

	if err := db.DbClient.ListObjects(&objects, params); err != nil {
		xmlError(c, http.StatusInternalServerError, "listing objects failed", err)
		return
	}

	response := generateListBucketResult(objects, params)
	c.XML(http.StatusOK, response)
}

func generateListBucketResult(objects []models.Object, params db.ListParams) ListBucketResult {
	var contents []Contents
	for _, object := range objects {
		contents = append(contents, Contents{
			Key:          object.Key,
			LastModified: object.UpdatedAt,
			ETag:         object.ETag,
			Size:         object.Size,
		})
	}

	return ListBucketResult{
		Name:        params.BucketName,
		Prefix:      params.Prefix,
		Marker:      params.Marker,
		MaxKeys:     params.Max,
		IsTruncated: false,
		Contents:    contents,
	}
}
