package model

type Material struct {
	ID    int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Items []Item `json:"items" gorm:"foreignKey:MaterialID"`
}

type MaterialInput struct {
	Name string `json:"name"`
}

type MaterialOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
