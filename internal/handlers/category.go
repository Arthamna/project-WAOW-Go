package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project_article/internal/dtos"
    "project_article/internal/services"
    "gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var request dtos.CategoryCreateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        categoryService := services.NewCategoryService(db)
        category, err := categoryService.Create(request)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, category)
    }
}

func GetCategories(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        categoryService := services.NewCategoryService(db)
        categories, err := categoryService.GetAll()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, categories)
    }
}

func GetCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        categoryService := services.NewCategoryService(db)
        category, err := categoryService.GetByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }

        c.JSON(http.StatusOK, category)
    }
}

func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var request dtos.CategoryUpdateRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        categoryService := services.NewCategoryService(db)
        category, err := categoryService.Update(id, request)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, category)
    }
}

func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        categoryService := services.NewCategoryService(db)
        err := categoryService.Delete(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
    }
}
