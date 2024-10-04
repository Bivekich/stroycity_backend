package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type MaterialService struct {
	repo repository.Material
}

func NewMaterialService(repo repository.Material) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) CreateMaterial(material model.Material) error {
	return s.repo.CreateMaterial(material)
}

func (s *MaterialService) DeleteMaterial(id int) error {
	return s.repo.DeleteMaterial(id)
}

func (s *MaterialService) GetMaterialList() ([]model.MaterialOutput, error) {
	materials, err := s.repo.GetMaterialList()
	if err != nil {
		return nil, err
	}
	var materialsInfo []model.MaterialOutput
	for _, material := range materials {
		materialsInfo = append(materialsInfo, model.MaterialOutput{
			ID:   material.ID,
			Name: material.Name,
		})
	}
	return materialsInfo, nil
}
