// handlers/event.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAPIInfo returns information about the API
func GetAPIInfo(c *gin.Context) {
	response := gin.H{
		"code":       200,
		"message":    "Event Booking API built with Go and Supabase",
		"maintainer": "Rowland Adimoha",
		"source":     "https://github.com/rowjay007/event-bookie",
	}
	c.JSON(http.StatusOK, response)
}
