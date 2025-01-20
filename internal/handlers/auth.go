package handlers

import (
	// "go/token"
	"net/http"
	"project_article/internal/dtos"
	"project_article/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var request dtos.UserRegisterRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        authService := services.NewAuthService(db)
        user, err := authService.Register(request)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, user)
    }
}

func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var request dtos.UserLoginRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        authService := services.NewAuthService(db)
        user, err := authService.Login(request)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}

func RegisterAdmin(db *gorm.DB) gin.HandlerFunc {
    return func (c *gin.Context){
        var request dtos.AdminRegisterRequest 
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        authService := services.NewAuthService(db)
        user, err := authService.RegisterAdmin(request)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, user)
    }
}