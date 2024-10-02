package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateMaterial(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	var input model.Material

	// Ошибка при разборе JSON данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при создании материала
	err := h.services.CreateMaterial(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to create material: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Material created successfully") // 201 Created
}

func (h *Handler) DeleteMaterial(c *gin.Context) {
	// Проверка роли пользователя
	if role := c.GetString("role"); role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	// Получение и проверка material_id
	materialIdStr := c.Query("material_id")
	materialId, err := strconv.Atoi(materialIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid material ID: "+err.Error()) // 400 Bad Request
		return
	}

	// Ошибка при удалении материала
	err = h.services.DeleteMaterial(materialId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to delete material: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, "Material deleted successfully") // 200 OK
}

func (h *Handler) GetMaterialList(c *gin.Context) {
	// Получение списка материалов
	materials, err := h.services.GetMaterialList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve material list: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, materials) // 200 OK
}
