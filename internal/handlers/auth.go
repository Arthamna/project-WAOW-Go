package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project_article/internal/dtos"
    "project_article/internal/services"
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
        token, err := authService.Login(request)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}
