package repo

import (
	model "HW/app/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetOrder(userName string, id string) *model.Order {
	var order model.Order
	r.db.Where(&model.Order{ID: id, UserName: userName}).
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		FirstOrCreate(&order)

	return &order
}

func (r *OrderRepository) CreateOrder(c *model.Order) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepository) GetAllOrders(userName string) []model.Order {
	var orders []model.Order
	r.db.Where(&model.Order{UserName: userName}).
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Category").
		Find(&orders)

	return orders
}

func (r *OrderRepository) UpdateOrder(c *model.Order) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
