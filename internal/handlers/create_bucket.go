package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBucket(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}