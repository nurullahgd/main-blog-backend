package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminUser struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string         `json:"username" gorm:"not null;unique"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"-" gorm:"not null"`                            // Password is not exposed in JSON
	Role      string         `json:"role" gorm:"type:varchar(20);default:'admin'"` // admin, super_admin
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// AdminUserCreate represents the data needed to create a new admin user
type AdminUserCreate struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin super_admin"`
}

// AdminUserResponse represents the admin user data that will be sent in responses
type AdminUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}
