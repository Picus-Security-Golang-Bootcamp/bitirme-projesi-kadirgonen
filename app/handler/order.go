package handler

import (
	"net/http"

	model "HW/app/models"
	"HW/app/service"
	"HW/pkg/logger"

	"github.com/gin-gonic/gin"
)

type (
	Order struct {
		orderService service.OrderService
		logger       logger.Logger
	}

	newOrderRequest struct {
		Name        string `json:"name" binding:"required"`
		Address     string `json:"address" binding:"required"`
		Phone       string `json:"phone" binding:"required"`
	}

	orderResponse struct {
		ID    string             `json:"id"`
		Items []*model.OrderItem `json:"items"`
	}
)

func NewOrder(os service.OrderService, l logger.Logger) *Order {
	return &Order{os, l}
}

func (c *Order) GetAllOrders(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	orders := c.orderService.GetAllOrders(userName)
	g.JSON(http.StatusOK, orders)
}

func (c *Order) CreateOrder(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	var req newOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - createOrder")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	err := c.orderService.CreateOrder(userName, req.Name, req.Address, req.PhoneNumber)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}

func (c *Order) CancelOrder(g *gin.Context) {
	userName := g.GetString("Email")
	if len(userName) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	id := g.Param("id")
	if len(id) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	err := c.orderService.CancelOrder(userName, id)
	if err != nil {
		errorResponse(g, http.StatusNotFound, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}
