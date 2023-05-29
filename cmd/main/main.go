package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matbujnowicz/s3storage/internal/db"
	"github.com/matbujnowicz/s3storage/internal/routes"
)

func main() {
	r := gin.Default()

	db.Connect()
	routes.SetupRoutes(r)

	r.Run()
}
