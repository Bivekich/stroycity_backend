package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"stroycity/pkg/model"
	"time"
)

// CreateReview обрабатывает создание нового отзыва.
// @Summary Создание нового отзыва
// @Description Позволяет покупателю создать отзыв для товара.
// @Tags reviews
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {JWT}"
// @Param review body model.Review true "Информация о отзыве"
// @Success 201 {string} string "Review created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 403 {string} string "Forbidden"
// @Failure 500 {string} string "Internal Server Error"
// @Router /buyer/review [post]
func (h *Handler) CreateReview(c *gin.Context) {
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource")
		return
	}

	var review model.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 Bad Request
		return
	}

	review.BuyerID = buyerID
	review.CreatedAt = time.Now()

	if err := h.services.CreateReview(review); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusCreated, "Review created successfully") // 201 Created
}

// GetReviews обрабатывает получение отзывов для товара.
// @Summary Получение отзывов для товара
// @Description Позволяет получить список отзывов для указанного товара по его ID.
// @Tags reviews
// @Accept json
// @Produce json
// @Param item_id query int true "ID товара"
// @Success 200 {array} model.Review "Список отзывов для товара"
// @Failure 400 {string} string "Invalid item ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /review [get]
func (h *Handler) GetReviews(c *gin.Context) {
	itemIDStr := c.Query("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item ID") // 400 Bad Request
		return
	}

	reviews, err := h.services.GetReviewsByItemID(itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, reviews) // 200 OK
}
