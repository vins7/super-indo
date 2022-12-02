package categorycontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	cDb "github.com/vins7/super-indo/app/db/category_db"
	cService "github.com/vins7/super-indo/app/service/category_service"
)

func CategoryRoute(rg *gin.RouterGroup, db *gorm.DB, m gin.HandlerFunc) {
	r := rg.Group("/category", m)
	u := NewCategoryController(cService.NewCategoryService(cDb.NewDBCategory(db)))
	{
		r.POST("/add", u.AddCategory)
		r.GET("/all", u.GetAllCategory, m)
		r.GET("/by-id", u.GetByID, m)
	}
}
