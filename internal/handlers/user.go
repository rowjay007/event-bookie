package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H
// @Router /api/v1/users/create [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}


// GetAllUsers godoc
// @Summary Get all users with filtering, sorting, and pagination
// @Description Get all users with optional filtering, sorting, and pagination
// @Tags users
// @Param name query string false "Filter by name"
// @Param email query string false "Filter by email"
// @Param sort_by query string false "Sort by field (e.g., name, email)"
// @Param sort_order query string false "Sort order (asc or desc)"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} map[string]interface{} "Users and total count"
// @Failure 500 {object} gin.H "Failed to fetch users"
// @Router /api/v1/users [get]
func (uh *UserHandler) GetAllUsers(c *gin.Context) {
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

	users, total, err := uh.UserService.GetAllUsers(params, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	response := gin.H{
		"users": users,
		"total": total,
	}

	c.JSON(http.StatusOK, response)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} gin.H
// @Router /api/v1/users/{id} [get]
func (uh *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := uh.UserService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H
// @Router /api/v1/users/{id} [put]
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.UserService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Failure 404 {object} gin.H
// @Router /api/v1/users/{id} [delete]
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.UserService.DeleteUser(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Login godoc
// @Summary User login
// @Description User login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login"
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /api/v1/auth/login [post]
func (uh *UserHandler) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uh.UserService.AuthenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := uh.UserService.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Forgot password
// @Tags Auth
// @Accept json
// @Produce json
// @Param forgotPassword body ForgotPasswordRequest true "Forgot Password"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/v1/auth/forgot-password [post]
func (uh *UserHandler) ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest ForgotPasswordRequest
	if err := c.ShouldBindJSON(&forgotPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uh.UserService.ForgotPassword(forgotPasswordRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reset_token": token})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags Auth
// @Accept json
// @Produce json
// @Param resetPassword body ResetPasswordRequest true "Reset Password"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/v1/auth/reset-password [post]
func (uh *UserHandler) ResetPassword(c *gin.Context) {
	var resetPasswordRequest ResetPasswordRequest
	if err := c.ShouldBindJSON(&resetPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.UserService.ResetPassword(resetPasswordRequest.Email, resetPasswordRequest.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// Logout godoc
// @Summary User logout
// @Description User logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Router /api/v1/auth/logout [post]
func (uh *UserHandler) Logout(c *gin.Context) {
	// For stateless JWT, logout can be a client-side operation.
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// GetUserProfile godoc
// @Summary Get user profile
// @Description Get user profile
// @Tags Secured
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} gin.H
// @Router /api/v1/secured/profile [get]
func (uh *UserHandler) GetUserProfile(c *gin.Context) {
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User email not found"})
		return
	}

	user, err := uh.UserService.GetUserByEmail(userEmail.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}
