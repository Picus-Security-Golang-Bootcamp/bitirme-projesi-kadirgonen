package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint      `gorm:"primary_key;" json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Phone      string    `json:"phone"`
	Roles      []*Role   `gorm:"many2many:user_roles;" json:"roles"`
}
type Role struct {
	gorm.Model
	Name       string    `gorm:"primary_key;" json:"name"`
}

func NewUser(email string, password string, roles []*Role) *User {
	return &User{
		Email:    email,
		Password: password,
		Roles:    roles,
	}
}
func NewRole(name string) *Role {
	return &Role{
		Name: name,
	}
}
func (User) TableName() string {
	return "user"
}
func (Role) TableName() string {
	return "role"
}
