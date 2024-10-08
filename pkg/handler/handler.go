package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "stroycity/docs"
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

	//SWAGGER
	////////////////////////////////////////////////////////////
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	////////////////////////////////////////////////////////////

	// EVERYONE
	////////////////////////////////////////////////////////////
	router.GET("/category", h.GetCategoryList)
	router.GET("/brand", h.GetBrandList)
	router.GET("/material", h.GetMaterialList)
	router.GET("/review", h.GetReviews)

	item := router.Group("/item")
	{
		item.POST("", h.GetItemList)
		item.GET("", h.GetItemById)
	}
	////////////////////////////////////////////////////////////

	//SIGN UP
	////////////////////////////////////////////////////////////
	signUp := router.Group("/sign-up")
	{
		signUp.POST("/seller", h.SellerSignUp)
		signUp.POST("/buyer", h.BuyerSignUp)
	}
	////////////////////////////////////////////////////////////

	//SIGN IN
	////////////////////////////////////////////////////////////
	signIn := router.Group("/sign-in")
	{
		signIn.POST("/seller", h.SellerSignIn)
		signIn.POST("/buyer", h.BuyerSignIn)
		signIn.POST("/admin", h.AdminSignIn)
	}
	////////////////////////////////////////////////////////////

	//ADMIN
	////////////////////////////////////////////////////////////
	admin := router.Group("/admin", h.UserIdentity)
	{
		admin.POST("/sign-up", h.AdminSignUp)

		category := admin.Group("/category")
		{
			category.POST("", h.CreateCategory)
			category.DELETE("", h.DeleteCategory)
		}

		brand := admin.Group("/brand")
		{
			brand.POST("", h.CreateBrand)
			brand.DELETE("", h.DeleteBrand)
		}

		material := admin.Group("/material")
		{
			material.POST("", h.CreateMaterial)
			material.DELETE("", h.DeleteMaterial)
		}
	}
	////////////////////////////////////////////////////////////

	//SELLER
	////////////////////////////////////////////////////////////
	seller := router.Group("/seller", h.UserIdentity)
	{
		seller.GET("", h.GetSeller)
		seller.PATCH("", h.UpdateSeller)

		sellerItem := seller.Group("/item")
		{
			sellerItem.POST("", h.CreateItem)
			sellerItem.PUT("", h.UpdateItem)
			sellerItem.POST("/image", h.UploadImage)
		}

		seller.GET("/statistic", h.GetSellerEarnings)

	}
	////////////////////////////////////////////////////////////

	//BUYER
	////////////////////////////////////////////////////////////
	buyer := router.Group("/buyer", h.UserIdentity)
	{
		buyer.GET("", h.GetBuyer)
		buyer.PATCH("", h.UpdateBuyer)

		order := buyer.Group("/order")
		{
			order.GET("", h.GetOrder)
			order.POST("", h.CreateOrder)
		}

		cart := buyer.Group("/cart")
		{
			cart.GET("", h.GetCart)
			cart.POST("", h.AddToCart)
			cart.DELETE("", h.RemoveFromCart)
		}

		favoritest := buyer.Group("/favorites")
		{
			favoritest.POST("", h.AddToFavorites)
			favoritest.DELETE("", h.RemoveFromFavorites)
		}

		buyer.POST("/review", h.CreateReview)
	}
	////////////////////////////////////////////////////////////

	return router
}
