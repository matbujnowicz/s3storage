package db

import (
	"fmt"
	"os"
	"strconv"

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

func (c *PostgresClient) ListObjects(objects *[]models.Object, params ListParams) error {
	tx := c.Db.Where("bucket = ?", params.BucketName)

	if params.Marker != "" {
		markerObject := models.Object{}
		if err := tx.Where("key = ?", params.Marker).First(&markerObject).Error; err != nil {
			return err
		}

		tx = c.Db.Where("bucket = ?", params.BucketName).Where("id > ?", markerObject.ID)
	}

	if params.Prefix != "" {
		tx.Where("key LIKE ?", params.Prefix+"%")
	}

	if params.Max != "" {
		maxInt, err := strconv.Atoi(params.Max)
		if err != nil || maxInt < 1 {
			return fmt.Errorf("max-keys value of %v is invalid", params.Max)
		}
		tx.Limit(maxInt)
	}

	if err := tx.Find(objects).Error; err != nil {
		return err
	}
	return nil
}
