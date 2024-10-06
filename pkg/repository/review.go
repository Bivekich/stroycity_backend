package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(review model.Review) error {
	return r.db.Model(&model.Review{}).Create(&review).Error
}

func (r *ReviewRepository) GetReviewsByItemID(itemID int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Model(&model.Review{}).Where("item_id = ?", itemID).Find(&reviews).Error
	return reviews, err
}
