package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type OrganizerService struct {
	OrganizerRepo *repository.OrganizerRepository
}

func NewOrganizerService(repo *repository.OrganizerRepository) *OrganizerService {
	return &OrganizerService{OrganizerRepo: repo}
}

func (os *OrganizerService) CreateOrganizer(organizer *models.Organizer) error {
	return os.OrganizerRepo.Create(organizer)
}

func (os *OrganizerService) GetAllOrganizers(params map[string]string, offset, limit int) ([]models.Organizer, int64, error) {
	return os.OrganizerRepo.GetAll(params, offset, limit)
}

func (os *OrganizerService) GetOrganizerByID(id uint) (*models.Organizer, error) {
	return os.OrganizerRepo.GetByID(id)
}

func (os *OrganizerService) UpdateOrganizer(organizer *models.Organizer) error {
	return os.OrganizerRepo.Update(organizer)
}

func (os *OrganizerService) DeleteOrganizer(id uint) error {
	return os.OrganizerRepo.Delete(id)
}
