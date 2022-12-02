package cartservice

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
	db "github.com/vins7/super-indo/app/db/cart_db"
	dbProduct "github.com/vins7/super-indo/app/db/product_db"
	"github.com/vins7/super-indo/app/model"
)

type CartService struct {
	repo        db.DBCartRepo
	repoProduct dbProduct.DBProductRepo
}

func NewCartService(repo db.DBCartRepo, repoProduct dbProduct.DBProductRepo) CartServiceRepo {
	return &CartService{
		repo:        repo,
		repoProduct: repoProduct,
	}
}

func (p *CartService) Add(data *model.AddCartRequest) error {

	product, err := p.repoProduct.GetByID(data.ProductId)
	if err != nil {
		return err
	}

	dataProduct := &model.Product{}
	if err := mapstructure.Decode(product, &dataProduct); err != nil {
		return err
	}

	qty, err := strconv.Atoi(data.Quantity)
	if err != nil {
		return err
	}

	totalPrice := qty * dataProduct.Price

	return p.repo.Add(&model.Cart{
		ProductId:  data.ProductId,
		UserId:     data.UserId,
		Quantity:   data.Quantity,
		TotalPrice: strconv.Itoa(totalPrice),
		Price:      strconv.Itoa(dataProduct.Price),
		Status:     "PENDING",
	})
}
func (p *CartService) GetAllByUser(userID string) (out []*model.Cart, e error) {
	out = []*model.Cart{}
	res, err := p.repo.GetCart(userID)
	if err != nil {
		return
	}
	if err := mapstructure.Decode(res, &out); err != nil {
		return out, err
	}
	return out, nil
}
