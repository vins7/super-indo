package controller

import (
	"time"

	c "github.com/vins7/super-indo/app/controller/user_controller"
	"github.com/vins7/super-indo/app/db"
	"github.com/vins7/super-indo/config"

	"github.com/gin-gonic/gin"
	"github.com/vins7/module-middleware/middleware"
	cart "github.com/vins7/super-indo/app/controller/cart_controller"
	cat "github.com/vins7/super-indo/app/controller/category_controller"
	p "github.com/vins7/super-indo/app/controller/product_controller"
)

type Route struct{}

func NewRoutes(r *gin.Engine) {

	rg := r.Group("/v1")
	cfg := config.GetConfig()
	m := middleware.NewJWTManager(cfg.Server.Secret, time.Duration(2)*time.Hour)

	dbUser := db.UserDB
	c.UserRoute(rg, dbUser)
	cat.CategoryRoute(rg, dbUser, m.Middleware(cfg.Server.Secret))
	p.ProductRoute(rg, dbUser, m.Middleware(cfg.Server.Secret))
	cart.CartRoute(rg, dbUser, m.Middleware(cfg.Server.Secret))
}
