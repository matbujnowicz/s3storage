package models

type Bucket struct {
	BaseModel
	Name string `json:"name" gorm:"unique"`
}
