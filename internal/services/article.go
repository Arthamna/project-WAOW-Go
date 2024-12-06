package services

import (
    "errors"
    "github.com/google/uuid"
    "github.com/gosimple/slug"
    "gorm.io/gorm"
    "project_article/internal/dtos"
    "project_article/internal/models"
    "project_article/internal/repositories"
    "time"
)

type ArticleService struct {
    articleRepo *repositories.ArticleRepository
}

func NewArticleService(db *gorm.DB) *ArticleService {
    return &ArticleService{
        articleRepo: repositories.NewArticleRepository(db),
    }
}

func (s *ArticleService) Create(req dtos.ArticleCreateRequest, authorID string) (*dtos.ArticleResponse, error) {
    article := &models.Article{
        ArticleID:  uuid.New().String(),
        Title:      req.Title,
        Content:    req.Content,
        Slug:       slug.Make(req.Title),
        CategoryID: req.CategoryID,
        AuthorID:   authorID,
        ViewCount:  0,
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
    }

    if err := s.articleRepo.Create(article); err != nil {
        return nil, err
    }

    return &dtos.ArticleResponse{
        ArticleID:  article.ArticleID,
        Title:      article.Title,
        Content:    article.Content,
        Slug:       article.Slug,
        ViewCount:  article.ViewCount,
        CategoryID: article.CategoryID,
        AuthorID:   article.AuthorID,
        CreatedAt:  article.CreatedAt.Format(time.RFC3339),
        UpdatedAt:  article.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *ArticleService) GetAll() ([]dtos.ArticleResponse, error) {
    articles, err := s.articleRepo.FindAll()
    if err != nil {
        return nil, err
    }

    var response []dtos.ArticleResponse
    for _, article := range articles {
        response = append(response, dtos.ArticleResponse{
            ArticleID:  article.ArticleID,
            Title:      article.Title,
            Content:    article.Content,
            Slug:       article.Slug,
            ViewCount:  article.ViewCount,
            CategoryID: article.CategoryID,
            AuthorID:   article.AuthorID,
            CreatedAt:  article.CreatedAt.Format(time.RFC3339),
            UpdatedAt:  article.UpdatedAt.Format(time.RFC3339),
        })
    }

    return response, nil
}

func (s *ArticleService) GetByID(id string) (*dtos.ArticleResponse, error) {
    article, err := s.articleRepo.FindByID(id)
    if err != nil {
        return nil, err
    }

    return &dtos.ArticleResponse{
        ArticleID:  article.ArticleID,
        Title:      article.Title,
        Content:    article.Content,
        Slug:       article.Slug,
        ViewCount:  article.ViewCount,
        CategoryID: article.CategoryID,
        AuthorID:   article.AuthorID,
        CreatedAt:  article.CreatedAt.Format(time.RFC3339),
        UpdatedAt:  article.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *ArticleService) Update(id string, req dtos.ArticleUpdateRequest, userID string) (*dtos.ArticleResponse, error) {
    article, err := s.articleRepo.FindByID(id)
    if err != nil {
        return nil, err
    }

    if article.AuthorID != userID {
        return nil, errors.New("unauthorized to update this article")
    }

    if req.Title != "" {
        article.Title = req.Title
        article.Slug = slug.Make(req.Title)
    }
    if req.Content != "" {
        article.Content = req.Content
    }
    if req.CategoryID != "" {
        article.CategoryID = req.CategoryID
    }
    article.UpdatedAt = time.Now()

    if err := s.articleRepo.Update(article); err != nil {
        return nil, err
    }

    return &dtos.ArticleResponse{
        ArticleID:  article.ArticleID,
        Title:      article.Title,
        Content:    article.Content,
        Slug:       article.Slug,
        ViewCount:  article.ViewCount,
        CategoryID: article.CategoryID,
        AuthorID:   article.AuthorID,
        CreatedAt:  article.CreatedAt.Format(time.RFC3339),
        UpdatedAt:  article.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *ArticleService) Delete(id string, userID string) error {
    article, err := s.articleRepo.FindByID(id)
    if err != nil {
        return err
    }

    if article.AuthorID != userID {
        return errors.New("unauthorized to delete this article")
    }

    return s.articleRepo.Delete(id)
}