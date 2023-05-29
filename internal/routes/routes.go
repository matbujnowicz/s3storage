package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.PUT("/:bucket", handlers.CreateBucket)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
