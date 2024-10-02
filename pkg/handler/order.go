package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	// Получение ID пользователя
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя, чтобы убедиться, что он является покупателем
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden для неавторизованного доступа
		return
	}

	var order model.Order
	order.BuyerID = buyerID

	// Валидация входных данных заказа
	if err := c.ShouldBindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 Bad Request для ошибки валидации
		return
	}

	// Создание заказа
	err := h.services.CreateOrder(order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error при ошибке сервера
		return
	}

	// Возвращение успешного ответа
	c.JSON(http.StatusCreated, "Order created successfully") // 201 Created после успешного создания заказа
}

func (h *Handler) GetOrder(c *gin.Context) {
	// Получение ID пользователя
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden для неавторизованного доступа
		return
	}

	// Получение ID заказа из параметров запроса
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid order ID") // 400 Bad Request, если ID невалиден
		return
	}

	// Получение заказа по ID
	order, err := h.services.GetOrderById(orderID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "Order not found") // 404 Not Found, если заказ не найден
		return
	}

	// Проверка, принадлежит ли заказ текущему пользователю
	if order.BuyerID != buyerID {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden, если заказ не принадлежит пользователю
		return
	}

	// Возвращение данных заказа
	c.JSON(http.StatusOK, order) // 200 OK для успешного получения данных
}
