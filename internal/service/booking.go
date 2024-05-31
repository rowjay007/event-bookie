package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type BookingService struct {
	BookingRepo *repository.BookingRepository
}

func NewBookingService(repo *repository.BookingRepository) *BookingService {
	return &BookingService{BookingRepo: repo}
}

func (bs *BookingService) CreateBooking(booking *models.Booking) error {
	return bs.BookingRepo.Create(booking)
}

func (bs *BookingService) GetAllBookings(params map[string]string, offset, limit int) ([]models.Booking, int64, error) {
	return bs.BookingRepo.GetAll(params, offset, limit)
}

func (bs *BookingService) GetBookingByID(id uint) (*models.Booking, error) {
	return bs.BookingRepo.GetByID(id)
}

func (bs *BookingService) UpdateBooking(booking *models.Booking) error {
	return bs.BookingRepo.Update(booking)
}

func (bs *BookingService) DeleteBooking(id uint) error {
	return bs.BookingRepo.Delete(id)
}
