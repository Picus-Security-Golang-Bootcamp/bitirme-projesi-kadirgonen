package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         int      `gorm:"primary_key;auto_increment" json:"id"`
	Name       string   `json:"name"`
	StockCode  string   `gorm:"unique" json:"stock_code"`
	Cost       float64  `json:"cost"`
	Number     int      `json:"number"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
}

func NewProduct(name string, stockCode string, cost float64, number int, categoryID int) *Product {
	return &Product{
		Name:       name,
		StockCode:  stockCode,
		Cost:       cost,
		Number:     number,
		CategoryID: categoryID,
	}
}

func (Product) TableName() string {
	return "product"
}
