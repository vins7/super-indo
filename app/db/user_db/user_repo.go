package userdb

import "github.com/vins7/super-indo/app/model"

type DBUserRepo interface {
	Login(*model.UserRequest) (*model.User, error)
	CreateUser(req *model.User) error
}
