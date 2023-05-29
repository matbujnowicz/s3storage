package db

type Client interface {
	Create(model interface{}) error
	List(models []interface{}) error
}

var DbClient Client
