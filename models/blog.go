package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	Slug      string         `json:"slug" gorm:"not null;unique"`
	MainImage string         `json:"main_image" gorm:"default:null"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	User      User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Status    string         `json:"status" gorm:"type:varchar(20);default:'draft'"` // draft, published, archived
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// BlogCreate represents the data needed to create a new blog
type BlogCreate struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Slug      string    `json:"slug" binding:"required"`
	MainImage string    `json:"main_image"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	Status    string    `json:"status" binding:"required,oneof=draft published archived"`
}

// BlogResponse represents the blog data that will be sent in responses
type BlogResponse struct {
	ID        uuid.UUID    `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Slug      string       `json:"slug"`
	MainImage string       `json:"main_image"`
	UserID    uuid.UUID    `json:"user_id"`
	User      UserResponse `json:"user,omitempty"`
	Status    string       `json:"status"`
	CreatedAt int64        `json:"created_at"`
	UpdatedAt int64        `json:"updated_at"`
}
