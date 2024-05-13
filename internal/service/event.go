package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type EventService struct {
    eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) *EventService {
    return &EventService{eventRepo}
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
    return s.eventRepo.GetAllEvents()
}

func (s *EventService) GetEventByID(id uint) (*models.Event, error) {
    return s.eventRepo.GetEventByID(id)
}

// Implement other methods like CreateEvent, UpdateEvent, DeleteEvent as needed
