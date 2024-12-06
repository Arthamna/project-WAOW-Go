package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project_article/internal/dtos"
    "project_article/internal/services"
    "gorm.io/gorm"
)

func CreateArticle(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var request dtos.ArticleCreateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        userID, _ := c.Get("user_id")
        articleService := services.NewArticleService(db)
        article, err := articleService.Create(request, userID.(string))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, article)
    }
}

func GetArticles(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        articleService := services.NewArticleService(db)
        articles, err := articleService.GetAll()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, articles)
    }
}

func GetArticle(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        articleService := services.NewArticleService(db)
        article, err := articleService.GetByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
            return
        }

        c.JSON(http.StatusOK, article)
    }
}

func UpdateArticle(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var request dtos.ArticleUpdateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        userID, _ := c.Get("user_id")
        articleService := services.NewArticleService(db)
        article, err := articleService.Update(id, request, userID.(string))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, article)
    }
}

func DeleteArticle(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        userID, _ := c.Get("user_id")
        articleService := services.NewArticleService(db)
        err := articleService.Delete(id, userID.(string))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
    }
}
