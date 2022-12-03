package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	BaseModel
	Name     string `gorm:"column:name" json:"name"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email;size:100" json:"email"`
}

func (User) TableName() string {
	return "t_User"
}

type CreateUserRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"Name"`
	Email    string `json:"email"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email   string `json:"email"`
	UserID  string `json:"userId"`
	Name    string `json:"name"`
	Counter int    `json:"counter"`
}
