package model

type Product struct {
	BaseModel
	Title       string   `gorm:"column:title" json:"title"`
	Description string   `gorm:"column:description" json:"description"`
	KategoryId  string   `gorm:"column:kategory_id;size:100" json:"kategory_id"`
	Kategory    Kategory `gorm:"foreignKey:KategoryId" json:"-"`
	Price       int      `gorm:"column:price" json:"price"`
}

func (Product) TableName() string {
	return "t_Product"
}

type GetProductByCatRequest struct {
	KategoryId string `json:"kategory_id"`
}

type GetAllProductRequest struct {
	ProductID string `json:"product_id"`
}
