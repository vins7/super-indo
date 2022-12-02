package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id string `json:"id" gorm:"column:ID;primaryKey"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	return nil
}
