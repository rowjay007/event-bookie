package handlers

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse is a standardized error response structure
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendErrorResponse sends a standardized error response
func SendErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{
		Code:    code,
		Message: message,
	})
	c.Abort()
}
