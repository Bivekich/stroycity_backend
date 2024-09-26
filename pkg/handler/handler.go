package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"stroycity/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	category := router.Group("/category")
	{
		category.POST("", h.CreateCategory)
		category.GET("", h.GetCategoryList)
		category.DELETE("", h.DeleteCategory)
	}

	brand := router.Group("/brand")
	{
		brand.POST("", h.CreateBrand)
		brand.GET("", h.GetBrandList)
		brand.DELETE("", h.DeleteBrand)
	}

	material := router.Group("/material")
	{
		material.POST("", h.CreateMaterial)
		material.GET("", h.GetMaterialList)
		material.DELETE("", h.DeleteMaterial)
	}

	signUp := router.Group("/sign_up")
	{
		signUp.POST("/seller", h.SellerSignUp)
		signUp.POST("/buyer", h.BuyerSignUp)
	}

	signIn := router.Group("/sign_in")
	{
		signIn.POST("/seller", h.SellerSignIn)
		signIn.POST("/buyer", h.BuyerSignIn)
	}

	item := router.Group("/item")
	{
		item.POST("/create", h.CreateItem)
		item.POST("", h.GetItemList)
		item.GET("", h.GetItemById)
		item.PUT("", h.UpdateItem)
		item.POST("/image", h.UploadImage)
	}

	seller := router.Group("/seller")
	{
		seller.GET("", h.GetSeller)
		seller.PATCH("", h.UpdateSeller)
	}

	buyer := router.Group("/buyer")
	{
		buyer.GET("", h.GetBuyer)
		buyer.PATCH("", h.UpdateBuyer)
	}

	order := router.Group("/order")
	{
		order.GET("", h.GetOrder)
		order.POST("", h.CreateOrder)
	}

	return router
}
