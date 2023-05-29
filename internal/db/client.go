package db

import "github.com/matbujnowicz/s3storage/internal/models"

type ListParams struct {
	BucketName string
	Prefix     string
	Marker     string
	Max        string
}

type Client interface {
	Delete(model interface{}) error
	CreateBucket(bucket *models.Bucket) error
	CreateObject(object *models.Object) error
	ListObjects(objects *[]models.Object, params ListParams) error
}

var DbClient Client
