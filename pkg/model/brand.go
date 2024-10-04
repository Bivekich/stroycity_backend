package model

type Brand struct {
	ID    int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Items []Item `json:"items" gorm:"foreignKey:BrandID"`
}

type BrandInput struct {
	Name string `json:"name"`
}

type BrandOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
