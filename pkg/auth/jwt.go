package auth

import (
	"errors"
	"time"
	"project_article/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDFromToken(token *jwt.Token) (string, error)
	GetUserRoleFromToken(token *jwt.Token) (string, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    "your-app",
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(user *models.User) (string, error) {
	claims := &Claims{
		UserID: user.UserID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserIDFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user_id in token")
	}
	return userID, nil
}

func (j *jwtService) GetUserRoleFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("invalid role in token")
	}
	return role, nil
}