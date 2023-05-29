package repositories

type DatabaseClient interface {
	CreateBucket(name string)
	CreateFile()
}
