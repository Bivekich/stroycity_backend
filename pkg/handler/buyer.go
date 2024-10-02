package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

func (h *Handler) BuyerSignUp(c *gin.Context) {
	var input model.Buyer

	// Ошибка при валидации данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при попытке зарегистрировать пользователя
	if err := h.services.BuyerSignUp(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to sign up: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Buyer signed up successfully") // 201 Created
}

func (h *Handler) GetBuyer(c *gin.Context) {
	id := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	// Ошибка при получении данных пользователя
	buyer, err := h.services.GetBuyer(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Buyer not found: "+err.Error()) // 404 Not Found
		return
	}
	c.JSON(http.StatusOK, buyer) // 200 OK
}

func (h *Handler) UpdateBuyer(c *gin.Context) {
	id := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	var input model.Buyer
	// Ошибка при валидации входных данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при обновлении данных
	if err := h.services.UpdateBuyer(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to update buyer: "+err.Error()) // 500 Internal Server Error
		return
	}
	c.JSON(http.StatusOK, "Buyer updated successfully") // 200 OK
}

func (h *Handler) BuyerSignIn(c *gin.Context) {
	var input model.LoginRequest

	// Ошибка при валидации данных для входа
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid login data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при аутентификации
	buyer, err := h.services.BuyerSignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid credentials: "+err.Error()) // 401 Unauthorized
		return
	}

	c.JSON(http.StatusOK, buyer) // 200 OK
}
