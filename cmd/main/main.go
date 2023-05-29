package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/handlers"
)

func main() {
	r := gin.Default()

	r.PUT("/:bucket", handlers.CreateBucket)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
