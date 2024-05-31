package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

type BookingHandler struct {
	BookingService *service.BookingService
}

func NewBookingHandler(service *service.BookingService) *BookingHandler {
	return &BookingHandler{BookingService: service}
}

// CreateBooking creates a new booking
// @Summary Create a new booking
// @Description Create a new booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Param booking body models.Booking true "Booking to create"
// @Success 201 {object} models.Booking
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Failed to create booking"
// @Router /api/v1/bookings [post]
func (bh *BookingHandler) CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := bh.BookingService.CreateBooking(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}
	c.JSON(http.StatusCreated, booking)
}

// GetAllBookings retrieves all bookings with optional filtering, sorting, and pagination
// @Summary Get all bookings
// @Description Get all bookings with optional filtering, sorting, and pagination
// @Tags Bookings
// @Accept json
// @Produce json
// @Param status query string false "Filter by status"
// @Param user_id query int false "Filter by user ID"
// @Param event_id query int false "Filter by event ID"
// @Param sort_by query string false "Field to sort by"
// @Param sort_order query string false "Sort order (asc or desc)"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} gin.H "Bookings and total count"
// @Failure 500 {object} gin.H "Failed to fetch bookings"
// @Router /api/v1/bookings [get]
func (bh *BookingHandler) GetAllBookings(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	offsetStr := queryParams.Get("offset")
	limitStr := queryParams.Get("limit")

	var offset, limit int
	var err error

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	} else {
		offset = -1
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	} else {
		limit = -1
	}

	params := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	bookings, total, err := bh.BookingService.GetAllBookings(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	response := gin.H{
		"bookings": bookings,
		"total":    total,
	}

	c.JSON(http.StatusOK, response)
}

// GetBookingByID retrieves a booking by ID
// @Summary Get booking by ID
// @Description Get booking by ID
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} models.Booking
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Failed to fetch booking"
// @Router /api/v1/bookings/{id} [get]
func (bh *BookingHandler) GetBookingByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	booking, err := bh.BookingService.GetBookingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// UpdateBooking updates an existing booking
// @Summary Update an existing booking
// @Description Update an existing booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body models.Booking true "Booking to update"
// @Success 200 {object} models.Booking
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Failed to update booking"
// @Router /api/v1/bookings/{id} [put]
func (bh *BookingHandler) UpdateBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking.ID = uint(id)
	if err := bh.BookingService.UpdateBooking(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}
	c.JSON(http.StatusOK, booking)
}

// DeleteBooking deletes a booking by ID
// @Summary Delete a booking by ID
// @Description Delete a booking by ID
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} gin.H "Booking deleted successfully"
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Failed to delete booking"
// @Router /api/v1/bookings/{id} [delete]
func (bh *BookingHandler) DeleteBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := bh.BookingService.DeleteBooking(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
