package controller

import (
	userdb "github.com/vins7/super-indo/app/db/user_db"
	userservice "github.com/vins7/super-indo/app/service/user_services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/user-management")
	u := NewUserController(userservice.NewUserService(userdb.NewDBUser(db)))

	r.POST("/login", u.Login)
	r.POST("/create-user", u.CreateUser)
}
