package categorydb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/vins7/super-indo/app/db"
	"github.com/vins7/super-indo/app/db/redis"
	"github.com/vins7/super-indo/app/model"
	"gorm.io/gorm"
)

type DBCategory struct {
	db *gorm.DB
}

func NewDBCategory(db *gorm.DB) db.DBBaseRepo {
	return &DBCategory{
		db: db,
	}
}

func (d *DBCategory) Add(in interface{}) error {

	req, _ := in.(*model.Kategory)
	if err := d.db.Debug().Save(&req).Error; err != nil {
		return err
	}

	_ = redis.DelCache(context.Background(), "ALLCATEGORY-*")
	return nil
}

func (d *DBCategory) GetAll() (interface{}, error) {

	data := []*model.Kategory{}
	c, err := redis.GetCache(context.Background(), "ALLCATEGORY-*")
	if err != nil {

		if err := d.db.Debug().Find(&data).Error; err != nil {
			return data, err
		}

		j, _ := json.Marshal(data)
		_ = redis.SetCache(context.Background(), "ALLCATEGORY-*", string(j), 8*time.Hour)
		return data, nil
	}

	if err := json.Unmarshal([]byte(c), &data); err != nil {
		_ = redis.DelCache(context.Background(), "ALLCATEGORY-*")
		return nil, err
	}

	return data, nil
}

func (d *DBCategory) GetByID(productID string) (interface{}, error) {

	data := &model.Kategory{}
	if err := d.db.Debug().
		Where("id = ?", productID).
		First(&data).Error; err != nil {

		return nil, err
	}

	return data, nil
}
