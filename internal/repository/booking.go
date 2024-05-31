package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type BookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (br *BookingRepository) Create(booking *models.Booking) error {
	return br.DB.Create(booking).Error
}

func (br *BookingRepository) GetAll(queryParams map[string]string, offset, limit int) ([]models.Booking, int64, error) {
	var bookings []models.Booking
	var total int64

	query := br.DB.Model(&models.Booking{})

	if status := queryParams["status"]; status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}

	if userID := queryParams["user_id"]; userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	if eventID := queryParams["event_id"]; eventID != "" {
		query = query.Where("event_id = ?", eventID)
	}

	if sortBy := queryParams["sort_by"]; sortBy != "" {
		order := "ASC"
		if sortOrder := queryParams["sort_order"]; sortOrder == "desc" {
			order = "DESC"
		}
		query = query.Order(sortBy + " " + order)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if offset >= 0 && limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&bookings).Error; err != nil {
		return nil, 0, err
	}

	return bookings, total, nil
}

func (br *BookingRepository) GetByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	if err := br.DB.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (br *BookingRepository) Update(booking *models.Booking) error {
	return br.DB.Save(booking).Error
}

func (br *BookingRepository) Delete(id uint) error {
	return br.DB.Delete(&models.Booking{}, id).Error
}
