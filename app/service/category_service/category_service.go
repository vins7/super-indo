package categoryservice

import (
	"github.com/mitchellh/mapstructure"
	db "github.com/vins7/super-indo/app/db/category_db"
	"github.com/vins7/super-indo/app/model"
)

type CategoryService struct {
	repo db.DBCategoryRepo
}

func NewCategoryService(repo db.DBCategoryRepo) CategoryServiceRepo {
	return &CategoryService{
		repo: repo,
	}
}

func (p *CategoryService) Add(data *model.Kategory) error {
	return p.repo.Add(data)
}
func (p *CategoryService) GetAll() (out []*model.Kategory, e error) {
	out = []*model.Kategory{}
	res, err := p.repo.GetAll()
	if err != nil {
		return
	}
	if err := mapstructure.Decode(res, &out); err != nil {
		return out, err
	}
	return out, nil
}
func (p *CategoryService) GetByID(pID string) (out *model.Kategory, e error) {
	out = &model.Kategory{}
	res, err := p.repo.GetByID(pID)
	if err != nil {
		return out, err
	}

	if err := mapstructure.Decode(res, &out); err != nil {
		return out, err
	}
	return out, nil
}
