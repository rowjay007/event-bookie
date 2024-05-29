package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

type EventHandler struct {
	EventService *service.EventService
}

func NewEventHandler(service *service.EventService) *EventHandler {
	return &EventHandler{EventService: service}
}

// @Summary Create a new event
// @Description Create a new event
// @Tags Events
// @Accept json
// @Produce json
// @Param event body models.Event true "Event object"
// @Success 201 {object} models.Event
// @Failure 400 {object} gin.H "Event object is invalid"
// @Failure 500 {object} gin.H "Failed to create event"
// @Router /api/v1/events [post]
func (eh *EventHandler) CreateEvent(c *gin.Context) {
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "02-01-2006 15:04:05"
	startTime, err := time.Parse(layout, input["start_time"].(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}
	endTime, err := time.Parse(layout, input["end_time"].(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	event := models.Event{
		Title:       input["title"].(string),
		Description: input["description"].(string),
		StartTime:   startTime,
		EndTime:     endTime,
		VenueID:     uint(input["venue_id"].(float64)), 
		OrganizerID: uint(input["organizer_id"].(float64)),
		CategoryID:  uint(input["category_id"].(float64)),
	}

	if err := eh.EventService.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusCreated, event)
}


// @Summary Retrieve all events with optional filtering, sorting, and pagination
// @Description Retrieve all events with optional filtering, sorting, and pagination
// @Produce json
// @Tags Events
// @Param offset query integer false "Offset for pagination"
// @Param limit query integer false "Limit for pagination"
// @Success 200 {object} gin.H "Events and total count"
// @Failure 400 {object} gin.H "Invalid query parameters"
// @Failure 500 {object} gin.H "Failed to fetch events"
// @Router /api/v1/events [get]
func (eh *EventHandler) GetAllEvents(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	offset, _ := strconv.Atoi(queryParams.Get("offset"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))

	// Construct a map of parameters
	params := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	events, total, err := eh.EventService.GetAllEvents(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	response := gin.H{
		"events": events,
		"total":  total,
	}

	// Send the response
	c.JSON(http.StatusOK, response)
}


// @Summary Retrieve an event by ID
// @Description Retrieve an event by ID
// @Produce json
// @Tags Events
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} gin.H "Invalid event ID"
// @Failure 500 {object} gin.H "Failed to fetch event"
// @Router /api/v1/events/{id} [get]
func (eh *EventHandler) GetEventByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	event, err := eh.EventService.GetEventByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// @Summary Update an existing event
// @Description Update an existing event
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body models.Event true "Event object"
// @Success 200 {object} models.Event
// @Failure 400 {object} gin.H "Event object is invalid"
// @Failure 500 {object} gin.H "Failed to update event"
// @Router /api/v1/events/{id} [put]
func (eh *EventHandler) UpdateEvent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    var input map[string]interface{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    event, err := eh.EventService.GetEventByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
        return
    }

    // Update fields if provided in the request
    if title, ok := input["title"].(string); ok {
        event.Title = title
    }
    if description, ok := input["description"].(string); ok {
        event.Description = description
    }
    if startTimeStr, ok := input["start_time"].(string); ok {
        startTime, err := time.Parse("02-01-2006 15:04:05", startTimeStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
            return
        }
        event.StartTime = startTime
    }
    if endTimeStr, ok := input["end_time"].(string); ok {
        endTime, err := time.Parse("02-01-2006 15:04:05", endTimeStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
            return
        }
        event.EndTime = endTime
    }
    if venueID, ok := input["venue_id"].(float64); ok {
        event.VenueID = uint(venueID)
    }
    if organizerID, ok := input["organizer_id"].(float64); ok {
        event.OrganizerID = uint(organizerID)
    }
    if categoryID, ok := input["category_id"].(float64); ok {
        event.CategoryID = uint(categoryID)
    }

    if err := eh.EventService.UpdateEvent(event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
        return
    }
    c.JSON(http.StatusOK, event)
}


// @Summary Delete an event by ID
// @Description Delete an event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {string} string "Event deleted successfully"
// @Failure 400 {object} gin.H "Invalid event ID"
// @Failure 500 {object} gin.H "Failed to delete event"
// @Router /api/v1/events/{id} [delete]
func (eh *EventHandler) DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := eh.EventService.DeleteEvent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
