package model

type FilterRequest struct {
	BrandIDs    []uint  `json:"brands"`
	SellerIDs   []uint  `json:"sellers"`
	CategoryIDs []uint  `json:"categories"`
	MaterialIDs []uint  `json:"materials"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
}
