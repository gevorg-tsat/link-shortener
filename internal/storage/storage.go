package storage

import "context"

type Storage interface {
	Get(ctx context.Context, shortURL string) (originalURL string, err error)
	Set(ctx context.Context, shortURL, originalURL string) error
}
