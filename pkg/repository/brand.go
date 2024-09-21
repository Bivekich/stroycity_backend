package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) CreateBrand(brand model.Brand) error {
	return r.db.Create(&brand).Error
}

func (r *BrandRepository) DeleteBrand(id int) error {
	return r.db.Delete(&model.Brand{}, id).Error
}

func (r *BrandRepository) GetBrandList() ([]model.Brand, error) {
	var brands []model.Brand
	if err := r.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}
