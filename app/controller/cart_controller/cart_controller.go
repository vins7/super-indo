package cartcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/vins7/module-middleware/middleware/model"
	"github.com/vins7/super-indo/app/model"
	cartSvc "github.com/vins7/super-indo/app/service/cart_service"
)

type CartController struct {
	cartSvc cartSvc.CartServiceRepo
}

func NewCartController(cartSvc cartSvc.CartServiceRepo) *CartController {

	return &CartController{
		cartSvc: cartSvc,
	}
}

func (p *CartController) Add(c *gin.Context) {
	req := &model.AddCartRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	res, ok := c.Get("user-data")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "error while get context"})
		return
	}

	data := res.(*m.UserClaims)
	req.UserId = data.UserID
	if err := p.cartSvc.Add(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": "Success add !"})
}

func (p *CartController) GetAll(c *gin.Context) {
	res, ok := c.Get("user-data")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "error while get context"})
		return
	}

	data := res.(*m.UserClaims)
	res, err := p.cartSvc.GetAllByUser(data.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"RESPONSE": res})
}
