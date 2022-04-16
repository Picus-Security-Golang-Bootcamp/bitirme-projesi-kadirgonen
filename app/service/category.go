package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	model "HW/app/models"
	"HW/app/repo"
)

type CategoryService struct {
	categoryRepo repo.CategoryRepository
}

func NewProductService(cr repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: cr,
	}
}

func (s *CategoryService) GetAllCategories(pageIndex, pageSize int) ([]model.Category, int) {

	return s.categoryRepo.GetAllCategories(pageIndex, pageSize)
}

func (s *CategoryService) GetCategory(categoryId int) *model.Category {
	return s.categoryRepo.GetCategory(categoryId)
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	categoryExists := s.categoryRepo.GetCategoryByName(category.Name)
	if categoryExists != nil {
		return errors.New("category with same name already exist in database")
	}

	err := s.categoryRepo.CreateCategory(category)
	if err != nil {
		return errors.New("an unknown error occurred during operation")
	}

	return nil
}

func (s *CategoryService) CreateBulkCategory(file multipart.File) (int, int, error) {
	reader := csv.NewReader(file)

	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return 0, 0, errors.New("unable to initialize csv reader")
	}

	addedCount := 0
	existingCount := 0

	for _, line := range lines[1:] {
		id, _ := strconv.Atoi(line[0])
		name := line[1]

		category := model.NewCategory(
			id,   // ID
			name, // Name
		)

		categoryExists := s.categoryRepo.GetCategoryByName(name)
		if categoryExists != nil {
			existingCount += 1
		} else {
			if err := s.categoryRepo.CreateCategory(category); err != nil {
				return addedCount, existingCount, fmt.Errorf("unable to create category: %s", name)
			}
			addedCount += 1
		}
	}

	return addedCount, existingCount, nil
}
