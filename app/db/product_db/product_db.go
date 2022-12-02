package productdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/vins7/super-indo/app/db/redis"
	"github.com/vins7/super-indo/app/model"
	"gorm.io/gorm"
)

type DBProduct struct {
	db *gorm.DB
}

func NewDBProduct(db *gorm.DB) DBProductRepo {
	return &DBProduct{
		db: db,
	}
}

func (d *DBProduct) Add(in interface{}) error {

	req, _ := in.(*model.Product)
	if err := d.db.Debug().Save(&req).Error; err != nil {
		return err
	}

	_ = redis.DelCache(context.Background(), req.KategoryId+"_ALLPRODUCT-*")
	return nil
}

func (d *DBProduct) GetAll() (interface{}, error) { return nil, nil }

func (d *DBProduct) GetAllByCat(catID string) (interface{}, error) {

	data := []*model.Product{}
	c, err := redis.GetCache(context.Background(), catID+"_ALLPRODUCT-*")
	if err != nil {

		if err := d.db.Debug().
			Where("kategory_id = ?", catID).
			Find(&data).Error; err != nil {

			return []*model.Product{}, err
		}

		j, _ := json.Marshal(data)
		_ = redis.SetCache(context.Background(), catID+"_ALLPRODUCT-*", string(j), 8*time.Hour)
		return data, nil
	}

	if err := json.Unmarshal([]byte(c), &data); err != nil {
		_ = redis.DelCache(context.Background(), catID+"_ALLPRODUCT-*")
		return nil, err
	}

	return data, nil
}

func (d *DBProduct) GetByID(productID string) (interface{}, error) {

	data := &model.Product{}
	if err := d.db.Debug().
		Where("id = ?", productID).
		First(&data).Error; err != nil {

		return nil, err
	}

	return data, nil
}
