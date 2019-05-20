package db

type DatabaseClient interface {
	Write(entityName string, obj interface{}) error
}