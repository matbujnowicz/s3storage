package db

type Client interface {
	Create(model interface{})
	List(models []interface{})
}
