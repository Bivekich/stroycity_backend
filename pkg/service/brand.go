package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type BrandService struct {
	repo repository.Brand
}

func NewBrandService(repo repository.Brand) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) CreateBrand(brand model.Brand) error {
	return s.repo.CreateBrand(brand)
}

func (s *BrandService) DeleteBrand(id int) error {
	return s.repo.DeleteBrand(id)
}

func (s *BrandService) GetBrandList() ([]model.BrandOutput, error) {
	brands, err := s.repo.GetBrandList()
	if err != nil {
		return nil, err
	}
	var brandsInfo []model.BrandOutput
	for _, brand := range brands {
		brandsInfo = append(brandsInfo, model.BrandOutput{
			ID:   brand.ID,
			Name: brand.Name,
		})
	}
	return brandsInfo, nil
}
