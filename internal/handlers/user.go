package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project_article/internal/dtos"
    "project_article/internal/services"
    "project_article/internal/repositories"
    "gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRepo := repositories.NewUserRepository(db)
        jwtService := services.NewJWTService("your-secret-key")
        userService := services.NewUserService(*userRepo, jwtService) 
        users, err := userService.GetAllUsers()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, users)
    }
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {

        var request dtos.UserRegisterRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        input := dtos.RegisterUserInput{
            Username:    request.Username,
            Email:       request.Email,
            Password:    request.Password,
            DisplayName: request.DisplayName,
            Bio:         request.Bio,
        }

        userRepo := repositories.NewUserRepository(db)
        jwtService := services.NewJWTService("your-secret-key")
        userService := services.NewUserService(*userRepo, jwtService) 
        user, err := userService.Register(input)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, user)
    }
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var request dtos.UserUpdateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        input := dtos.UpdateUserInput{
            Username:    request.Username,
            DisplayName: request.DisplayName,
            Bio:         request.Bio,
            Role:        "", 
        }

        userRepo := repositories.NewUserRepository(db)
        jwtService := services.NewJWTService("your-secret-key")
        userService := services.NewUserService(*userRepo, jwtService) 
        user, err := userService.UpdateUser(id, input)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        userRepo := repositories.NewUserRepository(db)
        jwtService := services.NewJWTService("your-secret-key")
        userService := services.NewUserService(*userRepo, jwtService)  
        err := userService.DeleteUser(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    }
}

