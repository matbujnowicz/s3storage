package models

type Object struct {
	BaseModel
	Key      string `json:"key" gorm:"uniqueIndex:idx_key_bucket"`
	FileName string `json:"file"`
	Bucket   string `json:"bucket" gorm:"foreignKey:BucketReferer;uniqueIndex:idx_key_bucket"`
	ETag     string `json:"ETag"`
}
