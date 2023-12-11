package server

import (
	"context"
	desc "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	pb "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gevorg-tsat/link-shortener/internal/storage"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerV1Server
	sg storage.Storage
}

func (s *ShortenerServer) Get(ctx context.Context, link *desc.ShortLink) (*desc.OriginalLink, error) {
	panic("NOT IMPLEMENTED")
}

func (s *ShortenerServer) Post(ctx context.Context, link *desc.ShortLink) (*desc.ShortLink, error) {
	panic("NOT IMPLEMENTED")
}
