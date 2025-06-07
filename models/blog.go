package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title      string         `json:"title" gorm:"not null"`
	Content    string         `json:"content" gorm:"type:text;not null"`
	Slug       string         `json:"slug" gorm:"not null;unique"`
	MainImage  string         `json:"main_image" gorm:"default:null"`
	UserID     string         `json:"user_id" gorm:"type:uuid;not nullc"`
	Visibility bool           `json:"visibility" gorm:"default:true"`
	Category   string         `json:"category" gorm:"not null"`
	Summary    string         `json:"summary" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// BlogCreate represents the data needed to create a new blog
type BlogCreate struct {
	Title      string    `json:"title" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	Slug       string    `json:"slug" binding:"required"`
	MainImage  string    `json:"main_image"`
	UserID     string    `json:"user_id" binding:"required"`
	Visibility bool      `json:"visibility" binding:"required"`
	Category   string    `json:"category" binding:"required"`
	Summary    string    `json:"summary" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// BlogResponse represents the blog data that will be sent in responses
type BlogResponse struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Slug       string    `json:"slug"`
	MainImage  string    `json:"main_image"`
	UserID     string    `json:"user_id"`
	Category   string    `json:"category"`
	Visibility bool      `json:"visibility"`
	Summary    string    `json:"summary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
