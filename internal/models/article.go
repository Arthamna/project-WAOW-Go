package models

import (
    "time"
    "gorm.io/gorm"
)

type Article struct {
    ArticleID    string         `gorm:"column:article_id;primaryKey"`
    CategoryID   string         `gorm:"column:category_id;index"` 
    AuthorID     string         `gorm:"column:author_id;index"`
    Title        string         `gorm:"column:title"`
    Content      string         `gorm:"column:content"`
    Slug         string         `gorm:"column:slug"`
    ViewCount    int            `gorm:"column:view_count"`
    CreatedAt    time.Time      `gorm:"column:created_at"`
    UpdatedAt    time.Time      `gorm:"column:updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
    Category     Category       `gorm:"belongsTo:Category;foreignKey:CategoryID;references:CategoryID"`//belongs so GORM Regcon
    Author       User          `gorm:"belongsTo:User;foreignKey:AuthorID;references:UserID"`
}