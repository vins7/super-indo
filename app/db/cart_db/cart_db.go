package cartdb

import (
	"github.com/vins7/super-indo/app/model"
	"gorm.io/gorm"
)

type DBCart struct {
	db *gorm.DB
}

func NewDBCart(db *gorm.DB) DBCartRepo {
	return &DBCart{
		db: db,
	}
}

func (t *DBCart) Add(in interface{}) error {
	req, _ := in.(*model.Cart)
	return t.db.Debug().Save(&req).Error
}

func (t *DBCart) Update(in interface{}) error {
	req, _ := in.(*model.Cart)
	return t.db.Debug().Updates(&req).Error
}

func (t *DBCart) GetCart(in interface{}) (data interface{}, e error) {
	UserID, _ := in.(string)
	data = []*model.Cart{}
	qry := t.db.Debug().Preload("Product").Preload("User").Where("user_id = ?", UserID)
	if e = qry.Find(&data).Error; e != nil {
		return data, e
	}
	return data, nil
}

func (t *DBCart) GetByID(id string) (data interface{}, e error) {
	if e = t.db.Debug().First(&data).Error; e != nil {
		return nil, e
	}
	return data, nil
}

func (t *DBCart) GetAll() (data interface{}, e error) { return nil, nil }
