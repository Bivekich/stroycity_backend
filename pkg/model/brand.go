package model

type Brand struct {
	ID    int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Items []Item `json:"items" gorm:"foreignKey:BrandID"`
}
