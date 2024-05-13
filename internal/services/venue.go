package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type VenueService struct {
    venueRepo repository.VenueRepository
}

func NewVenueService(venueRepo repository.VenueRepository) *VenueService {
    return &VenueService{venueRepo}
}

func (s *VenueService) GetAllVenues() ([]models.Venue, error) {
    return s.venueRepo.GetAllVenues()
}

func (s *VenueService) GetVenueByID(id uint) (*models.Venue, error) {
    return s.venueRepo.GetVenueByID(id)
}

// Implement other methods like CreateVenue, UpdateVenue, DeleteVenue as needed
