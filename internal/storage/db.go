package storage

import (
	"context"
	"fmt"
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

func NewDB(host, user, password, dbname string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&link{}); err != nil {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
		return nil, err
	}
	return db, nil
}

func (s *DB) Get(ctx context.Context, shortURL string) (originalURL string, err error) {
	panic("NOT IMPLEMENTED")
}

func (s *DB) Set(ctx context.Context, shortURL, originalURL string) error {
	panic("NOT IMPLEMENTED")
}
