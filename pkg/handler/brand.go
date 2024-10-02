package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

// CreateBrand
// @Summary      Create a new brand
// @Description  Creates a new brand with the provided details. Only accessible by admin.
// @Tags         brands
// @Accept       json
// @Produce      json
// @Param        input  body  model.Brand  true  "Brand data"
// @Success      201  {string}  string  "Brand created successfully"
// @Failure      400  {string}  string  "Invalid input data"
// @Failure      401  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Failed to create brand"
// @Router       /admin/brand [post]
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

// DeleteBrand
// @Summary      Delete a brand
// @Description  Deletes a brand by ID. Only accessible by admin.
// @Tags         brands
// @Produce      json
// @Param        brand_id  query  int  true  "Brand ID"
// @Success      200  {string}  string  "Brand deleted successfully"
// @Failure      400  {string}  string  "Invalid brand ID"
// @Failure      401  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Failed to delete brand"
// @Router       /admin/brand [delete]
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

// GetBrandList
// @Summary      Get list of brands
// @Description  Retrieves a list of all brands. Only accessible by admin.
// @Tags         brands
// @Produce      json
// @Success      200  {array}  model.Brand  "List of brands"
// @Failure      401  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Failed to get brand list"
// @Router       /brand [get]
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
