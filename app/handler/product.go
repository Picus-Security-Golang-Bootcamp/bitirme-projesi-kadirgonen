package handler

import (
	"net/http"
	"strconv"

	model "HW/app/models"
	"HW/app/service"
	"HW/pkg/logger"
	"HW/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type (
	Product struct {
		productService service.ProductService
		logger         logger.Logger
	}

	createProductRequest struct {
		Name       string  `json:"name" binding:"required"`
		StockCode  string  `json:"stock_code" binding:"required"`
		Cost       float64 `json:"cost" binding:"required"`
		Number     int     `json:"number" binding:"required"`
		CategoryID int     `json:"category_id" binding:"required"`
	}

	updateProductRequest struct {
		Name   string  `json:"name" binding:"required"`
		Cost   float64 `json:"cost" binding:"required"`
		Number int     `json:"number" binding:"required"`
	}

	productResponse struct {
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		StockCode    string  `json:"stock_code"`
		Cost         float64 `json:"cost"`
		Number       int     `json:"number"`
		CategoryID   int     `json:"category_id"`
		CategoryName string  `json:"category_name"`
	}

	searchResponse struct {
		Items []model.Product `json:"items"`
		Count int             `json:"count"`
	}
)

func NewProduct(cs service.ProductService, l logger.Logger) *Product {
	return &Product{cs, l}
}

func (c *Product) GetAllProducts(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	items, count := c.productService.GetAllProducts(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items

	g.JSON(http.StatusOK, paginatedResult)
}

func (c *Product) GetProduct(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	product := c.productService.GetProduct(int(id))
	if product == nil {
		errorResponse(g, http.StatusNotFound, "no record found")
		return
	}

	g.JSON(http.StatusOK, productResponse{
		ID: product.ID, Name: product.Name, StockCode: product.StockCode,
		Cost: product.Cost, Number: product.Number,
		CategoryName: product.Category.Name, CategoryID: product.Category.ID})
}

func (c *Product) SearchProducts(g *gin.Context) {
	searchQuery := g.Param("query")
	if len(searchQuery) == 0 {
		errorResponse(g, http.StatusBadRequest, "unable to get query")
		return
	}

	products := c.productService.SearchProducts(searchQuery)
	g.JSON(http.StatusOK, searchResponse{Items: products, Count: len(products)})
}

func (c *Product) CreateProduct(g *gin.Context) {
	var req createProductRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - createProduct")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	product := model.NewProduct(req.Name, req.StockCode, req.Cost, req.Number, req.CategoryID)
	err := c.productService.CreateProduct(product)
	if err != nil {
		c.logger.Error(err, "http - createProduct")
		errorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, productResponse{
		ID: product.ID, Name: product.Name, StockCode: product.StockCode,
		Cost: product.Cost, Number: product.Number,
		CategoryName: product.Category.Name, CategoryID: product.Category.ID})
}

func (c *Product) UpdateProduct(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorResponse(g, http.StatusBadRequest, "unable to get id")
		return
	}

	var req updateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - updateProduct")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	product := c.productService.GetProduct(int(id))
	product.Name = req.Name
	product.Cost = req.Cost
	product.Number = req.Number
	err = c.productService.UpdateProduct(product)
	if err != nil {
		c.logger.Error(err, "http - updateProduct")
		errorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, productResponse{
		ID: product.ID, Name: product.Name, StockCode: product.StockCode,
		Cost: product.Cost, Number: product.Number,
		CategoryName: product.Category.Name, CategoryID: product.Category.ID})
}

func (c *Product) DeleteProduct(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorResponse(g, http.StatusBadRequest, "unable to get id")
		return
	}

	err = c.productService.DeleteProduct(int(id))
	if err != nil {
		c.logger.Error(err, "http - deleteProduct")
		errorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(g, http.StatusOK, "Operation completed successfully.")
}
