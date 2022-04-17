package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID         int       `gorm:"primary_key;" json:"id"`
	Name       string    `json:"name"`
	Products   []Product `gorm:"foreignkey:CategoryID" json:"products"`
}

func NewCategory(id int, name string) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}

func (Category) TableName() string {
	return "category"
}
