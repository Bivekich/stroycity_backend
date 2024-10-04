package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

// CreateMaterial создаёт новый материал
// @Summary Create a new material
// @Description Create a new material in the system
// @Tags Materials
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.MaterialInput true "Material data"
// @Success 201 {string} string "Material created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to create material"
// @Router /admin/material [post]
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

// DeleteMaterial удаляет материал по ID
// @Summary Delete a material by ID
// @Description Delete a material from the system by its ID
// @Tags Materials
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param material_id query string true "Material ID"
// @Success 200 {string} string "Material deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid material ID"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to delete material"
// @Router /admin/material [delete]
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

// GetMaterialList возвращает список материалов
// @Summary Get material list
// @Description Retrieve a list of materials from the system
// @Tags Materials
// @Produce json
// @Success 200 {array} model.MaterialOutput "List of materials"
// @Failure 500 {object} ErrorResponse "Failed to retrieve material list"
// @Router /material [get]
func (h *Handler) GetMaterialList(c *gin.Context) {
	// Получение списка материалов
	materials, err := h.services.GetMaterialList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve material list: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, materials) // 200 OK
}
