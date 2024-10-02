package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

func (h *Handler) AdminSignUp(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	var input model.Admin

	// Ошибка при валидации данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при попытке зарегистрировать пользователя
	if err := h.services.AdminSignUp(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to sign up: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Admin signed up successfully") // 201 Created
}

func (h *Handler) AdminSignIn(c *gin.Context) {
	var input model.AdminLoginRequest

	// Ошибка при валидации данных для входа
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid login data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при аутентификации
	token, err := h.services.AdminSignIn(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid credentials: "+err.Error()) // 401 Unauthorized
		return
	}

	c.JSON(http.StatusOK, token) // 200 OK
}
