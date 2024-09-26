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
}

func NewOrderService(orderRepo repository.Order, itemRepo repository.Item, sellerRepo repository.Seller) *OrderService {
	return &OrderService{
		orderRepo:  orderRepo,
		itemRepo:   itemRepo,
		sellerRepo: sellerRepo,
	}
}

func (s *OrderService) CreateOrder(order model.Order) error {
	total := 0.0

	for _, orderItem := range order.OrderItems {
		item, err := s.itemRepo.GetItemById(orderItem.ItemID)
		if err != nil {
			return err
		}

		if item.Quantity < orderItem.Quantity {
			return fmt.Errorf("not enough stock for item: %s", item.Name)
		}
	}

	for _, orderItem := range order.OrderItems {
		item, err := s.itemRepo.GetItemById(orderItem.ItemID)
		if err != nil {
			return err
		}

		orderItem.UnitPrice = item.Price
		orderItem.Total = item.Price * float64(orderItem.Quantity)
		total += orderItem.Total

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

	order.Total = total

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOrderById(orderID int) (model.Order, error) {
	return s.orderRepo.GetOrderById(orderID)
}
