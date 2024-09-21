package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item model.Item) error {
	return s.repo.CreateItem(item)
}
