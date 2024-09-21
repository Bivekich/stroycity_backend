package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateItem(item model.Item) error {
	return r.db.Model(&model.Item{}).Create(&item).Error
}

func (r *ItemRepository) GetItemById(itemID int) (model.Item, error) {
	var item model.Item
	if err := r.db.First(&item, itemID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *ItemRepository) UpdateItem(item model.Item) error {
	return r.db.Save(&item).Error
}
