package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID         int       `gorm:"primary_key;" json:"id"`
	Name       string    `json:"name"`
	Products   []Product `gorm:"foreignkey:CategoryID" json:"products"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
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
