package userservice

import "github.com/vins7/super-indo/app/model"

type UserServiceRepo interface {
	Login(*model.UserRequest) (*model.UserResponse, error)
	CreateUser(req *model.CreateUserRequest) error
}
