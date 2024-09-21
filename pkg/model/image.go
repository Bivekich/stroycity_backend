package model

type Image struct {
	ID     int    `json:"id" gorm:"autoIncrement;primaryKey"`
	ItemID int    `json:"item_id" gorm:"not null"`
	URL    string `json:"url" gorm:"not null"`

	Item Item `gorm:"foreignKey:ItemID"`
}
