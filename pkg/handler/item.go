package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

// CreateItem создаёт новый товар
// @Summary Create a new item
// @Description Create a new item in the system
// @Tags Items
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param input body model.ItemInput true "Item data"
// @Success 201 {string} string "ItemID"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 401 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Failed to create item"
// @Router /seller/item [post]
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
	itemID, err := h.services.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to create item: "+err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, strconv.Itoa(itemID)) // 201 Created
}

// GetItemById возвращает товар по ID
// @Summary Get item by ID
// @Description Retrieve an item by its ID
// @Tags Items
// @Produce json
// @Param id query string true "Item ID"
// @Success 200 {object} model.CurrentItemInfo "Item details"
// @Failure 400 {object} ErrorResponse "Invalid item ID"
// @Failure 404 {object} ErrorResponse "Item not found"
// @Router /item [get]
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

// UpdateItem обновляет существующий товар
// @Summary Update an existing item
// @Description Update an item by its ID
// @Tags Items
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param id query string true "Item ID"
// @Param input body model.ItemInput true "Item data"
// @Success 200 {string} string "Item updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 401 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 404 {object} ErrorResponse "Item not found"
// @Failure 500 {object} ErrorResponse "Failed to update item"
// @Router /seller/item [patch]
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

// GetItemList возвращает список товаров с фильтрами
// @Summary Get item list
// @Description Retrieve a list of items, optionally filtered
// @Tags Items
// @Accept json
// @Produce json
// @Param filters body model.FilterRequest true "Filter criteria"
// @Success 200 {array} model.ItemInfo "List of items"
// @Failure 500 {object} ErrorResponse "Failed to get items"
// @Router /item [post]
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

// UploadImage загружает изображение для товара
// @Summary Upload an image for an item
// @Description Upload an image for a specific item by ID
// @Tags Items
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param item_id query string true "Item ID"
// @Param image formData file true "Image file"
// @Success 200 {object} SuccessResponse ""url": "url""
// @Failure 400 {object} ErrorResponse "Invalid item ID or failed to upload image"
// @Failure 401 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 404 {object} ErrorResponse "Item not found"
// @Failure 500 {object} ErrorResponse "Failed to save image"
// @Router /seller/item/image [post]
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
