package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddToCart(cartItem model.CartItem) error {
	return r.db.Create(&cartItem).Error
}

func (r *CartRepository) GetCartByBuyerID(buyerID string) (model.Cart, error) {
	var cart model.Cart
	err := r.db.Preload("CartItems.Item").Where("buyer_id = ?", buyerID).Find(&cart).Error
	return cart, err
}

func (r *CartRepository) UpdateCartItem(cartItem model.CartItem) error {
	return r.db.Save(&cartItem).Error
}

func (r *CartRepository) GetCartItemByID(cartItemID int) (model.CartItem, error) {
	var cartItem model.CartItem
	err := r.db.Where("id = ?", cartItemID).First(&cartItem).Error
	return cartItem, err
}

func (r *CartRepository) RemoveFromCart(cartItemID int) error {
	return r.db.Delete(&model.CartItem{}, cartItemID).Error
}

func (r *CartRepository) ClearCart(buyerID string) error {
	return r.db.Where("buyer_id = ?", buyerID).Delete(&model.CartItem{}).Error
}

func (r *CartRepository) GetCartItemsByBuyerID(buyerID string) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	if err := r.db.Where("buyer_id = ?", buyerID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
