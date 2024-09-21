package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateBrand(c *gin.Context) {
	var input model.Brand

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.CreateBrand(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteBrand(c *gin.Context) {
	brandIdStr := c.Query("brand_id")
	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.DeleteBrand(brandId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) GetBrandList(c *gin.Context) {
	brands, err := h.services.GetBrandList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, brands)
}
