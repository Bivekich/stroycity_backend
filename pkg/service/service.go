package service

import (
	"mime/multipart"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type Service struct {
	Category
	Brand
	Material
	Seller
	Item
	Buyer
	Order
	Admin
	Cart
	Review
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repos.Category),
		Brand:    NewBrandService(repos.Brand),
		Material: NewMaterialService(repos.Material),
		Seller:   NewSellerService(repos.Seller),
		Item:     NewItemService(repos.Item),
		Buyer:    NewBuyerService(repos.Buyer, repos.Item),
		Order:    NewOrderService(repos.Order, repos.Item, repos.Seller, repos.Cart),
		Admin:    NewAdminService(repos.Admin),
		Cart:     NewCartService(repos.Cart, repos.Item),
		Review:   NewReviewService(repos.Review),
	}
}

type Category interface {
	CreateCategory(category model.Category) error
	DeleteCategory(id int) error
	GetCategoryList() ([]model.CategoryOutput, error)
}

type Brand interface {
	CreateBrand(Brand model.Brand) error
	DeleteBrand(id int) error
	GetBrandList() ([]model.BrandOutput, error)
}

type Material interface {
	CreateMaterial(material model.Material) error
	DeleteMaterial(id int) error
	GetMaterialList() ([]model.MaterialOutput, error)
}

type Seller interface {
	SellerSignUp(seller model.Seller) error
	GetSeller(id string) (model.SellerOutput, error)
	UpdateSeller(id string, seller model.Seller) error
	SellerSignIn(mail, password string) (model.SellerSignInResponse, error)
	GetSellerEarnings(sellerID string) (float64, float64, error)
}

type Item interface {
	CreateItem(material model.Item) error
	GetItemById(itemID int) (model.CurrentItemInfo, error)
	UpdateItem(item model.Item) error
	GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs []uint, minPrice, maxPrice float64) ([]model.ItemInfo, error)
	GetAllItems() ([]model.ItemInfo, error)
	UploadImage(itemID int, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type Buyer interface {
	BuyerSignUp(buyer model.BuyerInput) error
	GetBuyer(id string) (model.BuyerOutput, error)
	UpdateBuyer(id string, buyer model.BuyerInput) error
	BuyerSignIn(mail, password string) (model.BuyerSignInResponse, error)
}

type Order interface {
	CreateOrder(buyerID string) error
	GetOrderById(orderID int) (model.OrderOutput, error)
	ClearCart(buyerID string) error
}

type Admin interface {
	AdminSignUp(admin model.Admin) error
	AdminSignIn(login, password string) (string, error)
}

type Cart interface {
	AddToCart(buyerID string, itemID int, quantity int) error
	GetCart(buyerID string) (model.CartOutput, error)
	UpdateCartItem(cartItemID int, quantity int) error
	RemoveFromCart(cartItemID int) error
}

type Review interface {
	CreateReview(review model.Review) error
	GetReviewsByItemID(itemID int) ([]model.Review, error)
}
