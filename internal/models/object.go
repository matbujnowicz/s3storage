package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Key      string `json:"key" gorm:"uniqueIndex:idx_key_bucket"`
	FileName string `json:"file"`
	Bucket   string `json:"bucket" gorm:"foreignKey:BucketReferer;uniqueIndex:idx_key_bucket"`
}
