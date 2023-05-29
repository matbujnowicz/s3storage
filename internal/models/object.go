package models

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Key    string `json:"key" gorm:"unique"`
	Bucket string `json:"bucket"`
	File   []byte `json:"file"`
}
