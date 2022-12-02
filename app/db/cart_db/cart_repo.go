package cartdb

import (
	"github.com/vins7/super-indo/app/db"
)

type DBCartRepo interface {
	db.DBBaseRepo
	Update(interface{}) error
	GetCart(in interface{}) (data interface{}, e error)
}
