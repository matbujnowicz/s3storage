package models

import "time"

// we are using this BaseModel instead of gorm's default Model to disable SOFT DELETE
// as we use UNIQUE indexes the SOFT DELETE is not effective for us (as it does not free up the UNIQUE indexes)
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
