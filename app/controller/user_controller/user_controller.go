package controller

import (
	"net/http"

	"github.com/vins7/super-indo/app/model"

	uSvc "github.com/vins7/super-indo/app/service/user_services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uSvc uSvc.UserServiceRepo
}

func NewUserController(uSvc uSvc.UserServiceRepo) *UserController {
	return &UserController{
		uSvc: uSvc,
	}
}

func (u *UserController) Login(c *gin.Context) {
	var req *model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	res, err := u.uSvc.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": res})
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req *model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	if err := u.uSvc.CreateUser(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "Success create user"})
}
