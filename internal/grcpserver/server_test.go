package grcpserver

import (
	"context"
	goerrors "errors"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	desc "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gevorg-tsat/link-shortener/internal/storage"
	"testing"
)

func TestShortenerServer_Get(t *testing.T) {
	cfg := &config.Config{HTTP: config.HTTP{Host: "0.0.0.0", Port: 8080}}
	s := New(storage.NewInMemory(), cfg)
	_, err := s.Get(context.Background(), s.buildShortLink("qwerty"))
	if !goerrors.Is(err, errors.TooShortIdentifier) {
		t.Errorf("wrong result; expected err TooShortIdentifier, got %v", err)
	}
	_, err = s.Get(context.Background(), s.buildShortLink("qwertyerwerwesdf"))
	if !goerrors.Is(err, errors.TooLongIdentifier) {
		t.Errorf("wrong result; expected err TooLongIdentifier, got %v", err)
	}
	_, err = s.Get(context.Background(), s.buildShortLink("google.com"))
	if !goerrors.Is(err, errors.InvalidIdentifier) {
		t.Errorf("wrong result; expected err InvalidIdentifier, got %v", err)
	}
	_, err = s.Get(context.Background(), s.buildShortLink("googlekcom"))
	if !goerrors.Is(err, errors.NotFound) {
		t.Errorf("wrong result; expected err NotFound, got %v", err)
	}
	_, err = s.Get(context.Background(), &desc.ShortLink{Url: "google.com"})
	if !goerrors.Is(err, errors.WrongURL) {
		t.Errorf("wrong result; expected err WrongUrl, got %v", err)
	}
}

func TestShortenerServer_Post(t *testing.T) {
	cfg := &config.Config{HTTP: config.HTTP{Host: "0.0.0.0", Port: 8080}}
	s := New(storage.NewInMemory(), cfg)
	shortLink, err := s.Post(context.Background(), &desc.OriginalLink{Url: "http://google.com"})
	if err != nil {
		t.Error(err)
	}
	original, err := s.Get(context.Background(), shortLink)
	if err != nil {
		t.Error(err)
	}
	if original.Url != "http://google.com" {
		t.Errorf("wrong result; expected `http://google.com`; got %v", original.Url)
	}
	_, err = s.Post(context.Background(), &desc.OriginalLink{Url: "google.com"})
	if !goerrors.Is(err, errors.BadURL) {
		t.Errorf("wrong result; expected err BadUrl, got %v", err)
	}
}
