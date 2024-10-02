package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateBrand(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	var input model.Brand
	// Валидация JSON данных
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Создание бренда
	if err := h.services.CreateBrand(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to create brand: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Brand created successfully") // 201 Created
}

func (h *Handler) DeleteBrand(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	brandIdStr := c.Query("brand_id")
	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid brand ID: "+err.Error()) // 400 Bad Request
		return
	}

	// Удаление бренда
	if err := h.services.DeleteBrand(brandId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to delete brand: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, "Brand deleted successfully") // 200 OK
}

func (h *Handler) GetBrandList(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	// Получение списка брендов
	brands, err := h.services.GetBrandList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to get brand list: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, brands) // 200 OK
}
