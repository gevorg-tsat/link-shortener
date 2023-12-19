package storage

import (
	"context"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	"sync"
)

type InMemory struct {
	shortToOriginal sync.Map
	originalToShort sync.Map
}

// Create Map storage
func NewInMemory() *InMemory {
	return &InMemory{}
}

// Get original link from map by short link
func (s *InMemory) Get(_ context.Context, identifier string) (originalURL string, err error) {
	value, ok := s.shortToOriginal.Load(identifier)
	if !ok {
		return "", errors.NotFound
	}
	originalURL = value.(string)
	return originalURL, nil
}

// Set in map short link and original link
func (s *InMemory) Set(_ context.Context, identifier, originalURL string) error {
	value, ok := s.shortToOriginal.Load(identifier)
	if ok && value == originalURL {
		return nil
	}
	s.shortToOriginal.Store(identifier, originalURL)
	s.originalToShort.Store(originalURL, identifier)
	return nil
}

// GetShortLink from map by original link
func (s *InMemory) GetShortLink(_ context.Context, originalURL string) (shortURL string, err error) {
	value, ok := s.originalToShort.Load(originalURL)
	if !ok {
		return "", errors.NotFound
	}
	shortURL = value.(string)
	return shortURL, nil
}

// Mock shutdowning for map
func (s *InMemory) Shutdown() {

}
