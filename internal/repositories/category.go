package repositories

import (
    "gorm.io/gorm"
    "project_article/internal/models"
)

type CategoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
    return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
    return r.db.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
    var categories []models.Category
    err := r.db.Find(&categories).Error
    return categories, err
}

func (r *CategoryRepository) FindByID(id string) (*models.Category, error) {
    var category models.Category
    err := r.db.First(&category, "category_id = ?", id).Error
    return &category, err
}

func (r *CategoryRepository) Update(category *models.Category) error {
    return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id string) error {
    return r.db.Delete(&models.Category{}, "category_id = ?", id).Error
}