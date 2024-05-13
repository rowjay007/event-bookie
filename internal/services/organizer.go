package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type OrganizerService struct {
    organizerRepo repository.OrganizerRepository
}

func NewOrganizerService(organizerRepo repository.OrganizerRepository) *OrganizerService {
    return &OrganizerService{organizerRepo}
}

func (s *OrganizerService) GetAllOrganizers() ([]models.Organizer, error) {
    return s.organizerRepo.GetAllOrganizers()
}

func (s *OrganizerService) GetOrganizerByID(id uint) (*models.Organizer, error) {
    return s.organizerRepo.GetOrganizerByID(id)
}

// Implement other methods like CreateOrganizer, UpdateOrganizer, DeleteOrganizer as needed
