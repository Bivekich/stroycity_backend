package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) CreateMaterial(material model.Material) error {
	return r.db.Create(&material).Error
}

func (r *MaterialRepository) DeleteMaterial(id int) error {
	return r.db.Delete(&model.Material{}, id).Error
}

func (r *MaterialRepository) GetMaterialList() ([]model.Material, error) {
	var materials []model.Material
	if err := r.db.Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}
