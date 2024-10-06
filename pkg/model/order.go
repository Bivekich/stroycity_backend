package model

type Order struct {
	ID         int         `json:"id" gorm:"autoIncrement;primaryKey"`
	BuyerID    string      `json:"buyer_id"`
	Buyer      Buyer       `gorm:"foreignKey:BuyerID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	Total      float64     `json:"total" gorm:"not null"`
	Status     string      `json:"status" gorm:"not null"`
}

type OrderOutput struct {
	BuyerID string          `json:"buyer_id"`
	Total   float64         `json:"total" gorm:"not null"`
	Status  string          `json:"status" gorm:"not null"`
	Items   []OrderItemInfo `json:"items"`
}

type OrderItemInfo struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
