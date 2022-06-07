package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         string       `gorm:"primary_key;" json:"id"`
	UserName   string       `json:"user_name"`
	Status     string       `json:"status"`
	Name       string       `json:"name"`
	Address    string       `json:"address"`
	Phone      string       `json:"phone"`
	Items      []*OrderItem `gorm:"foreignkey:OrderID;" json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID    string    `gorm:"primary_key" json:"order_id"`
	ProductID  int       `gorm:"primary_key" json:"product_id"`
	Product    *Product  `gorm:"foreignkey:ProductID;references:ID" json:"product"`
}

func NewOrder(userName string, name string, address string, phone string) (*Order, error) {
	if len(userName) == 0 {
		return nil, fmt.Errorf("userName field is required")
	}
	return &Order{
		UserName: userName,
		Status:   "unachieved",
		Name:     name,
		Address:  address,
		Phone:    phone,
		Items:    nil,
	}, nil
}

func NewOrderItem(orderId string, productId int) (*OrderItem, error) {
	return &OrderItem{
		OrderID:   orderId,
		ProductID: productId,
	}, nil
}

func (Order) TableName() string {
	return "order"
}

func (b *Order) BeforeCreate(db *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}

func (OrderItem) TableName() string {
	return "order_item"
}

func (b *Order) AddItem(item *OrderItem) error {
	b.Items = append(b.Items, item)
	return nil
}
