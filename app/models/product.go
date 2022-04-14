package repo

import (
	"errors"

	model "HW/app/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAllProducts(pageIndex, pageSize int) ([]model.Product, int) {
	var products []model.Product
	var count int64

	r.db.Preload("Category").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

func (r *ProductRepository) GetAllProductsCategoryId(categoryId int) ([]model.Product, int) {
	var products []model.Product
	var count int64
	r.db.Preload("Category").Find(&products).Where(&model.Product{CategoryID: categoryId}).Count(&count)

	return products, int(count)
}

func (r *ProductRepository) GetProduct(id int) *model.Product {
	var product model.Product
	result := r.db.Preload("Category").First(&product, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &product
}

func (r *ProductRepository) GetProductStockCode(stockCode string) *model.Product {
	var product model.Product
	result := r.db.Preload("Category").Where(&model.Product{StockCode: stockCode}).First(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &product
}

func (r *ProductRepository) SearchProduct(productname string) []model.Product {
	var products []model.Product
	r.db.Preload("Category").Where("name LIKE ? OR stockCode LIKE ?", "%"+productname+"%", "%"+productname+"%").Find(&products)

	return products
}

func (r *ProductRepository) CreateProduct(c *model.Product) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) UpdateProduct(Product *model.Product) error {
	result := r.db.Save(&Product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	result := r.db.Delete(&model.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
