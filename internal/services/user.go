package services

import (
	"errors"
	"os"
	"project_article/internal/dtos"
	"project_article/internal/models"
	"project_article/internal/repositories"
	"project_article/pkg/auth"
	// "project_article/pkg/auth/jwt"
	// "project_article/internal/services"
	// "project_article/pkg/common"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input dtos.RegisterUserInput) (*models.User, error)
	Login(input dtos.LoginInput) (*dtos.LoginResponse, error)
	GetAllUsers() ([]models.User, error)
	// GetAll() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, input dtos.UpdateUserInput) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepository repositories.UserRepository
	jwtAuth        auth.JWTService
}

// func NewJWTService(secretKey string) *auth.JWTService {
// 	secretKey = os.Getenv("JWT_SECRET_KEY")
// 	 auth.NewJWTService(secretKey)
//     return &auth.JWTService{secretKey: secretKey}
// }
func NewJWTService(secretKey string) auth.JWTService {
	secretKey = os.Getenv("JWT_SECRET_KEY")
	return auth.NewJWTService(secretKey)
}

func NewUserService(userRepo repositories.UserRepository, jwtAuth auth.JWTService) UserService {
	// userService := NewUserService(*repositories.NewUserRepository(db), auth.NewJWTService()) //perbaiki ini
	return &userService{
		userRepository: userRepo,
		jwtAuth:        jwtAuth,
	}
}

func (s *userService) Register(input dtos.RegisterUserInput) (user *models.User, err error) {
	existingUser, _ := s.userRepository.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user = &models.User{
		Username:         input.Username,
		Email:           input.Email,
		PasswordHash:    string(hashedPassword),
		DisplayName:     input.DisplayName,
		Bio:             input.Bio,
		ProfilePictureURL: "#",
		Role:            "USER", 
		RegistrationDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return s.userRepository.Create(user)
}

func (s *userService) Login(input dtos.LoginInput) (*dtos.LoginResponse, error) {
	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtAuth.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) UpdateUser(id string, input dtos.UpdateUserInput) (*models.User, error) {
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

	user.UpdatedAt = time.Now()
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}

