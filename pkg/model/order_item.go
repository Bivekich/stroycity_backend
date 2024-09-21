package model

type OrderItem struct {
	ID        int     `json:"id" gorm:"autoIncrement;primaryKey"`
	OrderID   int     `json:"order_id" gorm:"not null"`
	ItemID    int     `json:"item_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	UnitPrice float64 `json:"unit_price" gorm:"not null"`
	Total     float64 `json:"total" gorm:"not null"`

	Item  Item  `gorm:"foreignKey:ItemID"`
	Order Order `gorm:"foreignKey:OrderID"`
}
