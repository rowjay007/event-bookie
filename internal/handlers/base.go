package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeHandler is a handler function to serve the welcome message
func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"maintainer": "Rowland Adimoha",
		"message":    "Event Booking API built with Go and Supabase",
		"source":     "https://github.com/rowjay007/event-bookie",
	})
}
