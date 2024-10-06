package model

import "time"

type Review struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey"`
	ItemID    int       `json:"item_id" gorm:"not null"`
	BuyerID   string    `json:"buyer_id"  gorm:"not null"`
	Rating    float64   `json:"rating"  gorm:"not null"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}
