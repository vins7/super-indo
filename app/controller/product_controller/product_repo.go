package productcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	pDb "github.com/vins7/super-indo/app/db/product_db"
	pService "github.com/vins7/super-indo/app/service/product_service"
)

func ProductRoute(rg *gin.RouterGroup, db *gorm.DB, m gin.HandlerFunc) {
	r := rg.Group("/product", m)
	u := NewProductController(pService.NewProductService(pDb.NewDBProduct(db)))
	{
		r.POST("/add", u.AddProduct)
		r.GET("/all", u.GetAllProduct, m)
		r.GET("/by-id", u.GetByID, m)
	}
}
