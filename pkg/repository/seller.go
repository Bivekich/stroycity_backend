package repository

import (
	"errors"
	"gorm.io/gorm"
	"stroycity/pkg/model"
	"time"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) SellerSignUp(seller model.Seller) error {
	return r.db.Model(&model.Seller{}).Create(&seller).Error
}

func (r *SellerRepository) GetSeller(id string) (model.Seller, error) {
	var seller model.Seller
	if err := r.db.Model(&model.Seller{}).Preload("Items").Where("id = ?", id).First(&seller).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func (r *SellerRepository) UpdateSeller(seller model.Seller) error {
	return r.db.Save(&seller).Error
}

func (r *SellerRepository) SellerSignIn(mail, password string) (model.Seller, error) {
	var seller model.Seller
	if err := r.db.Model(&model.Seller{}).Where("email = ?", mail).First(&seller).Error; err != nil {
		return seller, errors.New("Пользователя с такой почтой не существует!")
	}

	if err := r.db.Model(&model.Seller{}).Preload("Items").Where("email = ? AND password = ?", mail, password).First(&seller).Error; err != nil {
		return seller, errors.New("Неверный пароль!")
	}

	return seller, nil
}

func (r *SellerRepository) GetSellerEarnings(sellerID string) (float64, float64, error) {
	orderItems := []model.OrderItem{}
	if err := r.db.Model(&model.OrderItem{}).Where("seller_id = ?", sellerID).Find(&orderItems).Error; err != nil {
		return 0, 0, err
	}
	lastWeek, currentWeek := 0.0, 0.0
	currentTime := time.Now()
	for _, orderItem := range orderItems {
		if currentTime.Before(orderItem.CreatedAt) && currentTime.Add(-7*24*time.Hour).After(orderItem.CreatedAt) {
			currentWeek += orderItem.Total
		} else if currentTime.Add(-7*24*time.Hour).Before(orderItem.CreatedAt) && currentTime.Add(-14*24*time.Hour).After(orderItem.CreatedAt) {
			lastWeek += orderItem.Total
		}
	}
	return currentWeek, lastWeek, errors.New("")
}
