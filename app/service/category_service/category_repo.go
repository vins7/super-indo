package categoryservice

import "github.com/vins7/super-indo/app/model"

type CategoryServiceRepo interface {
	Add(*model.Kategory) error
	GetAll() ([]*model.Kategory, error)
	GetByID(string) (*model.Kategory, error)
}
