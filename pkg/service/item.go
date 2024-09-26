package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
	"time"
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

func (s *ItemService) GetItemById(itemID int) (model.Item, error) {
	return s.repo.GetItemById(itemID)
}

func (s *ItemService) UpdateItem(item model.Item) error {
	return s.repo.UpdateItem(item)
}

func (s *ItemService) GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs []uint, minPrice, maxPrice float64) ([]model.ItemInfo, error) {
	items, err := s.repo.GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}
	itemInfos := model.ConvertItemsToItemInfo(items)
	return itemInfos, nil
}

func (s *ItemService) GetAllItems() ([]model.ItemInfo, error) {
	items, err := s.repo.GetAllItems()
	if err != nil {
		return nil, err
	}
	itemInfos := model.ConvertItemsToItemInfo(items)
	return itemInfos, nil
}

func (s *ItemService) UploadImage(itemID int, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d_%d%s", itemID, time.Now().Unix(), ext)
	filePath := fmt.Sprintf("uploads/%s", fileName)

	if err := s.saveFile(file, filePath); err != nil {
		return "", err
	}

	image := model.Image{
		ItemID: itemID,
		URL:    filePath,
	}
	if err := s.repo.SaveImage(image); err != nil {
		return "", err
	}

	return filePath, nil
}

func (s *ItemService) saveFile(file multipart.File, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
