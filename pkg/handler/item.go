package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

func (h *Handler) CreateItem(c *gin.Context) {
	var input model.Item

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.CreateItem(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "OK!")
}
