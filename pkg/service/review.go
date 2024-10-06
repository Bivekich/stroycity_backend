package service

import (
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func NewReviewService(repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(review model.Review) error {
	return s.repo.CreateReview(review)
}

func (s *ReviewService) GetReviewsByItemID(itemID int) ([]model.Review, error) {
	return s.repo.GetReviewsByItemID(itemID)
}
