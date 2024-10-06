package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddToCart
// @Summary      Add an item to cart
// @Description  Adds a specified item with a specified quantity to the buyer's cart. Only accessible by buyers.
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer {JWT}"
// @Param        input  body  model.AddToCartInput  true  "Item ID and Quantity"
// @Success      200  {string}  string  "Item added to cart"
// @Failure      400  {string}  string  "Invalid input"
// @Failure      403  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /cart/add [post]
func (h *Handler) AddToCart(c *gin.Context) {
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя, чтобы убедиться, что он является покупателем
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	var input struct {
		ItemID   int `json:"item_id"`
		Quantity int `json:"quantity"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input") // 400 Bad Request
		return
	}

	err := h.services.AddToCart(buyerID, input.ItemID, input.Quantity)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"}) // 200 OK
}

// GetCart
// @Summary      Get cart contents
// @Description  Retrieves the contents of the buyer's cart. Only accessible by buyers.
// @Tags         cart
// @Produce      json
// @Param Authorization header string true "Bearer {JWT}"
// @Success      200  {array}  model.CartOutput  "Cart items"
// @Failure      403  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /cart [get]
func (h *Handler) GetCart(c *gin.Context) {
	buyerID := c.GetString("user_id")

	// Проверка роли пользователя, чтобы убедиться, что он является покупателем
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	cart, err := h.services.GetCart(buyerID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, cart) // 200 OK
}

// RemoveFromCart
// @Summary      Remove an item from the cart
// @Description  Removes a specified item from the buyer's cart by cart item ID. Only accessible by buyers.
// @Tags         cart
// @Produce      json
// @Param Authorization header string true "Bearer {JWT}"
// @Param        cart_item_id  query  int  true  "Cart Item ID"
// @Success      200  {string}  string  "Item removed from cart"
// @Failure      400  {string}  string  "Invalid cart item ID"
// @Failure      403  {string}  string  "You are not authorized to access this resource"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /cart/remove [delete]
func (h *Handler) RemoveFromCart(c *gin.Context) {
	if role := c.GetString("role"); role != "buyer" {
		newErrorResponse(c, http.StatusForbidden, "You are not authorized to access this resource") // 403 Forbidden
		return
	}

	cartItemIDStr := c.Query("cart_item_id")
	cartItemID, err := strconv.Atoi(cartItemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid cart item ID") // 400 Bad Request
		return
	}

	err = h.services.RemoveFromCart(cartItemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 Internal Server Error
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"}) // 200 OK
}
