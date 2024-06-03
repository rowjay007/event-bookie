package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

type OrganizerHandler struct {
	OrganizerService *service.OrganizerService
}

func NewOrganizerHandler(service *service.OrganizerService) *OrganizerHandler {
	return &OrganizerHandler{OrganizerService: service}
}

// CreateOrganizer creates a new organizer
// @Summary Create a new organizer
// @Description Creates a new organizer with the provided details
// @Tags Organizers
// @Accept json
// @Produce json
// @Param input body models.Organizer true "Organizer object to be created"
// @Success 201 {object} models.Organizer
// @Failure 400 {object} gin.H "Organizer information is invalid"
// @Failure 500 {object} gin.H "Failed to create organizer"
// @Router /api/v1/organizers [post]
func (oh *OrganizerHandler) CreateOrganizer(c *gin.Context) {
	var organizer models.Organizer
	if err := c.ShouldBindJSON(&organizer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := oh.OrganizerService.CreateOrganizer(&organizer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organizer"})
		return
	}
	c.JSON(http.StatusCreated, organizer)
}

// GetAllOrganizers retrieves all organizers with optional filtering, sorting, and pagination
// @Summary Retrieve all organizers
// @Description Retrieves all organizers with optional filtering, sorting, and pagination
// @Tags Organizers
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} gin.H "An object containing organizers and total count"
// @Failure 400 {object} gin.H "Invalid query params"
// @Failure 500 {object} gin.H "Failed to fetch organizers"
// @Router /api/v1/organizers [get]
func (oh *OrganizerHandler) GetAllOrganizers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	offset, _ := strconv.Atoi(queryParams.Get("offset"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))

	params := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	organizers, total, err := oh.OrganizerService.GetAllOrganizers(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organizers"})
		return
	}

	response := gin.H{
		"organizers": organizers,
		"total":      total,
	}
	c.JSON(http.StatusOK, response)
}

// GetOrganizerByID retrieves an organizer by ID
// @Summary Retrieve an organizer by ID
// @Description Retrieves an organizer by ID
// @Tags Organizers
// @Accept json
// @Produce json
// @Param id path int true "Organizer ID"
// @Success 200 {object} models.Organizer
// @Failure 404 {object} gin.H "Organizer not found"
// @Failure 500 {object} gin.H "Failed to fetch organizer"
// @Router /api/v1/organizers/{id} [get]
func (oh *OrganizerHandler) GetOrganizerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	organizer, err := oh.OrganizerService.GetOrganizerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch organizer"})
		return
	}
	c.JSON(http.StatusOK, organizer)
}

// UpdateOrganizer updates an existing organizer
// @Summary Update an existing organizer
// @Description Updates an existing organizer with the provided details
// @Tags Organizers
// @Accept json
// @Produce json
// @Param id path int true "Organizer ID"
// @Param input body models.Organizer true "Updated organizer object"
// @Success 200 {object} models.Organizer
// @Failure 400 {object} gin.H "Invalid request body"
// @Failure 404 {object} gin.H "Organizer not found"
// @Failure 500 {object} gin.H "Failed to update organizer"
// @Router /api/v1/organizers/{id} [put]
func (oh *OrganizerHandler) UpdateOrganizer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var organizer models.Organizer
	if err := c.ShouldBindJSON(&organizer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizer.ID = uint(id)
	if err := oh.OrganizerService.UpdateOrganizer(&organizer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organizer"})
		return
	}
	c.JSON(http.StatusOK, organizer)
}

// DeleteOrganizer deletes an organizer by ID
// @Summary Delete an organizer by ID
// @Description Deletes an organizer by ID
// @Tags Organizers
// @Accept json
// @Produce json
// @Param id path int true "Organizer ID"
// @Success 200 {object} gin.H "Organizer deleted successfully"
// @Failure 404 {object} gin.H	"Organizer not found"
// @Failure 500 {object} gin.H	"Failed to delete organizer"
// @Router /api/v1/organizers/{id} [delete]
func (oh *OrganizerHandler) DeleteOrganizer(c *gin.Context) {
id, _ := strconv.Atoi(c.Param("id"))
if err := oh.OrganizerService.DeleteOrganizer(uint(id)); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organizer"})
return
}
c.JSON(http.StatusOK, gin.H{"message": "Organizer deleted successfully"})
}
