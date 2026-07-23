package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ReleaseType mendefinisikan tipe rilis.
type ReleaseType string

const (
	ReleaseTypeDataset     ReleaseType = "dataset"
	ReleaseTypeArticle     ReleaseType = "article"
	ReleaseTypeInfographic ReleaseType = "infographic"
)

// IsValid memeriksa apakah ReleaseType valid.
func (r ReleaseType) IsValid() bool {
	switch r {
	case ReleaseTypeDataset, ReleaseTypeArticle, ReleaseTypeInfographic:
		return true
	}
	return false
}

// ReleaseStatus mendefinisikan status rilis.
type ReleaseStatus string

const (
	ReleaseStatusDraft     ReleaseStatus = "draft"
	ReleaseStatusPublished ReleaseStatus = "published"
	ReleaseStatusArchived  ReleaseStatus = "archived"
)

// IsValid memeriksa apakah ReleaseStatus valid.
func (r ReleaseStatus) IsValid() bool {
	switch r {
	case ReleaseStatusDraft, ReleaseStatusPublished, ReleaseStatusArchived:
		return true
	}
	return false
}

// Release adalah model untuk rilis data (dataset, artikel, infographic).
type Release struct {
	BaseModel
	Title         string        `gorm:"type:varchar(500);not null" json:"title"`
	Slug          string        `gorm:"type:varchar(500);not null;uniqueIndex" json:"slug"`
	Description   string        `gorm:"type:text" json:"description,omitempty"`
	ReleaseType   ReleaseType   `gorm:"type:varchar(50);not null;index" json:"release_type"`
	Status        ReleaseStatus `gorm:"type:varchar(50);not null;default:draft;index" json:"status"`
	Year          int           `gorm:"not null;index" json:"year"`
	CoverImageURL string        `gorm:"type:text" json:"cover_image_url,omitempty"`
	Tags          TagsJSON      `gorm:"type:text[]" json:"tags,omitempty"`
	ViewCount     int           `gorm:"not null;default:0" json:"view_count"`
	PublishedAt   *time.Time    `json:"published_at,omitempty"`
	CreatedBy     *uuid.UUID    `gorm:"type:uuid;index" json:"created_by,omitempty"`

	// Relations
	DatasetMetadata *DatasetMetadata `gorm:"foreignKey:ReleaseID" json:"dataset_metadata,omitempty"`
	ArticleMetadata *ArticleMetadata `gorm:"foreignKey:ReleaseID" json:"article_metadata,omitempty"`
}

// TableName overrides the table name.
func (Release) TableName() string {
	return "releases"
}

// TagsJSON is a custom type for storing tags as PostgreSQL text array.
type TagsJSON []string

// Scan implements the Scanner interface for TagsJSON.
func (t *TagsJSON) Scan(value interface{}) error {
	if value == nil {
		*t = TagsJSON{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, t)
	case string:
		return json.Unmarshal([]byte(v), t)
	default:
		return errors.New("unsupported type for TagsJSON")
	}
}

// Value implements the driver Valuer interface for TagsJSON.
func (t TagsJSON) Value() (driver.Value, error) {
	if t == nil {
		return "{}", nil
	}
	return json.Marshal(t)
}

// DatasetMetadata adalah metadata khusus untuk rilis tipe dataset.
type DatasetMetadata struct {
	ID              uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ReleaseID       uuid.UUID    `gorm:"type:uuid;not null;index" json:"release_id"`
	FileURL         string       `gorm:"type:text;not null" json:"file_url"`
	FileFormat      string       `gorm:"type:varchar(50);not null" json:"file_format"`
	FileSize        int64        `gorm:"not null;default:0" json:"file_size"`
	RowCount        *int         `json:"row_count,omitempty"`
	ColumnCount     *int         `json:"column_count,omitempty"`
	Columns         *ColumnsJSON `gorm:"type:jsonb" json:"columns,omitempty"`
	DataSource      string       `gorm:"type:varchar(500)" json:"data_source,omitempty"`
	DataPeriodStart *time.Time   `json:"data_period_start,omitempty"`
	DataPeriodEnd   *time.Time   `json:"data_period_end,omitempty"`
	UpdateFrequency string       `gorm:"type:varchar(100)" json:"update_frequency,omitempty"`
	IsGeospatial    bool         `gorm:"not null;default:false" json:"is_geospatial"`
	CreatedAt       time.Time    `gorm:"not null;default:NOW()" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"not null;default:NOW()" json:"updated_at"`
}

// TableName overrides the table name.
func (DatasetMetadata) TableName() string {
	return "dataset_metadata"
}

// ColumnsJSON is a custom type for storing columns metadata as JSONB.
type ColumnsJSON []map[string]interface{}

// Scan implements the Scanner interface for ColumnsJSON.
func (c *ColumnsJSON) Scan(value interface{}) error {
	if value == nil {
		*c = ColumnsJSON{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("unsupported type for ColumnsJSON")
	}
	return json.Unmarshal(bytes, c)
}

// Value implements the driver Valuer interface for ColumnsJSON.
func (c ColumnsJSON) Value() (driver.Value, error) {
	if c == nil {
		return json.Marshal([]map[string]interface{}{})
	}
	return json.Marshal(c)
}

// ArticleMetadata adalah metadata khusus untuk rilis tipe artikel.
type ArticleMetadata struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ReleaseID      uuid.UUID `gorm:"type:uuid;not null;index" json:"release_id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	Excerpt        string    `gorm:"type:text" json:"excerpt,omitempty"`
	AuthorName     string    `gorm:"type:varchar(255)" json:"author_name,omitempty"`
	ReadingTimeMin *int      `json:"reading_time_minutes,omitempty"`
	IsFeatured     bool      `gorm:"not null;default:false;index" json:"is_featured"`
	Category       string    `gorm:"type:varchar(100);index" json:"category,omitempty"`
	CreatedAt      time.Time `gorm:"not null;default:NOW()" json:"created_at"`
	UpdatedAt      time.Time `gorm:"not null;default:NOW()" json:"updated_at"`
}

// TableName overrides the table name.
func (ArticleMetadata) TableName() string {
	return "article_metadata"
}
