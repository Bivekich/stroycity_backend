package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateItem(c *gin.Context) {
	// Проверка роли пользователя
	sellerId := c.GetString("user_id")
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	var input model.Item
	input.SellerID = sellerId

	// Валидация JSON данных
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	// Создание товара
	if err := h.services.CreateItem(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to create item: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Item created successfully") // 201 Created
}

func (h *Handler) GetItemById(c *gin.Context) {
	// Получение ID из запроса
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item ID: "+err.Error()) // 400 Bad Request
		return
	}

	// Получение товара по ID
	item, err := h.services.GetItemById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Item not found: "+err.Error()) // 404 Not Found
		return
	}
	c.JSON(http.StatusOK, item) // 200 OK
}

func (h *Handler) UpdateItem(c *gin.Context) {
	// Получение ID из запроса
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item ID: "+err.Error()) // 400 Bad Request
		return
	}

	// Проверка роли пользователя
	sellerId := c.GetString("user_id")
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	// Получение товара для проверки прав доступа
	item, err := h.services.GetItemById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Item not found: "+err.Error()) // 404 Not Found
		return
	}

	// Проверка прав на изменение товара
	if item.SellerID != sellerId {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to modify this item") // 401 Unauthorized
		return
	}

	var input model.Item
	// Валидация данных
	if err = c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data: "+err.Error()) // 400 Bad Request
		return
	}

	input.ID = id
	// Обновление товара
	if err = h.services.UpdateItem(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to update item: "+err.Error()) // 500 Internal Server Error
		return
	}
	c.JSON(http.StatusOK, "Item updated successfully") // 200 OK
}

func (h *Handler) GetItemList(c *gin.Context) {
	var filters model.FilterRequest

	// Проверка, есть ли фильтры
	if err := c.ShouldBindJSON(&filters); err != nil && err.Error() == "EOF" {
		// Получение всех товаров без фильтрации
		items, err := h.services.GetAllItems()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "Failed to get items") // 500 Internal Server Error
			return
		}
		c.JSON(http.StatusOK, items) // 200 OK
		return
	}

	// Получение товаров с фильтрами
	items, err := h.services.GetItems(filters.BrandIDs, filters.SellerIDs, filters.CategoryIDs, filters.MaterialIDs, filters.MinPrice, filters.MaxPrice)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Filtering error: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, items) // 200 OK
}

func (h *Handler) UploadImage(c *gin.Context) {
	// Получение ID товара из запроса
	itemIDStr := c.Query("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item ID: "+err.Error()) // 400 Bad Request
		return
	}

	// Проверка роли пользователя
	sellerId := c.GetString("user_id")
	if role := c.GetString("role"); role != "seller" {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to access this resource") // 401 Unauthorized
		return
	}

	// Проверка прав на добавление изображения
	item, err := h.services.GetItemById(itemID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Item not found: "+err.Error()) // 404 Not Found
		return
	}
	if item.SellerID != sellerId {
		newErrorResponse(c, http.StatusUnauthorized, "You are not authorized to modify this item") // 401 Unauthorized
		return
	}

	// Загрузка файла
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to upload image: "+err.Error()) // 400 Bad Request
		return
	}
	defer file.Close()

	// Сохранение изображения
	filePath, err := h.services.UploadImage(itemID, file, fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to save image: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": filePath}) // 200 OK
}
