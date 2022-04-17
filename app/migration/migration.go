package migration

import (
	"fmt"

	model "HW/app/models"

	"gorm.io/gorm"
)

var roles = []model.Role{
	{Name: "admin"},
	{Name: "customer"},
}

var users = []model.User{
        {Name: "kadir", Email: "kg@kgstore.com", Password: "kg1453", Phone: "05436680353", Roles: []*model.Role{&roles[0]}},
}

var categories = []model.Category{
	{Name: "Bilgisayar", ID: 1},
}

func Execute(db *gorm.DB) error {
	// Check if migration done
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return fmt.Errorf("seeder - Load - db.Migrator.GetTables: %w", err)
	}

	if len(tables) == 0 {
		// Auto create tables
		err = db.AutoMigrate(
			&model.User{},
			&model.Role{},
			&model.Cart{},
			&model.CartItem{},
			&model.Category{},
			&model.Product{},
			&model.Order{},
			&model.OrderItem{},
		)

		if err != nil {
			return fmt.Errorf("seeder - Load - db.AutoMigrate: %w", err)
		}

		for i := range categories {
			err := db.Model(&model.Category{}).Create(&categories[i]).Error
			if err != nil {
				return fmt.Errorf("seeder - Load - Model(Category).Create: %w", err)
			}
		}

		for i := range roles {
			err := db.Model(&model.Role{}).Create(&roles[i]).Error
			if err != nil {
				return fmt.Errorf("seeder - Load - Model(Role).Create: %w", err)
			}
		}

		for i := range users {
			err := db.Model(&model.User{}).Create(&users[i]).Error
			if err != nil {
				return fmt.Errorf("seeder - Load - Model(User).Create: %w", err)
			}
		}
	}

	return nil
}
