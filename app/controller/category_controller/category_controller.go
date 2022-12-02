package categorycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vins7/super-indo/app/model"
	pSvc "github.com/vins7/super-indo/app/service/category_service"
)

type CategoryController struct {
	pSvc pSvc.CategoryServiceRepo
}

func NewCategoryController(pSvc pSvc.CategoryServiceRepo) *CategoryController {

	return &CategoryController{
		pSvc: pSvc,
	}
}

func (p *CategoryController) AddCategory(c *gin.Context) {
	req := &model.Kategory{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	if err := p.pSvc.Add(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": "Success add category !"})
}

func (p *CategoryController) GetAllCategory(c *gin.Context) {
	res, err := p.pSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": res})
}

func (p *CategoryController) GetByID(c *gin.Context) {
	req := &model.Kategory{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	res, err := p.pSvc.GetByID(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": res})
}
