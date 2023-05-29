package models

type Bucket struct {
	BaseModel
	Name string `gorm:"unique"`
}
