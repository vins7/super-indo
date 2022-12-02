package userdb

import (
	"errors"
	"fmt"

	"github.com/vins7/super-indo/app/model"

	"gorm.io/gorm"
)

type DBUser struct {
	db *gorm.DB
}

func NewDBUser(db *gorm.DB) DBUserRepo {
	return &DBUser{
		db: db,
	}
}

func (d *DBUser) CreateUser(in *model.User) error {

	if err := d.db.Debug().Save(&in).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%s", "username atau password anda salah !")
		}
		return err
	}

	return nil
}

func (d *DBUser) Login(in *model.UserRequest) (*model.User, error) {

	data := &model.User{}
	if err := d.db.Debug().Where("username = ?", in.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s", "username atau password anda salah !")
		}
		return nil, err
	}

	return data, nil
}
