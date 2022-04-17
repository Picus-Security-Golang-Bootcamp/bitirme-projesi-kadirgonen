package database_handler

import (
	"time"

	"HW/config"

	"go.uber.org/zap"
	gormPsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {

	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER_NAME"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open(gormPsql.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot connect to database", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("Cannot get sql.DB from database", zap.Error(err))
	}

	origin.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)

	return db
}
