package handler

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

type TokenResponse struct {
	Token string `json:"token"`
}
