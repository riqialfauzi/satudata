package domain

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel adalah base struct yang diembed di semua model.
type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `gorm:"not null;default:NOW()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null;default:NOW()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate akan generate UUID sebelum insert.
func (b *BaseModel) BeforeCreate() error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
