package dtos

type ArticleCreateRequest struct {
    Title      string `json:"title" binding:"required"`
    Content    string `json:"content" binding:"required"`
    CategoryID string `json:"category_id" binding:"required"`
}

type ArticleUpdateRequest struct {
    Title      string `json:"title"`
    Content    string `json:"content"`
    CategoryID string `json:"category_id"`
}

type ArticleResponse struct {
    ArticleID  string `json:"article_id"`
    Title      string `json:"title"`
    Content    string `json:"content"`
    Slug       string `json:"slug"`
    ViewCount  int    `json:"view_count"`
    CategoryID string `json:"category_id"`
    AuthorID   string `json:"author_id"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}