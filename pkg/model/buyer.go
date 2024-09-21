package model

type Buyer struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"not null"`
	Email     string  `json:"email" gorm:"not null;unique"`
	Password  string  `json:"password" gorm:"not null"`
	Orders    []Order `json:"orders" gorm:"foreignKey:BuyerID"`
	Favorites []Item  `json:"favorites" gorm:"many2many:buyer_favorites"`
}

type BuyerOutput struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Orders    []Order    `json:"orders"`
	Favorites []ItemInfo `json:"favorites"`
}

type BuyerSignInResponse struct {
	Token string      `json:"token"`
	Buyer BuyerOutput `json:"buyer"`
}
