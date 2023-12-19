package storage

import "context"

type Storage interface {
	// Get original link from storage by short link
	Get(ctx context.Context, shortURL string) (originalURL string, err error)

	// GetShortLink from storage by original link
	GetShortLink(ctx context.Context, originalURL string) (shortURL string, err error)

	// Set in storage short link and original link
	Set(ctx context.Context, shortURL, originalURL string) error

	// Shutdown connection with storage
	Shutdown()
}
