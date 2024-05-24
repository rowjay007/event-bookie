package service

import (
	"errors"

	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
	"github.com/rowjay007/event-bookie/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (us *UserService) CreateUser(user *models.User) error {
	existingUser, err := us.UserRepository.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return us.UserRepository.Create(user)
}

func (us *UserService) GetAllUsers(queryParams map[string]string, offset, limit int) ([]models.User, int64, error) {
	return us.UserRepository.GetAll(queryParams, offset, limit)
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.UserRepository.GetByID(id)
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	return us.UserRepository.GetByEmail(email)
}

func (us *UserService) UpdateUser(user *models.User) error {
	return us.UserRepository.Update(user)
}

func (us *UserService) DeleteUser(user *models.User) error {
	return us.UserRepository.Delete(user)
}

func (us *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := us.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (us *UserService) GenerateJWT(email, role string) (string, error) {
	return utils.GenerateJWT(email, role)
}

func (us *UserService) ForgotPassword(email string) (string, error) {
	_, err := us.UserRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	resetToken, err := utils.GenerateResetToken()
	if err != nil {
		return "", err
	}
	return resetToken, nil
}

func (us *UserService) ResetPassword(email, newPassword string) error {
	user, err := us.UserRepository.GetByEmail(email)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return us.UserRepository.Update(user)
}
