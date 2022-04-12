package repo

import (
	"errors"

	model "HW/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]model.User, int) {
	var users []model.User
	var count int64

	r.db.Find(&users).Count(&count)

	return users, int(count)
}

func (r *UserRepository) GetUser(id int) *model.User {
	var user model.User
	result := r.db.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (r *UserRepository) GetUserIdWithRoles(id int) *model.User {
	var user model.User
	result := r.db.Preload("Roles").First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (r *UserRepository) GetUserEmail(email string) *model.User {
	var user model.User
	result := r.db.Where(model.User{Email: email}).Attrs(model.User{}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (r *UserRepository) GetUserEmailPassword(email string, password string) *model.User {
	var user model.User
	result := r.db.Where(model.User{Email: email, Password: password}).Attrs(model.User{}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (r *UserRepository) CreateUser(c *model.User) error {
	result := r.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) UpdateUser(c *model.User) error {
	result := r.db.Save(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	result := r.db.Delete(&model.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
