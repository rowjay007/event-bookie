package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type VenueService struct {
	VenueRepo *repository.VenueRepository
}

func NewVenueService(repo *repository.VenueRepository) *VenueService {
	return &VenueService{VenueRepo: repo}
}

func (vs *VenueService) CreateVenue(venue *models.Venue) error {
	return vs.VenueRepo.Create(venue)
}

func (vs *VenueService) GetAllVenues(params map[string]string, offset, limit int) ([]models.Venue, int64, error) {
	return vs.VenueRepo.GetAll(params, offset, limit)
}

func (vs *VenueService) GetVenueByID(id uint) (*models.Venue, error) {
	return vs.VenueRepo.GetByID(id)
}

func (vs *VenueService) UpdateVenue(venue *models.Venue) error {
	return vs.VenueRepo.Update(venue)
}

func (vs *VenueService) DeleteVenue(id uint) error {
	return vs.VenueRepo.Delete(id)
}
