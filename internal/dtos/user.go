// internal/dtos/user.go
package dtos

import (
	"time"
	"project_article/internal/models"
)

type UserRegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	DisplayName string `json:"display_name" binding:"required"`
	Bio         string `json:"bio"`
}

type UserUpdateRequest struct {
	Username    string `json:"username" binding:""`
	Email       string `json:"email" binding:"email"`
	Password    string `json:"password" binding:""`
	DisplayName string `json:"display_name" binding:""`
	Bio         string `json:"bio"`
	Role        string `json:"role"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


type UserResponse struct {
	UserID            string    `json:"user_id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	DisplayName       string    `json:"display_name"`
	Bio               string    `json:"bio"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Role              string    `json:"role"`
	RegistrationDate  time.Time `json:"registration_date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}


type AuthResponse struct{
	User  UserResponse `json:"user"`
    Token string       `json:"token"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		UserID:            user.UserID,
		Username:          user.Username,
		Email:             user.Email,
		DisplayName:       user.DisplayName,
		Bio:              user.Bio,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:             user.Role,
		RegistrationDate: user.RegistrationDate,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}
}

func ToUserResponseList(users []models.User) []*UserResponse {
	var responses []*UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}
type AdminRegisterRequest struct {
    Username    string `json:"username" binding:"required"`
    Email       string `json:"email" binding:"required,email"`
    Password    string `json:"password" binding:"required,min=6"`
    DisplayName string `json:"display_name" binding:"required"`
    Bio         string `json:"bio"`
    SecretKey   string `json:"secret_key" binding:"required"` 
}
