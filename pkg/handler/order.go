package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrder создает заказ на основе товаров в корзине пользователя
// @Summary Create a new order
// @Description Create an order from the items in the cart of the current buyer
// @Tags Orders
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Success 200 {object} SuccessResponse "Order created successfully"
// @Failure 400 {object} ErrorResponse "Cart is empty"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /buyer/order [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource")
		return
	}

	// Получение товаров из корзины
	cartItems, err := h.services.GetCart(buyerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(cartItems.CartItems) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Cart is empty")
		return
	}

	// Создание заказа
	err = h.services.CreateOrder(buyerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Очистка корзины после заказа
	err = h.services.ClearCart(buyerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

// GetOrder возвращает заказ по ID
// @Summary Get an order by ID
// @Description Retrieve an order by its ID for the buyer
// @Tags Orders
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param order_id query int true "Order ID"
// @Success 200 {object} model.Order "Order data"
// @Failure 400 {object} ErrorResponse "Invalid order ID"
// @Failure 403 {object} ErrorResponse "You are not authorized to access this resource"
// @Failure 404 {object} ErrorResponse "Order not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /buyer/order [get]
func (h *Handler) GetOrder(c *gin.Context) {
	// Получение ID пользователя
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource")
		return
	}

	// Получение ID заказа из параметров запроса
	orderIDStr := c.Query("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	// Получение заказа по ID
	order, err := h.services.GetOrderById(orderID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Order not found")
		return
	}

	// Проверка, принадлежит ли заказ текущему пользователю
	if order.BuyerID != buyerID {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource")
		return
	}

	// Возвращение данных заказа
	c.JSON(http.StatusOK, order)
}
