package main

import (
	"log"
	"time"

	"github.com/satudata/backend/internal/config"
	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/pkg/database"
	"github.com/satudata/backend/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	if err := logger.Init(cfg.App.LogLevel, cfg.App.Environment); err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	// Initialize database
	db, err := database.Init(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Starting seed...")

	// Seed users
	seedUsers(db)

	// Seed standards
	seedStandards(db)

	// Seed releases
	seedReleases(db)

	log.Println("Seed completed successfully!")
}

func seedUsers(db *gorm.DB) {
	// Admin user
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := domain.User{
		Email:        "admin@satudata.go.id",
		PasswordHash: string(adminHash),
		FullName:     "Admin Satudata",
		Role:         domain.UserRoleAdmin,
		IsActive:     true,
	}

	if err := db.Where("email = ?", admin.Email).FirstOrCreate(&admin).Error; err != nil {
		log.Printf("Failed to seed admin user: %v", err)
	} else {
		log.Printf("Admin user created: %s", admin.Email)
	}

	// Editor user
	editorHash, _ := bcrypt.GenerateFromPassword([]byte("editor123"), bcrypt.DefaultCost)
	editor := domain.User{
		Email:        "editor@satudata.go.id",
		PasswordHash: string(editorHash),
		FullName:     "Editor Satudata",
		Role:         domain.UserRoleEditor,
		IsActive:     true,
	}

	if err := db.Where("email = ?", editor.Email).FirstOrCreate(&editor).Error; err != nil {
		log.Printf("Failed to seed editor user: %v", err)
	} else {
		log.Printf("Editor user created: %s", editor.Email)
	}
}

func seedStandards(db *gorm.DB) {
	standards := []domain.Standard{
		{
			Title:       "Standar Data Statistik 2025",
			Description: "Standar data statistik nasional untuk tahun 2025",
			Year:        2025,
			Status:      domain.StandardStatusActive,
			Version:     "1.0",
			IsCurrent:   true,
		},
		{
			Title:       "Standar Data Statistik 2024",
			Description: "Standar data statistik nasional untuk tahun 2024",
			Year:        2024,
			Status:      domain.StandardStatusActive,
			Version:     "1.0",
			IsCurrent:   false,
		},
		{
			Title:       "Standar Data Geospasial 2025",
			Description: "Standar data geospasial untuk tahun 2025",
			Year:        2025,
			Status:      domain.StandardStatusDraft,
			Version:     "0.9",
			IsCurrent:   false,
		},
	}

	for _, s := range standards {
		if err := db.Where("year = ? AND version = ?", s.Year, s.Version).FirstOrCreate(&s).Error; err != nil {
			log.Printf("Failed to seed standard %d: %v", s.Year, err)
		} else {
			log.Printf("Standard created: %s (%d)", s.Title, s.Year)
		}
	}
}

func seedReleases(db *gorm.DB) {
	now := time.Now()
	published := now.Add(-24 * time.Hour)

	releases := []struct {
		release domain.Release
		dataset *domain.DatasetMetadata
		article *domain.ArticleMetadata
	}{
		{
			release: domain.Release{
				Title:       "Indeks Pembangunan Manusia 2024",
				Slug:        "indeks-pembangunan-manusia-2024",
				Description: "Data IPM tahun 2024 untuk seluruh provinsi di Indonesia",
				ReleaseType: domain.ReleaseTypeDataset,
				Status:      domain.ReleaseStatusPublished,
				Year:        2024,
				Tags:        []string{"ipm", "pembangunan", "manusia"},
				ViewCount:   1250,
				PublishedAt: &published,
			},
			dataset: &domain.DatasetMetadata{
				FileURL:         "https://storage.satudata.go.id/datasets/ipm-2024.csv",
				FileFormat:      "csv",
				FileSize:        245760,
				RowCount:        intPtr(34),
				ColumnCount:     intPtr(12),
				DataSource:      "Badan Pusat Statistik",
				UpdateFrequency: "tahunan",
				DataPeriodStart: datePtr(2024, 1, 1),
				DataPeriodEnd:   datePtr(2024, 12, 31),
			},
		},
		{
			release: domain.Release{
				Title:       "Inflasi Bulanan Januari 2025",
				Slug:        "inflasi-bulanan-januari-2025",
				Description: "Data inflasi bulan Januari 2025 berdasarkan kelompok pengeluaran",
				ReleaseType: domain.ReleaseTypeDataset,
				Status:      domain.ReleaseStatusPublished,
				Year:        2025,
				Tags:        []string{"inflasi", "ekonomi", "bulanan"},
				ViewCount:   890,
				PublishedAt: &published,
			},
			dataset: &domain.DatasetMetadata{
				FileURL:         "https://storage.satudata.go.id/datasets/inflasi-jan-2025.csv",
				FileFormat:      "csv",
				FileSize:        102400,
				RowCount:        intPtr(11),
				ColumnCount:     intPtr(8),
				DataSource:      "Badan Pusat Statistik",
				UpdateFrequency: "bulanan",
				DataPeriodStart: datePtr(2025, 1, 1),
				DataPeriodEnd:   datePtr(2025, 1, 31),
			},
		},
		{
			release: domain.Release{
				Title:       "Analisis Kemiskinan 2024",
				Slug:        "analisis-kemiskinan-2024",
				Description: "Analisis mendalam tentang tren kemiskinan di Indonesia tahun 2024",
				ReleaseType: domain.ReleaseTypeArticle,
				Status:      domain.ReleaseStatusPublished,
				Year:        2024,
				Tags:        []string{"kemiskinan", "analisis", "sosial"},
				ViewCount:   2100,
				PublishedAt: &published,
			},
			article: &domain.ArticleMetadata{
				Content:    "<h1>Analisis Kemiskinan 2024</h1><p>Tingkat kemiskinan di Indonesia terus menunjukkan tren penurunan yang signifikan...</p>",
				Excerpt:    "Tingkat kemiskinan di Indonesia terus menunjukkan tren penurunan yang signifikan pada tahun 2024...",
				AuthorName: "Tim Analis Satudata",
				Category:   "Sosial",
				IsFeatured: true,
			},
		},
		{
			release: domain.Release{
				Title:       "Infografis: Pertumbuhan Ekonomi 2024",
				Slug:        "infografis-pertumbuhan-ekonomi-2024",
				Description: "Infografis pertumbuhan ekonomi Indonesia triwulan IV 2024",
				ReleaseType: domain.ReleaseTypeInfographic,
				Status:      domain.ReleaseStatusDraft,
				Year:        2024,
				Tags:        []string{"ekonomi", "infografis", "pertumbuhan"},
				ViewCount:   0,
			},
		},
	}

	for _, item := range releases {
		var existing domain.Release
		result := db.Where("slug = ?", item.release.Slug).First(&existing)

		if result.Error != nil {
			// Create with metadata
			if item.dataset != nil {
				item.release.DatasetMetadata = item.dataset
			}
			if item.article != nil {
				item.release.ArticleMetadata = item.article
			}

			if err := db.Create(&item.release).Error; err != nil {
				log.Printf("Failed to seed release '%s': %v", item.release.Title, err)
			} else {
				log.Printf("Release created: %s", item.release.Title)
			}
		} else {
			log.Printf("Release already exists: %s", existing.Title)
		}
	}
}

func intPtr(i int) *int {
	return &i
}

func datePtr(year, month, day int) *time.Time {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &t
}
