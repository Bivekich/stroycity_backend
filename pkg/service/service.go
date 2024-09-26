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
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repos.Category),
		Brand:    NewBrandService(repos.Brand),
		Material: NewMaterialService(repos.Material),
		Seller:   NewSellerService(repos.Seller),
		Item:     NewItemService(repos.Item),
		Buyer:    NewBuyerService(repos.Buyer),
		Order:    NewOrderService(repos.Order, repos.Item, repos.Seller),
	}
}

type Category interface {
	CreateCategory(category model.Category) error
	DeleteCategory(id int) error
	GetCategoryList() ([]model.Category, error)
}

type Brand interface {
	CreateBrand(Brand model.Brand) error
	DeleteBrand(id int) error
	GetBrandList() ([]model.Brand, error)
}

type Material interface {
	CreateMaterial(material model.Material) error
	DeleteMaterial(id int) error
	GetMaterialList() ([]model.Material, error)
}

type Seller interface {
	SellerSignUp(seller model.Seller) error
	GetSeller(id string) (model.SellerOutput, error)
	UpdateSeller(id string, seller model.Seller) error
	SellerSignIn(mail, password string) (model.SellerSignInResponse, error)
}

type Item interface {
	CreateItem(material model.Item) error
	GetItemById(itemID int) (model.Item, error)
	UpdateItem(item model.Item) error
	GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs []uint, minPrice, maxPrice float64) ([]model.ItemInfo, error)
	GetAllItems() ([]model.ItemInfo, error)
	UploadImage(itemID int, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type Buyer interface {
	BuyerSignUp(buyer model.Buyer) error
	GetBuyer(id string) (model.BuyerOutput, error)
	UpdateBuyer(id string, buyer model.Buyer) error
	BuyerSignIn(mail, password string) (model.BuyerSignInResponse, error)
}

type Order interface {
	CreateOrder(order model.Order) error
	GetOrderById(orderID int) (model.Order, error)
}
