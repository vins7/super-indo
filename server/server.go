package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vins7/super-indo/app/controller"
	"github.com/vins7/super-indo/config"
)

func Server() {
	r := gin.Default()
	cfg := config.GetConfig()
	controller.NewRoutes(r)

	r.Run(":" + cfg.Server.Port)
}
