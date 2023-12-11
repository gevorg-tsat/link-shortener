package storage

import (
	"context"
	"sync"
)

type InMemory struct {
	m sync.Map
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (m *InMemory) Get(ctx context.Context, shortURL string) (originalURL string, err error) {
	panic("NOT IMPLEMENTED")
}

func (m *InMemory) Set(ctx context.Context, shortURL, originalURL string) error {
	panic("NOT IMPLEMENTED")
}
