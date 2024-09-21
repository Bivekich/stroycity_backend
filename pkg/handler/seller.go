package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"stroycity/pkg/model"
)

func (h *Handler) SellerSignUp(c *gin.Context) {
	var input model.Seller

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.SellerSignUp(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) GetSeller(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("id is required").Error())
		return
	}
	seller, err := h.services.GetSeller(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, seller)
}

func (h *Handler) UpdateSeller(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("id is required").Error())
		return
	}
	var input model.Seller
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.UpdateSeller(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) SellerSignIn(c *gin.Context) {
	var input model.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	seller, err := h.services.SellerSignIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, seller)
}
