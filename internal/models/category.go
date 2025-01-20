package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
    CategoryID   string         `gorm:"column:category_id;primaryKey"`
    Name         string         `gorm:"column:name"`
    Description  string         `gorm:"column:description"`
    Slug         string         `gorm:"column:slug"`
    CreatedAt    time.Time      `gorm:"column:created_at"`
    UpdatedAt    time.Time      `gorm:"column:updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
    Articles     []Article      `gorm:"foreignKey:CategoryID;references:CategoryID"` 
}

