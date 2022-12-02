package productdb

import "github.com/vins7/super-indo/app/db"

type DBProductRepo interface {
	db.DBBaseRepo
	GetAllByCat(catID string) (interface{}, error)
}
