package model

type CartItem struct {
	ID       int    `json:"id" gorm:"autoIncrement;primaryKey"`
	BuyerID  string `json:"buyer_id" gorm:"not null"`
	ItemID   int    `json:"item_id" gorm:"not null"`
	Quantity int    `json:"quantity" gorm:"not null"`
	Item     Item   `gorm:"foreignKey:ItemID"`
	Buyer    Buyer  `gorm:"foreignKey:BuyerID"`
}

type Cart struct {
	BuyerID   string     `json:"buyer_id" gorm:"primaryKey"`
	Buyer     Buyer      `gorm:"foreignKey:BuyerID"`
	CartItems []CartItem `json:"cart_items" gorm:"foreignKey:BuyerID"`
}

type CartOutput struct {
	BuyerID string         `json:"buyer_id"`
	Items   []CartItemInfo `json:"items"`
}

type CartItemInfo struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddToCartInput struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
