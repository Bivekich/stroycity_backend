package service

import (
	"github.com/gofrs/uuid"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type BuyerService struct {
	repo repository.Buyer
}

func NewBuyerService(repo repository.Buyer) *BuyerService {
	return &BuyerService{repo: repo}
}

func (s *BuyerService) BuyerSignUp(buyer model.Buyer) error {
	buyer.ID = uuid.Must(uuid.NewV4()).String()
	buyer.Password = GeneratePasswordHash(buyer.Password)
	return s.repo.BuyerSignUp(buyer)
}

func (s *BuyerService) GetBuyer(id string) (model.BuyerOutput, error) {

	buyer, err := s.repo.GetBuyer(id)
	if err != nil {
		return model.BuyerOutput{}, err
	}

	favorites := model.ConvertItemsToItemInfo(buyer.Favorites)
	buyerOutput := model.BuyerOutput{
		ID:        buyer.ID,
		Name:      buyer.Name,
		Email:     buyer.Email,
		Favorites: favorites,
		Orders:    buyer.Orders,
	}
	return buyerOutput, nil
}

func (s *BuyerService) UpdateBuyer(id string, buyer model.Buyer) error {
	buyer.ID = id
	return s.repo.UpdateBuyer(buyer)
}

func (s *BuyerService) BuyerSignIn(mail, password string) (model.BuyerSignInResponse, error) {
	var signInResponse model.BuyerSignInResponse

	hashedPassword := GeneratePasswordHash(password)

	buyer, err := s.repo.BuyerSignIn(mail, hashedPassword)
	if err != nil {
		return model.BuyerSignInResponse{}, err
	}

	token := CreateToken(buyer.ID)

	signInResponse.Token = token
	favorites := model.ConvertItemsToItemInfo(buyer.Favorites)
	signInResponse.Buyer = model.BuyerOutput{
		ID:        buyer.ID,
		Name:      buyer.Name,
		Email:     buyer.Email,
		Favorites: favorites,
		Orders:    buyer.Orders,
	}

	return signInResponse, err
}
