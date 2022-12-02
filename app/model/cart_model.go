package model

type Cart struct {
	BaseModel
	ProductId  string  `gorm:"column:product_id;size:100" json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductId" json:"-"`
	UserId     string  `gorm:"column:user_id;size:100"`
	User       User    `gorm:"foreignKey:UserId" json:"-"`
	Quantity   string  `gorm:"column:qty;size:100"`
	Price      string  `gorm:"column:price;size:100"`
	TotalPrice string  `gorm:"column:total_price;size:100"`
	Status     string  `gorm:"column:status;size:100"`
}

type AddCartRequest struct {
	ProductId string `gorm:"column:product_id;size:100" json:"product_id"`
	UserId    string `gorm:"column:user_id;size:100" json:"-"`
	Quantity  string `gorm:"column:qty;size:100" json:"qty"`
}

type GetCartResponse struct {
	Id         string `json:"id"`
	ProductId  string `gorm:"column:product_id;size:100" json:"product_id"`
	Quantity   string `gorm:"column:qty;size:100" json:"qty"`
	Price      string `gorm:"column:price;size:100" json:"price"`
	TotalPrice string `gorm:"column:total_price;size:100" json:"total_price" `
	Status     string `gorm:"column:status;size:100" json:"Status"`
}

func (Cart) TableName() string {
	return "t_Cart"
}

type GetCartRequest struct {
	UserID string
	Status string `json:"STATUS"`
}

type AcceptCartRequest struct {
	CartID string
}
