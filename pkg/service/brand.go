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

func (s *BrandService) GetBrandList() ([]model.Brand, error) {
	return s.repo.GetBrandList()
}
