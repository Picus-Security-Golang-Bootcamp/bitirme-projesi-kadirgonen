package repo

import (
	"errors"

	model "HW/app/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAllCategories(pageIndex, pageSize int) ([]model.Category, int) {
	var categories []model.Category
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)

	return categories, int(count)
}

func (r *CategoryRepository) GetCategory(id int) *model.Category {
	var product model.Category
	result := r.db.First(&product, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &product
}

func (r *CategoryRepository) GetCategoryByName(name string) *model.Category {
	var category model.Category
	result := r.db.Where(model.Category{Name: name}).Attrs(model.Category{}).First(&category)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &category
}

func (r *CategoryRepository) CreateCategory(Category *model.Category) error {
	result := r.db.Create(&Category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(Category model.Category) error {
	result := r.db.Save(&Category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	result := r.db.Delete(&model.Category{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
