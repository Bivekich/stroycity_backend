package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateMaterial(c *gin.Context) {
	var input model.Material

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.CreateMaterial(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteMaterial(c *gin.Context) {
	materialIdStr := c.Query("material_id")
	materialId, err := strconv.Atoi(materialIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.DeleteMaterial(materialId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) GetMaterialList(c *gin.Context) {
	materials, err := h.services.GetMaterialList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, materials)
}
