package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "project_article/internal/handlers"
    "project_article/pkg/middleware"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    // Auth routes
    auth := r.Group("/auth")
    {
        auth.POST("/register", handlers.Register(db))
        auth.POST("/login", handlers.Login(db))
    }

    // Protected routes
    api := r.Group("/api")
    api.Use(middleware.AuthMiddleware())
    {
        // Category routes
        api.POST("/categories", handlers.CreateCategory(db))
        api.GET("/categories", handlers.GetCategories(db))
        api.GET("/categories/:id", handlers.GetCategory(db))
        api.PUT("/categories/:id", handlers.UpdateCategory(db))
        api.DELETE("/categories/:id", handlers.DeleteCategory(db))

        // Article routes
        api.POST("/articles", handlers.CreateArticle(db))
        api.GET("/articles", handlers.GetArticles(db))
        api.GET("/articles/:id", handlers.GetArticle(db))
        api.PUT("/articles/:id", handlers.UpdateArticle(db))
        api.DELETE("/articles/:id", handlers.DeleteArticle(db))
    }

    // Admin routes
    admin := r.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        admin.GET("/users", handlers.GetUsers(db))
        admin.POST("/users", handlers.CreateUser(db))
        admin.PUT("/users/:id", handlers.UpdateUser(db))
        admin.DELETE("/users/:id", handlers.DeleteUser(db))
    }
}
