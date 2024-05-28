package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

// VenueHandler handles venue-related requests
type VenueHandler struct {
	VenueService *service.VenueService
}

// NewVenueHandler creates a new VenueHandler
func NewVenueHandler(service *service.VenueService) *VenueHandler {
	return &VenueHandler{VenueService: service}
}

// CreateVenue creates a new venue
// @Summary Create a new venue
// @Description Create a new venue
// @Tags Venues
// @Accept json
// @Produce json
// @Param venue body models.Venue true "Venue"
// @Success 201 {object} models.Venue
// @Failure 400 {object} gin.H "Venue is invalid"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/venues [post]
func (vh *VenueHandler) CreateVenue(c *gin.Context) {
	var venue models.Venue
	if err := c.ShouldBindJSON(&venue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := vh.VenueService.CreateVenue(&venue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create venue"})
		return
	}
	c.JSON(http.StatusCreated, venue)
}

// GetAllVenues retrieves all venues with optional filtering, sorting, and pagination
// @Summary Get all venues
// @Description Get all venues with optional filtering, sorting, and pagination
// @Tags Venues
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param name query string false "Name"
// @Param location query string false "Location"
// @Param sort_by query string false "Sort By"
// @Param sort_order query string false "Sort Order"
// @Success 200 {object} gin.H "Venues and total count"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/venues [get]
func (vh *VenueHandler) GetAllVenues(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	offset, _ := strconv.Atoi(queryParams.Get("offset"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))

	params := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	venues, total, err := vh.VenueService.GetAllVenues(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch venues"})
		return
	}

	response := gin.H{
		"venues": venues,
		"total":  total,
	}
	c.JSON(http.StatusOK, response)
}

// GetVenueByID retrieves a venue by ID
// @Summary Get a venue by ID
// @Description Get a venue by ID
// @Tags Venues
// @Accept json
// @Produce json
// @Param id path uint true "Venue ID"
// @Success 200 {object} models.Venue
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/venues/{id} [get]
func (vh *VenueHandler) GetVenueByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	venue, err := vh.VenueService.GetVenueByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch venue"})
		return
	}
	c.JSON(http.StatusOK, venue)
}

// UpdateVenue updates an existing venue
// @Summary Update an existing venue
// @Description Update an existing venue
// @Tags Venues
// @Accept json
// @Produce json
// @Param id path uint true "Venue ID"
// @Param venue body models.Venue true "Venue"
// @Success 200 {object} models.Venue
// @Failure 400 {object} gin.H "Venue is invalid"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/venues/{id} [put]
func (vh *VenueHandler) UpdateVenue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var venue models.Venue
	if err := c.ShouldBindJSON(&venue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	venue.ID = uint(id)
	if err := vh.VenueService.UpdateVenue(&venue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update venue"})
		return
	}
	c.JSON(http.StatusOK, venue)
}

// DeleteVenue deletes a venue by ID
// @Summary Delete a venue by ID
// @Description Delete a venue by ID
// @Tags Venues
// @Accept json
// @Produce json
// @Param id path uint true "Venue ID"
// @Success 200 {object} gin.H "Venue deleted successfully"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/venues/{id} [delete]
func (vh *VenueHandler) DeleteVenue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := vh.VenueService.DeleteVenue(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete venue"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Venue deleted successfully"})
}
