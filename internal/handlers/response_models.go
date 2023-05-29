package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func xmlError(c *gin.Context, code int, message string, err error) {
	errMessage := ""
	if err != nil {
		errMessage = err.Error()
	}
	c.XML(code, Error{
		Message: message,
		Error:   errMessage,
		Code:    code,
	})
}

type Error struct {
	Message string
	Error   string `xml:",omitempty"`
	Code    int
}

type Contents struct {
	Key          string
	LastModified time.Time
	ETag         string
	Size         int
}

type ListBucketResult struct {
	Name        string
	Prefix      string `xml:",omitempty"`
	Marker      string `xml:",omitempty"`
	MaxKeys     string `xml:",omitempty"`
	IsTruncated bool
	Contents    []Contents
}
