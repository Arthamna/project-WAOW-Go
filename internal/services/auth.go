package services

import (
	"errors"
	"log"
	"os"
	"project_article/internal/dtos"
	"project_article/internal/models"
	"project_article/internal/repositories"
	"project_article/pkg/auth"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
    userRepo *repositories.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{
        userRepo: repositories.NewUserRepository(db),
    }
}

func (s *AuthService) Register(req dtos.UserRegisterRequest) (*dtos.UserResponse, error) {
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        UserID:           uuid.New().String(),
        Username:         req.Username,
        Email:            req.Email,
        PasswordHash:     string(hashedPassword),
        DisplayName:      req.DisplayName,
        ProfilePictureURL: "#",
        Role:            "USER",
        RegistrationDate: time.Now(),
        CreatedAt:       time.Now(),
        UpdatedAt:       time.Now(),
    }

    if _, err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return &dtos.UserResponse{
        UserID:      user.UserID,
        Username:    user.Username,
        Email:       user.Email,
        DisplayName: user.DisplayName,
        Role:        user.Role,
    }, nil
}

func (s *AuthService) Login(req dtos.UserLoginRequest) (string, error) {
    user, err := s.userRepo.FindByEmail(req.Email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return "", errors.New("invalid credentials")
    }
	secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("JWT_SECRET_KEY tidak ditemukan di environment variables")
    }
	jwtService := auth.NewJWTService(secretKey) 
    token, err := jwtService.GenerateToken(user)
    if err != nil {
        return "", err
    }

    return token, nil

    // return auth.GenerateToken(user.UserID, user.Role)
}