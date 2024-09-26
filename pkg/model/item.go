package model

type Item struct {
	ID                int     `json:"id" gorm:"autoIncrement;primaryKey"`
	Name              string  `json:"name" gorm:"not null"`
	Description       string  `json:"description" gorm:"not null"`
	Article           string  `json:"article"`
	Price             float64 `json:"price" gorm:"not null"`
	PriceWithDiscount float64 `json:"price_with_discount" gorm:"not null"`
	Quantity          int     `json:"quantity" gorm:"default:0"`
	Length            int     `json:"length" gorm:"default:0"`
	Width             int     `json:"width" gorm:"default:0"`
	Height            int     `json:"height" gorm:"default:0"`
	Weight            int     `json:"weight" gorm:"default:0"`
	CategoryID        int     `json:"category_id" gorm:"not null"`
	BrandID           int     `json:"brand_id" gorm:"not null"`
	SellerID          string  `json:"seller_id" gorm:"not null"`
	MaterialID        int     `json:"material_id" gorm:"not null"`

	Category    Category `gorm:"foreignKey:CategoryID"`
	Brand       Brand    `gorm:"foreignKey:BrandID"`
	Seller      Seller   `gorm:"foreignKey:SellerID"`
	Material    Material `gorm:"foreignKey:MaterialID"`
	FavoritedBy []Buyer  `gorm:"many2many:buyer_favorites"`
	Images      []Image  `gorm:"foreignKey:ItemID"`
}

type ItemInfo struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type CurrentItemInfo struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Article           string   `json:"article"`
	Price             float64  `json:"price"`
	PriceWithDiscount float64  `json:"price_with_discount"`
	Quantity          int      `json:"quantity"`
	Length            int      `json:"length"`
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	Weight            int      `json:"weight"`
	Category          string   `json:"category"`
	Brand             string   `json:"brand"`
	Seller            string   `json:"seller"`
	Material          string   `json:"material"`
	Images            []string `json:"images"`
}

func ConvertItemsToItemInfo(items []Item) []ItemInfo {
	var itemInfos []ItemInfo

	for _, item := range items {
		itemInfo := ItemInfo{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		}
		itemInfos = append(itemInfos, itemInfo)
	}

	return itemInfos
}
