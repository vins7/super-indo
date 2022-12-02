package productservice

import "github.com/vins7/super-indo/app/model"

type ProductServiceRepo interface {
	Add(*model.Product) error
	GetAll(*model.GetProductByCatRequest) ([]*model.Product, error)
	GetByID(string) (*model.Product, error)
}
