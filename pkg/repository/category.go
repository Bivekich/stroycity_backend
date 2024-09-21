package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category model.Category) error {
	return r.db.Create(&category).Error
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	return r.db.Delete(&model.Category{}, id).Error
}

func (r *CategoryRepository) GetCategoryList() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
