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
	Category struct {
		categoryService service.CategoryService
		logger          logger.Logger
	}

	createCategoryRequest struct {
		Name string `json:"name" binding:"required"`
	}

	categoryResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	createBulkCategoryResponse struct {
		Added    int `json:"added_count"`
		Existing int `json:"existing_count"`
	}
)

func NewCategory(cs service.CategoryService, l logger.Logger) *Category {
	return &Category{cs, l}
}

func (c *Category) GetAllCategories(g *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(g)
	items, count := c.categoryService.GetAllCategories(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(g, count)
	paginatedResult.Items = items

	g.JSON(http.StatusOK, paginatedResult)
}

func (c *Category) GetCategory(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		errorResponse(g, http.StatusBadRequest, "unable to get parameters")
		return
	}

	category := c.categoryService.GetCategory(int(id))
	if category == nil {
		errorResponse(g, http.StatusNotFound, "no record found")
		return
	}

	g.JSON(http.StatusOK, categoryResponse{ID: category.ID, Name: category.Name})
}

func (c *Category) CreateCategory(g *gin.Context) {
	var req createCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		c.logger.Error(err, "http - createCategory")
		errorResponse(g, http.StatusBadRequest, "invalid request body")
		return
	}

	category := model.NewCategory(strconv.IntSize, req.Name)
	err := c.categoryService.CreateCategory(category)
	if err != nil {
		c.logger.Error(err, "http - createCategory")
		errorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, categoryResponse{ID: category.ID, Name: category.Name})
}

func (c *Category) CreateBulkCategory(g *gin.Context) {

	file, fileHead, err := g.Request.FormFile("file")
	if err != nil {
		c.logger.Error(err, "http - createBulkCategory")
		errorResponse(g, http.StatusBadRequest, err.Error())

		return
	}

	contentType := fileHead.Header.Values("Content-Type")
	if contentType[0] != "text/csv" {
		errorResponse(g, http.StatusBadRequest, "invalid file type")
	}

	added, existing, err := c.categoryService.CreateBulkCategory(file)
	if err != nil {
		errorResponse(g, http.StatusInternalServerError, err.Error())
	}

	g.JSON(http.StatusOK, createBulkCategoryResponse{Added: added, Existing: existing})
}

func (c *Category) DeleteCategory(g *gin.Context) {
	c.logger.Fatal("unimplemented exception")
}
