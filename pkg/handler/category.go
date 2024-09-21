package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateCategory(c *gin.Context) {
	var input model.Category

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.CreateCategory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.DeleteCategory(categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) GetCategoryList(c *gin.Context) {
	categories, err := h.services.GetCategoryList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}
