package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// AuditLog adalah model untuk mencatat semua aktivitas pengguna.
type AuditLog struct {
	ID         uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID     *uuid.UUID      `gorm:"type:uuid;index" json:"user_id,omitempty"`
	Action     string          `gorm:"type:varchar(100);not null;index" json:"action"`
	Resource   string          `gorm:"type:varchar(100);not null;index" json:"resource"`
	ResourceID string          `gorm:"type:varchar(100)" json:"resource_id,omitempty"`
	Details    *AuditLogDetail `gorm:"type:jsonb" json:"details,omitempty"`
	IPAddress  string          `gorm:"type:varchar(45)" json:"ip_address,omitempty"`
	UserAgent  string          `gorm:"type:text" json:"user_agent,omitempty"`
	CreatedAt  time.Time       `gorm:"not null;default:NOW();index" json:"created_at"`
}

// TableName overrides the table name.
func (AuditLog) TableName() string {
	return "audit_logs"
}

// AuditLogDetail adalah tipe untuk menyimpan detail audit log dalam JSONB.
type AuditLogDetail map[string]interface{}

// Scan implements the Scanner interface for AuditLogDetail.
func (a *AuditLogDetail) Scan(value interface{}) error {
	if value == nil {
		*a = AuditLogDetail{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("unsupported type for AuditLogDetail")
	}
	return json.Unmarshal(bytes, a)
}

// Value implements the driver Valuer interface for AuditLogDetail.
func (a AuditLogDetail) Value() (driver.Value, error) {
	if a == nil {
		return json.Marshal(map[string]interface{}{})
	}
	return json.Marshal(a)
}
