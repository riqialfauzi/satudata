package database

import (
	"fmt"
	"log"
	"time"

	"github.com/satudata/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Init membuka koneksi ke PostgreSQL menggunakan GORM.
func Init(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	gormLogger := logger.Default.LogMode(logger.Info)
	if cfg.SSLMode == "disable" && cfg.Host == "localhost" {
		gormLogger = logger.Default.LogMode(logger.Warn)
	}

	dialector := postgres.Open(dsn)

	var err error
	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("[Database] Successfully connected to PostgreSQL")
	return db, nil
}

// GetDB mengembalikan instance database global.
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("[Database] Database not initialized. Call Init() first.")
	}
	return db
}

// Close menutup koneksi database.
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get sql.DB for closing: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}

// HealthCheck memeriksa koneksi database.
func HealthCheck() error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	return sqlDB.Ping()
}
