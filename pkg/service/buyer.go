package service

import (
	"github.com/gofrs/uuid"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type BuyerService struct {
	repo     repository.Buyer
	itemRepo repository.Item
}

func NewBuyerService(repo repository.Buyer, itemRepo repository.Item) *BuyerService {
	return &BuyerService{repo: repo, itemRepo: itemRepo}
}

func (s *BuyerService) BuyerSignUp(buyerInput model.BuyerInput) error {
	buyer := model.Buyer{
		ID:       uuid.Must(uuid.NewV4()).String(),
		Name:     buyerInput.Name,
		Email:    buyerInput.Email,
		Password: buyerInput.Password,
	}
	return s.repo.BuyerSignUp(buyer)
}

func (s *BuyerService) GetBuyer(id string) (model.BuyerOutput, error) {

	buyer, err := s.repo.GetBuyer(id)
	if err != nil {
		return model.BuyerOutput{}, err
	}

	orders := []model.OrderOutput{}
	for _, order := range buyer.Orders {
		orders = append(orders, s.orderToOrderOutput(order))
	}

	favorites := model.ConvertItemsToItemInfo(buyer.Favorites)
	buyerOutput := model.BuyerOutput{
		ID:        buyer.ID,
		Name:      buyer.Name,
		Email:     buyer.Email,
		Favorites: favorites,
		Orders:    orders,
	}
	return buyerOutput, nil
}

func (s *BuyerService) UpdateBuyer(id string, buyer model.BuyerInput) error {
	updates := model.Buyer{
		ID:       id,
		Name:     buyer.Name,
		Email:    buyer.Email,
		Password: GeneratePasswordHash(buyer.Password),
	}
	return s.repo.UpdateBuyer(updates)
}

func (s *BuyerService) BuyerSignIn(mail, password string) (model.BuyerSignInResponse, error) {
	var signInResponse model.BuyerSignInResponse

	hashedPassword := GeneratePasswordHash(password)

	buyer, err := s.repo.BuyerSignIn(mail, hashedPassword)
	if err != nil {
		return model.BuyerSignInResponse{}, err
	}

	token := CreateToken(buyer.ID, "buyer")

	signInResponse.Token = token
	favorites := model.ConvertItemsToItemInfo(buyer.Favorites)

	orders := []model.OrderOutput{}
	for _, order := range buyer.Orders {
		orders = append(orders, s.orderToOrderOutput(order))
	}

	signInResponse.Buyer = model.BuyerOutput{
		ID:        buyer.ID,
		Name:      buyer.Name,
		Email:     buyer.Email,
		Favorites: favorites,
		Orders:    orders,
	}

	return signInResponse, err
}

func (s *BuyerService) AddToFavorites(buyerID string, itemID int) error {
	return s.repo.AddToFavorites(buyerID, itemID)
}

func (s *BuyerService) RemoveFromFavorites(buyerID string, itemID int) error {
	return s.repo.RemoveFromFavorites(buyerID, itemID)
}

func (s *BuyerService) orderToOrderOutput(order model.Order) model.OrderOutput {
	orderOutput := model.OrderOutput{}
	orderOutput.Total = order.Total
	orderOutput.Status = order.Status
	orderOutput.BuyerID = order.BuyerID
	for _, orderItem := range order.OrderItems {
		itemInfo, _ := s.itemRepo.GetItemById(orderItem.ID)

		orderOutput.Items = append(orderOutput.Items, model.OrderItemInfo{
			ID:       orderItem.ID,
			Name:     itemInfo.Name,
			Price:    itemInfo.Price,
			Quantity: itemInfo.Quantity,
		})
	}
	return orderOutput
}
