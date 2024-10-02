package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

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

func (h *Handler) GetCategoryList(c *gin.Context) {
	// Ошибка при получении списка категорий
	categories, err := h.services.GetCategoryList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve category list: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}
