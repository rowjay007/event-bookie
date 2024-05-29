package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (er *EventRepository) Create(event *models.Event) error {
	return er.DB.Create(event).Error
}

func (er *EventRepository) GetAll(params map[string]string, offset, limit int) ([]models.Event, int64, error) {
	var events []models.Event
	var total int64

	query := er.DB.Model(&models.Event{})

	if title, ok := params["title"]; ok && title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if sortBy, ok := params["sort_by"]; ok && sortBy != "" {
		order := "ASC"
		if sortOrder, ok := params["sort_order"]; ok && sortOrder == "desc" {
			order = "DESC"
		}
		query = query.Order(sortBy + " " + order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}



func (er *EventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	if err := er.DB.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (er *EventRepository) Update(event *models.Event) error {
	return er.DB.Save(event).Error
}

func (er *EventRepository) Delete(id uint) error {
	return er.DB.Delete(&models.Event{}, id).Error
}
