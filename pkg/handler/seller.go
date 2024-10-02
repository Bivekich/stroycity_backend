package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

func (h *Handler) SellerSignUp(c *gin.Context) {
	var input model.Seller

	// Валидация входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 Bad Request для ошибок валидации
		return
	}

	// Регистрация нового продавца
	if err := h.services.SellerSignUp(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error в случае ошибки сервера
		return
	}

	// 201 Created после успешной регистрации
	c.JSON(http.StatusCreated, "Seller registered successfully")
}

func (h *Handler) GetSeller(c *gin.Context) {
	sellerID := c.GetString("user_id")

	// Проверка роли пользователя, чтобы убедиться, что это продавец
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden, если доступ запрещён
		return
	}

	// Получение данных о продавце
	seller, err := h.services.GetSeller(sellerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error при ошибке на сервере
		return
	}

	// 200 OK и данные о продавце
	c.JSON(http.StatusOK, seller)
}

func (h *Handler) UpdateSeller(c *gin.Context) {
	sellerID := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden при недостатке прав
		return
	}

	// Получение данных для обновления
	var input model.Seller
	input.ID = sellerID

	// Валидация данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 Bad Request при ошибке валидации
		return
	}

	// Обновление данных продавца
	if err := h.services.UpdateSeller(sellerID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error при ошибке сервера
		return
	}

	// 200 OK после успешного обновления
	c.JSON(http.StatusOK, "Seller updated successfully")
}

func (h *Handler) SellerSignIn(c *gin.Context) {
	var input model.LoginRequest

	// Валидация входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 Bad Request для ошибок валидации
		return
	}

	// Проверка учетных данных
	seller, err := h.services.SellerSignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error()) // 401 Unauthorized при ошибке аутентификации
		return
	}

	// 200 OK и данные о продавце
	c.JSON(http.StatusOK, seller)
}
