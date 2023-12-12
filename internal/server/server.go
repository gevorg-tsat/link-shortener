package server

import (
	"context"
	inerrors "errors"
	"fmt"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	desc "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gevorg-tsat/link-shortener/internal/shorting"
	"github.com/gevorg-tsat/link-shortener/internal/storage"
	"log"
	"net/url"
	"strings"
)

type ShortenerServer struct {
	desc.UnimplementedShortenerV1Server
	sg  storage.Storage
	cfg config.Config
}

func New(sg storage.Storage, cfg config.Config) *ShortenerServer {
	return &ShortenerServer{sg: sg, cfg: cfg}
}

func (s *ShortenerServer) Get(ctx context.Context, link *desc.ShortLink) (*desc.OriginalLink, error) {
	shortLinkIdentifier := link.Url
	if slashInd := strings.Index(shortLinkIdentifier, "//"); slashInd != -1 {
		shortLinkIdentifier = shortLinkIdentifier[slashInd+2:]
	}
	if slashInd := strings.Index(shortLinkIdentifier, "/"); slashInd != -1 {
		shortLinkIdentifier = shortLinkIdentifier[slashInd+1:]
	}
	shortLinkIdentifier = strings.TrimRight(shortLinkIdentifier, "/")
	if len(shortLinkIdentifier) > 10 {
		return nil, errors.TooLongIdentifier
	}
	if len(shortLinkIdentifier) < 10 {
		return nil, errors.TooShortIdentifier
	}
	originalURL, err := s.sg.Get(ctx, shortLinkIdentifier)
	if err != nil {
		if inerrors.Is(err, errors.NotFound) {
			return nil, err
		}
		log.Println("storage.Get error:", err)
		return nil, errors.InternalServerError
	}
	return &desc.OriginalLink{Url: originalURL}, nil
}

func (s *ShortenerServer) Post(ctx context.Context, link *desc.OriginalLink) (*desc.ShortLink, error) {
	_, err := url.ParseRequestURI(link.Url)
	if err != nil {
		return nil, errors.BadURL
	}
	identifier := shorting.GenerateIdentifier()
	_, err = s.sg.Get(ctx, identifier)
	for !inerrors.Is(err, errors.NotFound) {
		identifier = shorting.GenerateIdentifier()
		_, err = s.sg.Get(ctx, identifier)
	}
	if err != nil && !inerrors.Is(err, errors.NotFound) {
		log.Println("storage.Get error:", err)
		return nil, errors.InternalServerError
	}
	if err = s.sg.Set(ctx, identifier, link.Url); err != nil {
		log.Println("storage.Set error:", err)
		return nil, errors.InternalServerError
	}
	return s.shortLinkBuilder(identifier), nil
}

func (s *ShortenerServer) shortLinkBuilder(identifier string) *desc.ShortLink {
	originalUrl := fmt.Sprintf("http://%v:%v/%v", s.cfg.Host, s.cfg.Port, identifier)
	return &desc.ShortLink{Url: originalUrl}
}
