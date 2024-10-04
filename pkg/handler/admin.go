package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

// AdminSignUp godoc
// @Summary Sign up as an admin
// @Tags admin
// @Description Registers a new admin user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.AdminLoginRequest true "Admin signup data"
// @Success 201 {object} SuccessResponse "Admin signed up successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to sign up"
// @Router /admin/sign-up [post]
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

	c.JSON(http.StatusCreated, SuccessResponse{Message: "Admin signed up successfully"}) // 201 Created
}

// AdminSignIn godoc
// @Summary Sign in as an admin
// @Tags admin
// @Description Logs in an admin user
// @Accept json
// @Produce json
// @Param input body model.AdminLoginRequest true "Admin login data"
// @Success 200 {object} TokenResponse "Authentication token"
// @Failure 400 {object} ErrorResponse "Invalid login data"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /sign-in/admin [post]
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

	c.JSON(http.StatusOK, TokenResponse{Token: token}) // 200 OK
}
