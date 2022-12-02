package userservice

import (
	"fmt"
	"time"

	m "github.com/vins7/module-middleware/middleware"
	userdb "github.com/vins7/super-indo/app/db/user_db"
	"github.com/vins7/super-indo/app/model"
	"github.com/vins7/super-indo/config"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userdb userdb.DBUserRepo
}

func NewUserService(userdb userdb.DBUserRepo) UserServiceRepo {
	return &UserService{
		userdb: userdb,
	}
}

func (u *UserService) CreateUser(req *model.CreateUserRequest) error {

	pass, err := m.HashPassword(req.Password)
	if err != nil {
		return err
	}

	data := &model.User{
		Username: req.Username,
		Password: string(pass),
		Email:    req.Email,
		Name:     req.Name,
	}

	if err := u.userdb.CreateUser(data); err != nil {
		return err
	}

	return nil
}

func (u *UserService) Login(req *model.UserRequest) (*model.UserResponse, error) {

	res, err := u.userdb.Login(req)
	if err != nil {
		return nil, err
	}

	return u.validation(res.Id, res.Name, res.Username, res.Email, req.Password, res.Password)
}

func (u *UserService) validation(userId, name, username, email, reqpassword, respassword string) (*model.UserResponse, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(respassword), []byte(reqpassword)); err != nil {
		return nil, fmt.Errorf("%s", "username atau password salah !")
	}

	cfg := config.GetConfig()
	token, err := m.NewJWTManager(cfg.Server.Secret, time.Duration(2)*time.Hour).
		GenerateToken(email, userId, username)
	if err != nil {
		return nil, err
	}

	out := &model.UserResponse{
		Nama:     name,
		Username: username,
		Token:    token,
	}

	return out, nil
}
