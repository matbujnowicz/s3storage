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

func Connect() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Berlin",
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

	DbClient = &PostgresClient{
		Db: db,
	}
}

func (c *PostgresClient) Create(model interface{}) error {
	if err := c.Db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (c *PostgresClient) List(models []interface{}) error {
	if err := c.Db.Find(models).Error; err != nil {
		return err
	}
	return nil
}
