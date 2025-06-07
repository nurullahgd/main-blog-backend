package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string         `json:"name" gorm:"not null"`
	Surname      string         `json:"surname" gorm:"not null"`
	Username     string         `json:"username" gorm:"not null;unique"`
	Email        string         `json:"email" gorm:"not null;unique"`
	Password     []byte         `json:"password" gorm:"not null"`
	ProfileImage string         `json:"profile_image" gorm:"default:null"`
	BlogCount    int            `json:"blog_count" gorm:"default:0"`
	Blogs        []Blog         `json:"blogs,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// UserCreate represents the data needed to create a new user
type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse represents the user data that will be sent in responses
type UserResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	ProfileImage string    `json:"profile_image"`
	BlogCount    int       `json:"blog_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserEdit struct {
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	ProfileImage string `json:"profile_image" binding:"required"`
	UpdatedAt    int64  `json:"updated_at" binding:"required"`
}
type UserPasswordChange struct {
	Password string `json:"password" binding:"required"`
}
