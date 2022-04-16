package service

import (
	"errors"

	model "HW/app/models"
	"HW/app/repo"
)

type CartService struct {
	cartRepo    repo.CartRepository
	productRepo repo.ProductRepository
}

func NewCartService(cr repo.CartRepository, pr repo.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cr,
		productRepo: pr,
	}
}

func (s *CartService) GetCart(userName string) *model.Cart {
	return s.cartRepo.GetCart(userName)
}

func (s *CartService) AddItemCart(userName string, productId int, number int) error {
	cart := s.cartRepo.GetCart(userName)
	if cart == nil {
		cart, _ := model.NewCart(userName)
		if err := s.cartRepo.CreatCart(cart); err != nil {
			return errors.New("unable to create basket")
		}
	}

	product := s.productRepo.GetProduct(productId)
	if product == nil {
		return errors.New("product not found")
	}

	if product.Number < number {
		return errors.New("not enough stock")
	}

	item, _ := model.NewCartItem(cart.ID, productId, number)
	if err := s.cartRepo.AddCartItem(item); err != nil {
		return errors.New("unable to add item to cart")
	}

	product.Number -= number
	if err := s.productRepo.UpdateProduct(product); err != nil {
		return errors.New("unable to update product stock info")
	}

	return nil
}

func (s *CartService) UpdateItemCart(userName string, productId int, number int) error {
	cart := s.cartRepo.GetCart(userName)
	if cart == nil {
		return errors.New("basket not found")
	}

	_, item := cart.SearchItemByProductId(productId)
	if item == nil {
		return errors.New("item not found")
	}

	product := s.productRepo.GetProduct(productId)
	if product == nil {
		return errors.New("product not found")
	}

	if product.Number < number {
		return errors.New("not enough stock")
	}

	product.Number = product.Number - (number - item.Number)
	if err := s.productRepo.UpdateProduct(product); err != nil {
		return errors.New("unable to update product stock info")
	}

	item.Number = number
	if err := s.cartRepo.UpdateCartItem(item); err != nil {
		return errors.New("unable to update item info")
	}

	return nil
}

func (s *CartService) RemoveItemCart(userName string, productId int) error {
	cart := s.cartRepo.GetCart(userName)
	if cart == nil {
		return errors.New("basket not found")
	}

	_, item := cart.SearchItemByProductId(productId)
	if item == nil {
		return errors.New("item not found")
	}

	product := s.productRepo.GetProduct(productId)
	if product == nil {
		return errors.New("product not found")
	}

	product.Number += item.Number
	if err := s.productRepo.UpdateProduct(product); err != nil {
		return errors.New("unable to update product stock info")
	}

	if err := s.cartRepo.DeleteCartItem(item); err != nil {
		return errors.New("unable to hard remove item")
	}

	return nil
}
