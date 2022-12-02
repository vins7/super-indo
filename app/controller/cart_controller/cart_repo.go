package cartcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	cartDb "github.com/vins7/super-indo/app/db/cart_db"
	pDb "github.com/vins7/super-indo/app/db/product_db"
	cartService "github.com/vins7/super-indo/app/service/cart_service"
)

func CartRoute(rg *gin.RouterGroup, db *gorm.DB, m gin.HandlerFunc) {
	r := rg.Group("/cart", m)
	u := NewCartController(cartService.NewCartService(cartDb.NewDBCart(db), pDb.NewDBProduct(db)))
	{
		r.POST("/add", u.Add)
		r.GET("/all", u.GetAll, m)
	}
}
