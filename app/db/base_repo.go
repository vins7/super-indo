package db

type DBBaseRepo interface {
	Add(interface{}) error
	GetAll() (interface{}, error)
	GetByID(string) (interface{}, error)
}
