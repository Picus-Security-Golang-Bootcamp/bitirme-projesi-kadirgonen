package repo

import (
	"errors"

	model "HW/app/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) GetCart(userName string) *model.Cart {
	var cart model.Cart
	result := r.db.Preload("Items").Preload("Items.Product").Preload("Items.Product.Category").FirstOrCreate(&cart, model.Cart{UserName: userName})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &cart
}

func (r *CartRepository) CreatCart(c *model.Cart) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) UpdateCart(c *model.Cart) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) DeleteCart(userName string) error {
	result := r.db.Where(&model.Cart{UserName: userName}).Delete(&model.Cart{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) AddCartItem(c *model.CartItem) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) UpdateCartItem(c *model.CartItem) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) DeleteCartItem(c *model.CartItem) error {
	result := r.db.Delete(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CartRepository) DeleteCartItemsByCartId(cartId string) error {
	result := r.db.Where(&model.CartItem{CartID: cartId}).Delete(&model.CartItem{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
