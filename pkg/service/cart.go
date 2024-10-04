package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(buyerID string, itemID int, quantity int) error {
	cartItem := model.CartItem{
		BuyerID:  buyerID,
		ItemID:   itemID,
		Quantity: quantity,
	}

	return s.repo.AddToCart(cartItem)
}

func (s *CartService) GetCart(buyerID string) (model.Cart, error) {
	return s.repo.GetCartByBuyerID(buyerID)
}

func (s *CartService) UpdateCartItem(cartItemID int, quantity int) error {
	cartItem, err := s.repo.GetCartItemByID(cartItemID)
	if err != nil {
		return err
	}

	cartItem.Quantity = quantity

	return s.repo.UpdateCartItem(cartItem)
}

func (s *CartService) RemoveFromCart(cartItemID int) error {
	return s.repo.RemoveFromCart(cartItemID)
}
