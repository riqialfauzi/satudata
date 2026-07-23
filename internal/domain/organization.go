package domain

// Organization adalah model untuk organisasi/unit kerja.
type Organization struct {
	BaseModel
	Slug        string `gorm:"type:varchar(200);not null;uniqueIndex" json:"slug"`
	Name        string `gorm:"type:varchar(300);not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
	ImageURL    string `gorm:"type:text" json:"image_url,omitempty"`
	Website     string `gorm:"type:varchar(500)" json:"website,omitempty"`
}

// TableName overrides the table name.
func (Organization) TableName() string {
	return "organizations"
}
