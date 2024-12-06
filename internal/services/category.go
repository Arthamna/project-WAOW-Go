package services

import (
    "github.com/google/uuid"
    "github.com/gosimple/slug"
    "gorm.io/gorm"
    "project_article/internal/dtos"
    "project_article/internal/models"
    "project_article/internal/repositories"
    "time"
)

type CategoryService struct {
    categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(db *gorm.DB) *CategoryService {
    return &CategoryService{
        categoryRepo: repositories.NewCategoryRepository(db),
    }
}

func (s *CategoryService) Create(req dtos.CategoryCreateRequest) (*dtos.CategoryResponse, error) {
    category := &models.Category{
        CategoryID:  uuid.New().String(),
        Name:        req.Name,
        Description: req.Description,
        Slug:        slug.Make(req.Name),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    if err := s.categoryRepo.Create(category); err != nil {
        return nil, err
    }

    return &dtos.CategoryResponse{
        CategoryID:  category.CategoryID,
        Name:        category.Name,
        Description: category.Description,
        Slug:        category.Slug,
        CreatedAt:   category.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   category.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *CategoryService) GetAll() ([]dtos.CategoryResponse, error) {
    categories, err := s.categoryRepo.FindAll()
    if err != nil {
        return nil, err
    }

    var response []dtos.CategoryResponse
    for _, category := range categories {
        response = append(response, dtos.CategoryResponse{
            CategoryID:  category.CategoryID,
            Name:        category.Name,
            Description: category.Description,
            Slug:        category.Slug,
            CreatedAt:   category.CreatedAt.Format(time.RFC3339),
            UpdatedAt:   category.UpdatedAt.Format(time.RFC3339),
        })
    }

    return response, nil
}

func (s *CategoryService) GetByID(id string) (*dtos.CategoryResponse, error) {
    category, err := s.categoryRepo.FindByID(id)
    if err != nil {
        return nil, err
    }

    return &dtos.CategoryResponse{
        CategoryID:  category.CategoryID,
        Name:        category.Name,
        Description: category.Description,
        Slug:        category.Slug,
        CreatedAt:   category.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   category.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *CategoryService) Update(id string, req dtos.CategoryUpdateRequest) (*dtos.CategoryResponse, error) {
    category, err := s.categoryRepo.FindByID(id)
    if err != nil {
        return nil, err
    }

    if req.Name != "" {
        category.Name = req.Name
        category.Slug = slug.Make(req.Name)
    }
    if req.Description != "" {
        category.Description = req.Description
    }
    category.UpdatedAt = time.Now()

    if err := s.categoryRepo.Update(category); err != nil {
        return nil, err
    }

    return &dtos.CategoryResponse{
        CategoryID:  category.CategoryID,
        Name:        category.Name,
        Description: category.Description,
        Slug:        category.Slug,
        CreatedAt:   category.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   category.UpdatedAt.Format(time.RFC3339),
    }, nil
}

func (s *CategoryService) Delete(id string) error {
    return s.categoryRepo.Delete(id)
}