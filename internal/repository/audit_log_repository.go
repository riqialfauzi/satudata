package repository

import (
	"context"
	"fmt"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// AuditLogRepository adalah implementasi dari AuditLogRepositoryInterface.
type AuditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository membuat instance baru AuditLogRepository.
func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

// Create mencatat audit log baru.
func (r *AuditLogRepository) Create(ctx context.Context, log *domain.AuditLog) error {
	if err := r.db.WithContext(ctx).Create(log).Error; err != nil {
		return fmt.Errorf("failed to create audit log: %w", err)
	}
	return nil
}

// GetAuditLogs mengambil daftar audit logs dengan filter.
func (r *AuditLogRepository) GetAuditLogs(ctx context.Context, filter AuditLogFilter) ([]domain.AuditLog, int64, error) {
	query := r.db.WithContext(ctx).Model(&domain.AuditLog{})

	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.Action != "" {
		query = query.Where("action = ?", filter.Action)
	}
	if filter.Resource != "" {
		query = query.Where("resource = ?", filter.Resource)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count audit logs: %w", err)
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	var logs []domain.AuditLog
	if err := query.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&logs).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get audit logs: %w", err)
	}

	return logs, total, nil
}
