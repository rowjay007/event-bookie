package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/service"
)

// UserHandler handles HTTP requests related to user management
type UserHandler struct {
    UserService *service.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{UserService: userService}
}

// CreateUser handles the creation of a new user
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Accept json
// @Produce json
// @Param user body models.User true "User object to be created"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := uh.UserService.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

// GetUserByID retrieves a user by its ID
// @Summary Get a user by ID
// @Description Retrieve a user by its ID
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /users/{id} [get]
func (uh *UserHandler) GetUserByID(c *gin.Context) {
    userID := c.Param("id")

    // Parse userID to uint
    id, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Convert uint64 to uint
    uid := uint(id)

    user, err := uh.UserService.GetUserByID(uid)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing user
// @Summary Update a user
// @Description Update an existing user with the provided details
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User object with updated details"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func (uh *UserHandler) UpdateUser(c *gin.Context) {
    userID := c.Param("id")

    // Parse userID to uint
    id, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Convert uint64 to uint
    uid := uint(id)

    // Fetch the user from the database
    user, err := uh.UserService.GetUserByID(uid)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Bind JSON request body to user struct
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the user in the database
    if err := uh.UserService.UpdateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user by its ID
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (uh *UserHandler) DeleteUser(c *gin.Context) {
    userID := c.Param("id")

    // Parse userID to uint
    id, err := strconv.ParseUint(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Convert uint64 to uint
    uid := uint(id)

    // Fetch the user from the database
    user, err := uh.UserService.GetUserByID(uid)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Delete the user from the database
    err = uh.UserService.DeleteUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
