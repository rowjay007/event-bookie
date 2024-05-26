package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category"
// @Success 201 {object} models.Category
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/categories [post]
func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := ch.CategoryService.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// GetAllCategories godoc
// @Summary Get all categories with filtering, sorting, and pagination
// @Description Get all categories with optional filtering, sorting, and pagination
// @Tags categories
// @Param name query string false "Filter by name"
// @Param sort_by query string false "Sort by field (e.g., name)"
// @Param sort_order query string false "Sort order (asc or desc)"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} map[string]interface{} "Categories and total count"
// @Failure 500 {object} gin.H "Failed to fetch categories"
// @Router /api/v1/categories [get]
func (ch *CategoryHandler) GetAllCategories(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	offsetStr := queryParams.Get("offset")
	limitStr := queryParams.Get("limit")

	var offset, limit int
	var err error

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	} else {
		offset = -1
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	} else {
		limit = -1
	}

	params := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	categories, total, err := ch.CategoryService.GetAllCategories(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	response := gin.H{
		"categories": categories,
		"total":      total,
	}

	// Send response
	c.JSON(http.StatusOK, response)
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Get a category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} gin.H "Category not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /api/v1/categories/{id} [get]
func (ch *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := ch.CategoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category"
// @Success 200 {object} models.Category
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 404 {object} gin.H "Category not found"
// @Failure 500 {object} gin.H "Failed to update category"
// @Router /api/v1/categories/{id} [put]
func (ch *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	category.ID = uint(id)
	if err := ch.CategoryService.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Param id path int true "Category ID"
// @Success 204 {object} gin.H "Category deleted"
// @Failure 404 {object} gin.H "Category not found"
// @Failure 500 {object} gin.H "Failed to delete category"
// @Router /api/v1/categories/{id} [delete]
func (ch *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := ch.CategoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := ch.CategoryService.DeleteCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Category deleted"})
}
