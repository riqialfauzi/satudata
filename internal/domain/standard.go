package domain

import (
	"github.com/google/uuid"
)

// StandardStatus mendefinisikan status standar data.
type StandardStatus string

const (
	StandardStatusActive   StandardStatus = "active"
	StandardStatusArchived StandardStatus = "archived"
	StandardStatusDraft    StandardStatus = "draft"
)

// IsValid memeriksa apakah StandardStatus valid.
func (s StandardStatus) IsValid() bool {
	switch s {
	case StandardStatusActive, StandardStatusArchived, StandardStatusDraft:
		return true
	}
	return false
}

// Standard adalah model untuk standar data tahunan.
type Standard struct {
	BaseModel
	Title       string         `gorm:"type:varchar(500);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	Year        int            `gorm:"not null;index" json:"year"`
	FileURL     string         `gorm:"type:text" json:"file_url,omitempty"`
	FileSize    int64          `gorm:"default:0" json:"file_size,omitempty"`
	Status      StandardStatus `gorm:"type:varchar(50);not null;default:active;index" json:"status"`
	Version     string         `gorm:"type:varchar(50);not null;default:1.0" json:"version"`
	IsCurrent   bool           `gorm:"not null;default:false;index" json:"is_current"`
	CreatedBy   *uuid.UUID     `gorm:"type:uuid;index" json:"created_by,omitempty"`
}

// TableName overrides the table name.
func (Standard) TableName() string {
	return "standards"
}
