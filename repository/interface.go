package repository

type Interface interface {
	Set(id string, name string)
	GetById(id string) (name string, found bool)
	GetAll() (items map[string]string)
	Delete(id string) error
}