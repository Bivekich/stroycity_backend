package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

// BuyerSignUp создаёт нового покупателя
// @Summary Create a new buyer
// @Description Register a new buyer in the system
// @Tags Buyers
// @Accept json
// @Produce json
// @Param input body model.Buyer true "Buyer data"
// @Success 201 {string} string "Buyer signed up successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 500 {object} ErrorResponse "Failed to sign up"
// @Router /sign-up/buyer [post]
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

// GetBuyer возвращает информацию о покупателе
// @Summary Get buyer information
// @Description Retrieve buyer information by ID
// @Tags Buyers
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Success 200 {object} model.BuyerOutput "Buyer data"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 404 {object} ErrorResponse "Buyer not found"
// @Router /buyer [get]
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

// UpdateBuyer обновляет информацию о покупателе
// @Summary Update buyer information
// @Description Update buyer details
// @Tags Buyers
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.Buyer true "Buyer data"
// @Success 200 {string} string "Buyer updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to update buyer"
// @Router /buyer [patch]
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

// BuyerSignIn аутентифицирует покупателя
// @Summary Buyer login
// @Description Authenticate a buyer with email and password
// @Tags Buyers
// @Accept json
// @Produce json
// @Param input body model.LoginRequest true "Login data"
// @Success 200 {object} model.BuyerOutput "Authenticated buyer data"
// @Failure 400 {object} ErrorResponse "Invalid login data"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /sign-in/buyer [post]
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
