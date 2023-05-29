package models

import "gorm.io/gorm"

type Bucket struct {
	gorm.Model
	Name string `json:"name" gorm:"primaryKey"`
}
