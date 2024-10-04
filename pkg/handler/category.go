package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

// CreateCategory создаёт новую категорию
// @Summary Create a new category
// @Description Create a new category in the system
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.CategoryInput true "Category data"
// @Success 200 {string} string "Category created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 401 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to create category"
// @Router /admin/category [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource")
		return
	}

	var input model.Category
	// Ошибка при разборе JSON данных
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error())
		return
	}

	// Ошибка при создании категории
	err := h.services.CreateCategory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to create category: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Category created successfully")
}

// DeleteCategory удаляет категорию
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags Categories
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param category_id query string true "Category ID"
// @Success 200 {string} string "Category deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid category ID"
// @Failure 401 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to delete category"
// @Router /admin/category [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource")
		return
	}

	// Получение и проверка category_id
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid category ID: "+err.Error())
		return
	}

	// Ошибка при удалении категории
	err = h.services.DeleteCategory(categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to delete category: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Category deleted successfully")
}

// GetCategoryList возвращает список категорий
// @Summary Get category list
// @Description Retrieve a list of all categories
// @Tags Categories
// @Produce json
// @Success 200 {array} model.CategoryOutput "List of categories"
// @Failure 500 {object} ErrorResponse "Failed to retrieve category list"
// @Router /category [get]
func (h *Handler) GetCategoryList(c *gin.Context) {
	// Ошибка при получении списка категорий
	categories, err := h.services.GetCategoryList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve category list: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}
