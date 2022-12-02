package productservice

import (
	"github.com/mitchellh/mapstructure"
	db "github.com/vins7/super-indo/app/db/product_db"
	"github.com/vins7/super-indo/app/model"
)

type ProductService struct {
	repo db.DBProductRepo
}

func NewProductService(repo db.DBProductRepo) ProductServiceRepo {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) Add(data *model.Product) error {
	return p.repo.Add(data)
}
func (p *ProductService) GetAll(req *model.GetProductByCatRequest) (out []*model.Product, e error) {

	out = []*model.Product{}
	res, err := p.repo.GetAllByCat(req.KategoryId)
	if err != nil {
		return out, err
	}

	if err := mapstructure.Decode(res, &out); err != nil {
		return out, err
	}
	return out, nil
}
func (p *ProductService) GetByID(pID string) (out *model.Product, e error) {
	out = &model.Product{}
	res, err := p.repo.GetByID(pID)
	if err != nil {
		return out, err
	}

	if err := mapstructure.Decode(res, &out); err != nil {
		return out, err
	}
	return out, nil
}
