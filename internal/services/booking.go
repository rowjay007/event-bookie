package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type BookingService struct {
    bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) *BookingService {
    return &BookingService{bookingRepo}
}

func (s *BookingService) GetAllBookings() ([]models.Booking, error) {
    return s.bookingRepo.GetAllBookings()
}

func (s *BookingService) GetBookingByID(id uint) (*models.Booking, error) {
    return s.bookingRepo.GetBookingByID(id)
}

// Implement other methods like CreateBooking, UpdateBooking, DeleteBooking as needed
