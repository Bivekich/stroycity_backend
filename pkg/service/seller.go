package service

import (
	"github.com/gofrs/uuid"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type SellerService struct {
	repo repository.Seller
}

func NewSellerService(repo repository.Seller) *SellerService {
	return &SellerService{repo: repo}
}

func (s *SellerService) SellerSignUp(seller model.Seller) error {
	seller.ID = uuid.Must(uuid.NewV4()).String()
	seller.Password = GeneratePasswordHash(seller.Password)
	return s.repo.SellerSignUp(seller)
}

func (s *SellerService) GetSeller(id string) (model.SellerOutput, error) {

	seller, err := s.repo.GetSeller(id)
	if err != nil {
		return model.SellerOutput{}, err
	}

	items := model.ConvertItemsToItemInfo(seller.Items)
	sellerOutput := model.SellerOutput{
		ID:       seller.ID,
		Name:     seller.Name,
		Email:    seller.Email,
		ShopName: seller.ShopName,
		Items:    items,
	}
	return sellerOutput, nil
}

func (s *SellerService) UpdateSeller(id string, seller model.Seller) error {
	seller.ID = id
	return s.repo.UpdateSeller(seller)
}

func (s *SellerService) SellerSignIn(mail, password string) (model.SellerSignInResponse, error) {
	var signInResponse model.SellerSignInResponse

	hashedPassword := GeneratePasswordHash(password)

	seller, err := s.repo.SellerSignIn(mail, hashedPassword)
	if err != nil {
		return model.SellerSignInResponse{}, err
	}

	token := CreateToken(seller.ID, "seller")

	items := model.ConvertItemsToItemInfo(seller.Items)

	signInResponse.Token = token
	signInResponse.Seller = model.SellerOutput{
		ID:       seller.ID,
		Name:     seller.Name,
		Email:    seller.Email,
		ShopName: seller.ShopName,
		Items:    items,
	}

	return signInResponse, err
}
