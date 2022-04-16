package service

import (
	"errors"
	"time"

	model "HW/app/models"
	"HW/app/repo"
)

type OrderService struct {
	orderRepo repo.OrderRepository
	cartRepo  repo.CartRepository
}

func NewOrderService(or repo.OrderRepository, cr repo.CartRepository) *OrderService {
	return &OrderService{
		orderRepo: or,
		cartRepo:  cr,
	}
}

func (s *OrderService) GetAllOrders(userName string) []model.Order {
	return s.orderRepo.GetAllOrders(userName)
}

func (s *OrderService) CreateOrder(userName string, name string, address string, phoneNumber string) error {
	order, _ := model.NewOrder(userName, name, address, phoneNumber)

	cart := s.cartRepo.GetCart(userName)
	if cart == nil {
		return errors.New("cart not found")
	}

	for _, v := range cart.Items {
		item, _ := model.NewOrderItem("", v.ProductID)
		if err := order.AddItem(item); err != nil {
			return errors.New("unable to add order item")
		}
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return errors.New("unable to create order")
	}

	if err := s.cartRepo.DeleteCart(cart.ID); err != nil {
		return errors.New("unable to clear basket")
	}

	return nil
}

func (s *OrderService) CancelOrder(userName string, orderId string) error {
	order := s.orderRepo.GetOrder(userName, orderId)
	if order == nil {
		return errors.New("order not found")
	}

	now := time.Now()
	orderDate := order.CreatedAt.In(time.UTC)
	days := now.Sub(orderDate).Hours() / 24
	if days > 14 {
		return errors.New("order cancel period ended")
	}

	order.Status = "canceled"
	if err := s.orderRepo.UpdateOrder(order); err != nil {
		return errors.New("unable to update order status")
	}

	return nil
}
