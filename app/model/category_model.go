package model

type Kategory struct {
	BaseModel
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
}

func (Kategory) TableName() string {
	return "t_Kategory"
}
