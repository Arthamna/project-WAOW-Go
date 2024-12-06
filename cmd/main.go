package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "project_article/internal/routes"
    "project_article/pkg/database"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("JWT_SECRET_KEY tidak ditemukan di environment variables")
    }
// jwtService := auth.NewJWTService(secretKey)

    db := database.ConnectToPostgresql()
    
    r := gin.Default()
    
    routes.SetupRoutes(r, db)
    
    r.Run(":8080")
}
