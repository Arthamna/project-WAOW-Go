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
    jwtService auth.JWTService
}

func NewAuthService(db *gorm.DB) *AuthService {
    secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("JWT_SECRET_KEY tidak ditemukan di environment variables")
    }
    
    return &AuthService{
        userRepo:    repositories.NewUserRepository(db),
        jwtService:  auth.NewJWTService(secretKey),
    }
}

func (s *AuthService) Register(req dtos.UserRegisterRequest) (*dtos.AuthResponse, error) {
    existingUser, _ := s.userRepo.FindByEmail(req.Email)
    if existingUser != nil {
        return nil, errors.New("email already registered")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Buat user baru
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

    createdUser, err := s.userRepo.Create(user)
    if err != nil {
        return nil, err
    }

    token, err := s.jwtService.GenerateToken(user)
    if err != nil {
        return nil, err
    }

    return &dtos.AuthResponse{
        User:  *dtos.ToUserResponse(createdUser),
        Token: token,
    }, nil
}

func (s *AuthService) RegisterAdmin(input dtos.AdminRegisterRequest) (*dtos.AuthResponse, error) {
    expectedKey := os.Getenv("ADMIN_SECRET_KEY")
    if expectedKey == "" {
        return nil, errors.New("admin secret key not configured")
    }
    if input.SecretKey != expectedKey {
        return nil, errors.New("invalid admin secret key")
    }

    existingUser, _ := s.userRepo.FindByEmail(input.Email)
    if existingUser != nil {
        return nil, errors.New("email already registered")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        UserID:           uuid.New().String(),
        Username:         input.Username,
        Email:           input.Email,
        PasswordHash:    string(hashedPassword),
        DisplayName:     input.DisplayName,
        Bio:             input.Bio,
        ProfilePictureURL: "#",
        Role:            "ADMIN",  
        RegistrationDate: time.Now(),
        CreatedAt:       time.Now(),
        UpdatedAt:       time.Now(),
    }

    createdUser, err := s.userRepo.Create(user)
    if err != nil {
        return nil, err
    }

    token, err := s.jwtService.GenerateToken(user)
    if err != nil {
        return nil, err
    }

    return &dtos.AuthResponse{
        User:  *dtos.ToUserResponse(createdUser),
        Token: token,
    }, nil
}

func (s *AuthService) Login(req dtos.UserLoginRequest) (*dtos.AuthResponse, error) {
    user, err := s.userRepo.FindByEmail(req.Email)
    if err != nil {
        return nil, errors.New("invalid email")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("invalid password")
    }

    token, err := s.jwtService.GenerateToken(user)
    if err != nil {
        return nil, err
    }

    return &dtos.AuthResponse{
        User:  *dtos.ToUserResponse(user),
        Token: token,
    }, nil
}