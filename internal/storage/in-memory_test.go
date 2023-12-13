package storage

import (
	"context"
	goerrors "errors"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	"testing"
)

func TestInMemory_Get(t *testing.T) {
	inMemory := NewInMemory()
	original, err := inMemory.Get(context.Background(), "qwerty")
	if !goerrors.Is(err, errors.NotFound) {
		t.Errorf("wrong result; expected NotFound error, got %v", err)
	}
	inMemory.shortToOriginal.Store("qwerty", "google.com")
	original, err = inMemory.Get(context.Background(), "qwerty")
	if err != nil {
		t.Errorf("wrong result; expected `google.com`, got an error:%v", err)
	}
	if original != "google.com" {
		t.Errorf("wrong result; expected `google.com`, got %v", original)
	}
}

func TestInMemory_Set(t *testing.T) {
	inMemory := NewInMemory()
	if err := inMemory.Set(context.Background(), "qwerty", "google.com"); err != nil {
		t.Error(err)
	}
	val, _ := inMemory.shortToOriginal.Load("qwerty")
	original := val.(string)
	if original != "google.com" {
		t.Errorf("wrong result; expected `google.com`, got %v", original)
	}
}

func TestInMemory_GetShortLink(t *testing.T) {
	inMemory := NewInMemory()
	inMemory.originalToShort.Store("google.com", "qwerty")
	identifier, err := inMemory.GetShortLink(context.Background(), "google.com")
	if err != nil {
		t.Error(err)
	}
	if identifier != "qwerty" {
		t.Errorf("wrong result; expected `qwerty`, got %v", identifier)
	}
}
