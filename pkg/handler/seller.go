package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

// SellerSignUp регистрирует нового продавца
// @Summary Register a new seller
// @Description Register a new seller with the provided details
// @Tags Sellers
// @Accept json
// @Produce json
// @Param input body model.SellerInput true "Seller data"
// @Success 201 {string} string "Seller registered successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 500 {object} ErrorResponse "Failed to register seller"
// @Router /sign-up/seller [post]
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

// GetSeller возвращает информацию о продавце
// @Summary Get seller information
// @Description Retrieve seller details by ID
// @Tags Sellers
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Success 200 {object} model.SellerOutput "Seller data"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to retrieve seller"
// @Router /seller [get]
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

// UpdateSeller обновляет информацию о продавце
// @Summary Update seller information
// @Description Update the details of the seller
// @Tags Sellers
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.SellerInput true "Updated seller data"
// @Success 200 {string} string "Seller updated successfully"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 500 {object} ErrorResponse "Failed to update seller"
// @Router /seller [patch]
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

// SellerSignIn аутентифицирует продавца
// @Summary Seller sign in
// @Description Authenticate a seller with their email and password
// @Tags Sellers
// @Accept json
// @Produce json
// @Param input body model.SellerSignInResponse true "Login credentials"
// @Success 200 {object} model.Seller "Seller data"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /sign-in/seller [post]
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

// GetSellerEarnings возвращает заработок продавца за текущую и прошлую неделю
// @Summary Get seller earnings
// @Description Retrieve seller earnings for the current and previous week
// @Tags Sellers
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Success 200 {object} model.Statistic "Earnings for current and last week"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to retrieve earnings"
// @Router /seller/statistic [get]
func (h *Handler) GetSellerEarnings(c *gin.Context) {
	sellerID := c.GetString("user_id")

	// Проверка роли пользователя, чтобы убедиться, что это продавец
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden, если доступ запрещён
		return
	}

	// Получение заработка продавца за текущую и прошлую недели
	currentWeek, lastWeek, err := h.services.GetSellerEarnings(sellerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error при ошибке сервера
		return
	}

	// Формирование ответа
	statistic := struct {
		CurrentWeek float64 `json:"current_week"`
		LastWeek    float64 `json:"last_week"`
	}{
		CurrentWeek: currentWeek,
		LastWeek:    lastWeek,
	}

	// Возвращаем ответ с кодом 200 OK
	c.JSON(http.StatusOK, statistic)
}
