package handler

import (
	"net/http"

	model "HW/app/models"
	"HW/app/service"
	"HW/pkg/logger"

	"github.com/gin-gonic/gin"
)

type (
	Cart struct {
		cartService service.CartService
		logger      logger.Logger
	}

	newCartItemRequest struct {
		ProductId int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	removeCartItemRequest struct {
		ProductId int `json:"product_id"`
	}

	cartResponse struct {
		ID    string            `json:"id"`
		Items []*model.CartItem `json:"items"`
	}
)

func NewCart(cs service.CartService, l logger.Logger) *Cart {
	return &Cart{cs, l}
}
func (c *Cart) GetCart(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	cart := c.cartService.GetCart(userName)
	if cart == nil {
		errorResponse(g, http.StatusNotFound, "no record found")
		return
	}

	g.JSON(http.StatusOK, cartResponse{ID: cart.ID, Items: cart.Items})
}
func (c *Cart) AddCartItem(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	var req newCartItemRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - addCartItem")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	err := c.cartService.AddItemCart(userName, req.ProductId, req.Quantity)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}
func (c *Cart) UpdateCartItem(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	var req newCartItemRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - updateBasketItem")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	err := c.cartService.UpdateItemCart(userName, req.ProductId, req.Quantity)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}
func (c *Cart) RemoveCartItem(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	var req removeCartItemRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - removeBasketItem")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	err := c.cartService.RemoveItemCart(userName, req.ProductId)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}
