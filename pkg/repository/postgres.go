package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Item{},
		&model.Brand{},
		&model.Category{},
		&model.Material{},
		&model.Order{},
		&model.Seller{},
		&model.Buyer{},
		&model.OrderItem{},
		&model.Image{},
		&model.Admin{},
		&model.CartItem{},
		&model.Review{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
