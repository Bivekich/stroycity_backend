package model

type Order struct {
	ID         int         `json:"id" gorm:"autoIncrement;primaryKey"`
	BuyerID    string      `json:"buyer_id"`
	Buyer      Buyer       `gorm:"foreignKey:BuyerID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	Total      float64     `json:"total" gorm:"not null"`
	Status     string      `json:"status" gorm:"not null"`
}
