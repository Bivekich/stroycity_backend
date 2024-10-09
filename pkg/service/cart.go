package service

import (
	"errors"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type CartService struct {
	repo     repository.Cart
	itemRepo repository.Item
}

func NewCartService(repo repository.Cart, itemRepo repository.Item) *CartService {
	return &CartService{repo: repo, itemRepo: itemRepo}
}

func (s *CartService) AddToCart(buyerID string, itemID int, quantity int) error {
	usersCart, err := s.repo.GetCartByBuyerID(buyerID)
	if err != nil {
		return err
	}
	for _, cartItem := range usersCart.CartItems {
		if cartItem.ItemID == itemID {
			err = s.UpdateCartItem(cartItem.ID, quantity)
			if err != nil {
				return err
			}
			return nil
		}
	}
	cartItem := model.CartItem{
		BuyerID:  buyerID,
		ItemID:   itemID,
		Quantity: quantity,
	}

	return s.repo.AddToCart(cartItem)
}

func (s *CartService) GetCart(buyerID string) (model.CartOutput, error) {
	cart, err := s.repo.GetCartByBuyerID(buyerID)
	if err != nil {
		return model.CartOutput{}, err
	}
	cartOutput := model.CartOutput{}
	cartOutput.BuyerID = cart.BuyerID
	for _, cartItem := range cart.CartItems {
		itemInfo, _ := s.itemRepo.GetItemById(cartItem.ItemID)

		cartOutput.Items = append(cartOutput.Items, model.CartItemInfo{
			ID:       itemInfo.ID,
			Name:     itemInfo.Name,
			Price:    itemInfo.Price,
			Quantity: cartItem.Quantity,
		})
	}
	return cartOutput, nil
}

func (s *CartService) UpdateCartItem(cartItemID int, quantity int) error {
	cartItem, err := s.repo.GetCartItemByID(cartItemID)
	if err != nil {
		return err
	}

	cartItem.Quantity = quantity

	return s.repo.UpdateCartItem(cartItem)
}

func (s *CartService) RemoveFromCart(userID string, itemID int) error {
	cart, err := s.repo.GetCartByBuyerID(userID)
	if err != nil {
		return err
	}
	for _, cartItem := range cart.CartItems {
		if cartItem.ItemID == itemID {
			err = s.repo.RemoveFromCart(cartItem.ID)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Item not found")
}
