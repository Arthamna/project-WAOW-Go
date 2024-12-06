package dtos

type CategoryCreateRequest struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}

type CategoryUpdateRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

type CategoryResponse struct {
    CategoryID  string `json:"category_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Slug        string `json:"slug"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

