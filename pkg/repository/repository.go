package repository

import (
	"gorm.io/gorm"
	"stroycity/pkg/model"
)

type Repository struct {
	Category
	Brand
	Material
	Seller
	Item
	Buyer
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		Brand:    NewBrandRepository(db),
		Material: NewMaterialRepository(db),
		Seller:   NewSellerRepository(db),
		Item:     NewItemRepository(db),
		Buyer:    NewBuyerRepository(db),
		Order:    NewOrderRepository(db),
	}
}

type Category interface {
	CreateCategory(category model.Category) error
	DeleteCategory(id int) error
	GetCategoryList() ([]model.Category, error)
}

type Brand interface {
	CreateBrand(brand model.Brand) error
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
	GetSeller(id string) (model.Seller, error)
	UpdateSeller(seller model.Seller) error
	SellerSignIn(mail, password string) (model.Seller, error)
}

type Item interface {
	CreateItem(item model.Item) error
	GetItemById(itemID int) (model.Item, error)
	UpdateItem(item model.Item) error
	GetItems(brandIDs, sellerIDs, categoryIDs, materialIDs []uint, minPrice, maxPrice float64) ([]model.Item, error)
	GetAllItems() ([]model.Item, error)
	SaveImage(image model.Image) error
}

type Buyer interface {
	BuyerSignUp(buyer model.Buyer) error
	GetBuyer(id string) (model.Buyer, error)
	UpdateBuyer(buyer model.Buyer) error
	BuyerSignIn(mail, password string) (model.Buyer, error)
}

type Order interface {
	CreateOrder(order model.Order) error
	GetOrderById(orderID int) (model.Order, error)
}
