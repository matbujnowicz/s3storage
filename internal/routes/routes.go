package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.PUT("/:bucket", handlers.CreateBucket)
	r.PUT("/:bucket/:key", handlers.CreateObject)
	r.GET("/:bucket", handlers.ListObjects)
}
