package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) GetItemById(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	item, err := h.services.GetItemById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var input model.Item

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	input.ID = id
	if err := h.services.UpdateItem(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "OK!")
}

func (h *Handler) GetItemList(c *gin.Context) {
	var filters model.FilterRequest

	if err := c.ShouldBindJSON(&filters); err != nil && err.Error() == "EOF" {
		items, err := h.services.GetAllItems()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, errors.New("failed to get items").Error())
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}

	items, err := h.services.GetItems(filters.BrandIDs, filters.SellerIDs, filters.CategoryIDs, filters.MaterialIDs, filters.MinPrice, filters.MaxPrice)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("Ошибка фильтрации").Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) UploadImage(c *gin.Context) {
	itemIDStr := c.Query("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, errors.New("invalid item ID").Error())
		return
	}

	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, errors.New("failed to upload image").Error())
		return
	}
	defer file.Close()

	filePath, err := h.services.UploadImage(itemID, file, fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": filePath})
}
