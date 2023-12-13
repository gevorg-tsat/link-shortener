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

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (s *InMemory) Get(_ context.Context, identifier string) (originalURL string, err error) {
	value, ok := s.shortToOriginal.Load(identifier)
	if !ok {
		return "", errors.NotFound
	}
	originalURL = value.(string)
	return originalURL, nil
}

func (s *InMemory) Set(_ context.Context, identifier, originalURL string) error {
	value, ok := s.shortToOriginal.Load(identifier)
	if ok && value == originalURL {
		return nil
	}
	s.shortToOriginal.Store(identifier, originalURL)
	s.originalToShort.Store(originalURL, identifier)
	return nil
}

func (s *InMemory) GetShortLink(_ context.Context, originalURL string) (shortURL string, err error) {
	value, ok := s.originalToShort.Load(originalURL)
	if !ok {
		return "", errors.NotFound
	}
	shortURL = value.(string)
	return shortURL, nil
}
