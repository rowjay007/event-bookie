package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type CategoryService struct {
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryRepo}
}

func (cs *CategoryService) CreateCategory(category *models.Category) error {
	return cs.CategoryRepository.Create(category)
}

func (cs *CategoryService) GetAllCategories(queryParams map[string]string, offset, limit int) ([]models.Category, int64, error) {
	return cs.CategoryRepository.GetAll(queryParams, offset, limit)
}

func (cs *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return cs.CategoryRepository.GetByID(id)
}

func (cs *CategoryService) UpdateCategory(category *models.Category) error {
	return cs.CategoryRepository.Update(category)
}

func (cs *CategoryService) DeleteCategory(category *models.Category) error {
	return cs.CategoryRepository.Delete(category)
}
