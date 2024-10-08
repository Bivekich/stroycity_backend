package repository

import (
	"errors"
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type BuyerRepository struct {
	db *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}

func (r *BuyerRepository) BuyerSignUp(buyer model.Buyer) error {
	return r.db.Model(&model.Buyer{}).Create(&buyer).Error
}

func (r *BuyerRepository) GetBuyer(id string) (model.Buyer, error) {
	var buyer model.Buyer
	if err := r.db.Model(&model.Buyer{}).Preload("Orders").Preload("Favorites").Where("id = ?", id).First(&buyer).Error; err != nil {
		return buyer, err
	}
	return buyer, nil
}

func (r *BuyerRepository) UpdateBuyer(buyer model.Buyer) error {
	return r.db.Save(&buyer).Error
}

func (r *BuyerRepository) BuyerSignIn(mail, password string) (model.Buyer, error) {
	var buyer model.Buyer
	if err := r.db.Model(&model.Buyer{}).Where("email = ?", mail).First(&buyer).Error; err != nil {
		return buyer, errors.New("Пользователя с такой почтой не существует!")
	}

	if err := r.db.Model(&model.Buyer{}).Preload("Orders").Preload("Favorites").Where("email = ? AND password = ?", mail, password).First(&buyer).Error; err != nil {
		return buyer, errors.New("Неверный пароль!")
	}

	return buyer, nil
}

func (r *BuyerRepository) AddToFavorites(buyerID string, itemID int) error {
	var buyer model.Buyer
	if err := r.db.Preload("Favorites").First(&buyer, "id = ?", buyerID).Error; err != nil {
		return err
	}

	// Проверка, что товара нет в избранном
	for _, item := range buyer.Favorites {
		if item.ID == itemID {
			return nil // Товар уже в избранном
		}
	}

	// Добавляем товар в избранное
	return r.db.Model(&buyer).Association("Favorites").Append(&model.Item{ID: itemID})
}

// RemoveFromFavorites удаляет товар из избранного покупателя
func (r *BuyerRepository) RemoveFromFavorites(buyerID string, itemID int) error {
	var buyer model.Buyer
	if err := r.db.Preload("Favorites").First(&buyer, "id = ?", buyerID).Error; err != nil {
		return err
	}

	// Удаляем товар из избранного
	return r.db.Model(&buyer).Association("Favorites").Delete(&model.Item{ID: itemID})
}
