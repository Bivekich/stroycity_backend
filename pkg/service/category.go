package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category model.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}

func (s *CategoryService) GetCategoryList() ([]model.Category, error) {
	return s.repo.GetCategoryList()
}
