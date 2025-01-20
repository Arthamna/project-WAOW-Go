package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project_article/internal/dtos"
    "project_article/internal/services"
    "project_article/internal/repositories"
    "gorm.io/gorm"
)

// type UserHandler struct {
//     userService services.UserService
// }

func GetUsers(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRepository := repositories.NewUserRepository(db)
        userService := services.NewUserService(*userRepository)
        users, err := services.UserService.GetAllUsers(userService)
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
        
        authService := services.NewAuthService(db)
        user, err := authService.Register(request)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, user)
    }
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRepository := repositories.NewUserRepository(db)
        userService := services.NewUserService(*userRepository)
        id := c.Param("id")
        var request dtos.UserUpdateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        input := dtos.UserUpdateRequest(request)

        user, err := services.UserService.UpdateUser(userService ,id, input)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRepository := repositories.NewUserRepository(db)
        userService := services.NewUserService(*userRepository)
        id := c.Param("id")
        err := services.UserService.DeleteUser(userService,id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    }
}
