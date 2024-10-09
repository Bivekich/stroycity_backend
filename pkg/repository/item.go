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

func (r *ItemRepository) CreateItem(item model.Item) (int, error) {
	if err := r.db.Model(&model.Item{}).Create(&item).Error; err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r *ItemRepository) GetItemById(itemID int) (model.Item, error) {
	var item model.Item
	if err := r.db.Preload("Brand").Preload("Seller").Preload("Category").Preload("Material").Preload("Images").First(&item, itemID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *ItemRepository) UpdateItem(item model.Item) error {
	return r.db.Save(&item).Error
}

func (r *ItemRepository) GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs []uint, minPrice, maxPrice float64, query string) ([]model.Item, error) {
	var items []model.Item

	params := r.db.Model(&model.Item{}).Preload("Brand").Preload("Seller").Preload("Category").Preload("Material").Preload("Images")

	if len(brandIDs) > 0 {
		params = params.Where("brand_id IN ?", brandIDs)
	}

	if len(sellerIDs) > 0 {
		params = params.Where("seller_id IN ?", sellerIDs)
	}

	if len(categoryIDs) > 0 {
		params = params.Where("category_id IN ?", categoryIDs)
	}

	if len(materialIDs) > 0 {
		params = params.Where("material_id IN ?", materialIDs)
	}

	if minPrice > 0 {
		params = params.Where("price_with_discount >= ?", minPrice)
	}
	if maxPrice > 0 {
		params = params.Where("price_with_discount <= ?", maxPrice)
	}
	if len(query) > 0 {
		params = params.Where("name ILIKE ?", "%"+query+"%")
	}

	err := params.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) SaveImage(image model.Image) error {
	return r.db.Model(&model.Image{}).Create(&image).Error
}
