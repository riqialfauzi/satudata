package domain

import (
	"time"
)

// UserRole mendefinisikan role pengguna.
type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleEditor UserRole = "editor"
	UserRoleViewer UserRole = "viewer"
)

// IsValid memeriksa apakah UserRole valid.
func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleEditor, UserRoleViewer:
		return true
	}
	return false
}

// User adalah model untuk autentikasi dan otorisasi pengguna.
type User struct {
	BaseModel
	Email        string     `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"-"`
	FullName     string     `gorm:"type:varchar(255);not null" json:"full_name"`
	Role         UserRole   `gorm:"type:varchar(50);not null;default:editor;index" json:"role"`
	IsActive     bool       `gorm:"not null;default:true" json:"is_active"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
}

// TableName overrides the table name.
func (User) TableName() string {
	return "users"
}

// IsAdmin memeriksa apakah user memiliki role admin.
func (u *User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

// IsEditor memeriksa apakah user memiliki role editor atau admin.
func (u *User) IsEditor() bool {
	return u.Role == UserRoleEditor || u.Role == UserRoleAdmin
}
