package repositories

import (
    "gorm.io/gorm"
    "project_article/internal/models"
)

type ArticleRepository struct {
    db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
    return &ArticleRepository{db}
}

func (r *ArticleRepository) Create(article *models.Article) error {
    return r.db.Create(article).Error
}

func (r *ArticleRepository) FindAll() ([]models.Article, error) {
    var articles []models.Article
    err := r.db.Preload("Category").Preload("Author").Find(&articles).Error
    return articles, err
}

func (r *ArticleRepository) FindByID(id string) (*models.Article, error) {
    var article models.Article
    err := r.db.Preload("Category").Preload("Author").First(&article, "article_id = ?", id).Error
    return &article, err
}

func (r *ArticleRepository) Update(article *models.Article) error {
    return r.db.Save(article).Error
}

func (r *ArticleRepository) Delete(id string) error {
    return r.db.Delete(&models.Article{}, "article_id = ?", id).Error
}