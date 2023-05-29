package models

type Object struct {
	BaseModel
	Key    string `json:"key" gorm:"uniqueIndex:idx_key_bucket"`
	Bucket string `json:"bucket" gorm:"foreignKey:BucketReferer;uniqueIndex:idx_key_bucket"`
	Size   int    `json:"size"`
	ETag   string `json:"ETag"`
}
