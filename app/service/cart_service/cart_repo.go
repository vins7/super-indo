package cartservice

import "github.com/vins7/super-indo/app/model"

type CartServiceRepo interface {
	Add(*model.AddCartRequest) error
	GetAllByUser(string) ([]*model.Cart, error)
}
