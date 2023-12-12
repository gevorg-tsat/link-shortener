package storage

import "context"

type Storage interface {
	Get(ctx context.Context, shortURL string) (originalURL string, err error)
	GetShortLink(_ context.Context, originalURL string) (shortURL string, err error)
	Set(ctx context.Context, shortURL, originalURL string) error
}
