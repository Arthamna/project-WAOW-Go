package services

import (
	// "errors"
	// "os"
	"project_article/internal/dtos"
	"project_article/internal/models"
	"project_article/internal/repositories"
	// "project_article/pkg/auth"

	// "project_article/pkg/auth/jwt"
	// "project_article/internal/services"
	// "project_article/pkg/common"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, input dtos.UserUpdateRequest) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepository repositories.UserRepository
}


func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}


func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) UpdateUser(id string, input dtos.UserUpdateRequest) (*models.User, error) {
	user, err := s.userRepository.FindByID(id)
	
	if err != nil {
		return nil, err
	}

	if input.Username != "" {
		user.Username = input.Username
	}
	if input.DisplayName != "" {
		user.DisplayName = input.DisplayName
	}
	if input.Bio != "" {
		user.Bio = input.Bio
	}
	if input.Role != "" {
		user.Role = input.Role
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            return nil, err
        }
        user.PasswordHash = string(hashedPassword)
    }

	user.UpdatedAt = time.Now()
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}

