package storage

import (
	"context"
	"fmt"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type link struct {
	ShortURL    string `gorm:"type:char(10);primaryKey"`
	OriginalURL string `gorm:"unique;not_null"`
}

// Connect to db
func NewDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&link{}); err != nil {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
		return nil, err
	}
	return &DB{db: db}, nil
}

// Get original link from db by short link
func (s *DB) Get(ctx context.Context, identifier string) (originalURL string, err error) {
	val := link{ShortURL: identifier}
	if res := s.db.WithContext(ctx).First(&val); res.RowsAffected == 0 {
		return "", errors.NotFound
	}
	return val.OriginalURL, nil
}

// Set in db short link and original link
func (s *DB) Set(ctx context.Context, identifier, originalURL string) error {
	search := link{OriginalURL: originalURL, ShortURL: identifier}
	res := s.db.WithContext(ctx).Create(&search)
	return res.Error
}

// GetShortLink from db by original link
func (s *DB) GetShortLink(ctx context.Context, originalURL string) (shortURL string, err error) {
	var val link
	if res := s.db.WithContext(ctx).First(&val, &link{OriginalURL: originalURL}); res.RowsAffected == 0 {
		return "", errors.NotFound
	}
	return val.ShortURL, nil
}

// Shutdown connection with db
func (s *DB) Shutdown() {
	dbInstance, _ := s.db.DB()
	err := dbInstance.Close()
	if err != nil {
		panic(err)
	}
}
