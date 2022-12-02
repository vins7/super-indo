package productcontroller

import (
	"net/http"

	"github.com/vins7/super-indo/app/model"
	pSvc "github.com/vins7/super-indo/app/service/product_service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	pSvc pSvc.ProductServiceRepo
}

func NewProductController(pSvc pSvc.ProductServiceRepo) *ProductController {

	return &ProductController{
		pSvc: pSvc,
	}
}

func (p *ProductController) AddProduct(c *gin.Context) {
	req := &model.Product{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	if err := p.pSvc.Add(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": "Success add product !"})
}

func (p *ProductController) GetAllProduct(c *gin.Context) {
	req := &model.GetProductByCatRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	res, err := p.pSvc.GetAll(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": res})
}

func (p *ProductController) GetByID(c *gin.Context) {
	req := &model.GetAllProductRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	res, err := p.pSvc.GetByID(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": res})
}
