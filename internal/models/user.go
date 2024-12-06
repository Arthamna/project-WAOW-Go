package models

import (
	"time"
	"gorm.io/gorm"
)

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_CUST  = "CUSTOMER"
)

type User struct {
    UserID           string         `gorm:"column:user_id;primaryKey"`
    Username         string         `gorm:"column:username;unique"`
    Email            string         `gorm:"column:email;unique"`
    PasswordHash     string         `gorm:"column:password_hash"`
    DisplayName      string         `gorm:"column:display_name"`
    Bio             string         `gorm:"column:bio"`
    ProfilePictureURL string         `gorm:"column:profile_picture_url"`
    RegistrationDate time.Time      `gorm:"column:registration_date"`
    Role            string         `gorm:"column:role"`
    CreatedAt       time.Time      `gorm:"column:created_at"`
    UpdatedAt       time.Time      `gorm:"column:updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
}
