package db

import (
	"fmt"
	"os"

	"github.com/matbujnowicz/s3storage/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)

type PostgresClient struct {
	Db *gorm.DB
}

var DbClient PostgresClient

func Connect() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database, err: %v", err))
	}

	db.AutoMigrate(&models.Bucket{}, &models.Object{})

	DbClient = PostgresClient{
		Db: db,
	}
}

func Create(c *PostgresClient, model interface{}) {
	c.Db.Create(model)
}

func List(c *PostgresClient, models interface{}) {
	c.Db.Find(models)
}
