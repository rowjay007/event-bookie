package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type EventService struct {
	EventRepo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{EventRepo: repo}
}

func (es *EventService) CreateEvent(event *models.Event) error {
	return es.EventRepo.Create(event)
}


func (es *EventService) GetAllEvents(params map[string]string, offset, limit int) ([]models.Event, int64, error) {
	events, total, err := es.EventRepo.GetAll(params, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return events, total, nil
}


func (es *EventService) GetEventByID(id uint) (*models.Event, error) {
	return es.EventRepo.GetByID(id)
}

func (es *EventService) UpdateEvent(event *models.Event) error {
	return es.EventRepo.Update(event)
}

func (es *EventService) DeleteEvent(id uint) error {
	return es.EventRepo.Delete(id)
}
