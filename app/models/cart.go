package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID         string      `gorm:"primary_key;" json:"id"`
	UserName   string      `json:"user_name"`
	Items      []*CartItem `gorm:"foreignkey:CartID;" json:"items"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
	Deleted_at time.Time   `json:"deleted_at"`
}

type CartItem struct {
	gorm.Model
	CartID     string    `gorm:"primary_key" json:"cart_id"`
	ProductID  int       `gorm:"primary_key" json:"product_id"`
	Product    *Product  `gorm:"foreignkey:ProductID;references:ID" json:"product"`
	Number     int       `gorm:"not null" json:"number"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

func NewCart(userName string) (*Cart, error) {
	if len(userName) == 0 {
		return nil, fmt.Errorf("userName field is required")
	}
	return &Cart{
		UserName: userName,
		Items:    nil,
	}, nil
}

func NewCartItem(id string, productId int, number int) (*CartItem, error) {
	if number <= 0 {
		return nil, fmt.Errorf("quantity must be greater than zero")
	}
	return &CartItem{
		CartID:    id,
		ProductID: productId,
		Number:    number,
	}, nil
}

func (Cart) TableName() string {
	return "cart"
}

func (c *Cart) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}

func (CartItem) TableName() string {
	return "cart_item"
}

func (c *Cart) SearchItemByProductId(productId int) (int, *CartItem) {
	for i, n := range c.Items {
		if n.ProductID == productId {
			return i, n
		}
	}
	return -1, nil
}
