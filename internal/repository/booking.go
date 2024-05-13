package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type BookingRepository struct {
    db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
    return &BookingRepository{db}
}

func (r *BookingRepository) GetAllBookings() ([]models.Booking, error) {
    var bookings []models.Booking
    if err := r.db.Find(&bookings).Error; err != nil {
        return nil, err
    }
    return bookings, nil
}

func (r *BookingRepository) GetBookingByID(id uint) (*models.Booking, error) {
    var booking models.Booking
    if err := r.db.First(&booking, id).Error; err != nil {
        return nil, err
    }
    return &booking, nil
}

// Implement other methods like CreateBooking, UpdateBooking, DeleteBooking as needed
