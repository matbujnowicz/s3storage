package models

type Object struct {
	BaseModel
	Key    string `gorm:"uniqueIndex:idx_key_bucket"`
	Bucket string `gorm:"uniqueIndex:idx_key_bucket"`
	Size   int
	ETag   string
}
