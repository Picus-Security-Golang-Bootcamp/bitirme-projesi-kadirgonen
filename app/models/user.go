package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint      `gorm:"primary_key;" json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Phone      string    `json:"phone"`
	Roles      []*Role   `gorm:"many2many:user_roles;" json:"roles"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}
type Role struct {
	gorm.Model
	Name       string    `gorm:"primary_key;" json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
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
