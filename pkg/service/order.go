package service

import (
	"fmt"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type OrderService struct {
	orderRepo  repository.Order
	itemRepo   repository.Item
	sellerRepo repository.Seller
	cartRepo   repository.Cart
}

func NewOrderService(orderRepo repository.Order, itemRepo repository.Item, sellerRepo repository.Seller, cartRepo repository.Cart) *OrderService {
	return &OrderService{
		orderRepo:  orderRepo,
		itemRepo:   itemRepo,
		sellerRepo: sellerRepo,
		cartRepo:   cartRepo,
	}
}

func (s *OrderService) CreateOrder(buyerID string) error {
	// Получаем товары из корзины покупателя
	cartItems, err := s.cartRepo.GetCartItemsByBuyerID(buyerID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	total := 0.0
	orderItems := []model.OrderItem{}

	// Проверяем наличие товаров на складе
	for _, cartItem := range cartItems {
		item, err := s.itemRepo.GetItemById(cartItem.ItemID)
		if err != nil {
			return err
		}

		if item.Quantity < cartItem.Quantity {
			return fmt.Errorf("not enough stock for item: %s", item.Name)
		}

		// Добавляем в список товаров для заказа
		orderItem := model.OrderItem{
			ItemID:    cartItem.ItemID,
			Quantity:  cartItem.Quantity,
			UnitPrice: item.Price,
			Total:     item.Price * float64(cartItem.Quantity),
		}
		orderItems = append(orderItems, orderItem)
		total += orderItem.Total
	}

	// Создаем заказ
	order := model.Order{
		BuyerID:    buyerID,
		OrderItems: orderItems,
		Total:      total,
	}

	// Сохраняем заказ
	if err := s.orderRepo.CreateOrder(order); err != nil {
		return err
	}

	// Обновляем количество товара и баланс продавца
	for _, orderItem := range order.OrderItems {
		item, err := s.itemRepo.GetItemById(orderItem.ItemID)
		if err != nil {
			return err
		}

		item.Quantity -= orderItem.Quantity
		if err := s.itemRepo.UpdateItem(item); err != nil {
			return err
		}

		seller, err := s.sellerRepo.GetSeller(item.SellerID)
		if err != nil {
			return err
		}
		seller.Balance += orderItem.Total
		if err := s.sellerRepo.UpdateSeller(seller); err != nil {
			return err
		}
	}

	// Очищаем корзину
	if err := s.cartRepo.ClearCart(buyerID); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOrderById(orderID int) (model.Order, error) {
	return s.orderRepo.GetOrderById(orderID)
}

func (s *OrderService) ClearCart(buyerID string) error {
	return s.cartRepo.ClearCart(buyerID)
}
