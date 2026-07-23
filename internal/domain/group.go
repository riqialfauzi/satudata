package domain

// Group adalah model untuk kategori/grup dataset.
type Group struct {
	BaseModel
	Slug        string `gorm:"type:varchar(200);not null;uniqueIndex" json:"slug"`
	Name        string `gorm:"type:varchar(300);not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	ImageURL    string `gorm:"type:text" json:"image_url,omitempty"`
}

// TableName overrides the table name.
func (Group) TableName() string {
	return "groups"
}
