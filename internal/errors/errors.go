package errors

import (
	"errors"
)

var (
	NotFound           = errors.New("not found")
	BadURL             = errors.New("doesn't look like url")
	TooLongIdentifier  = errors.New("too long identifier")
	TooShortIdentifier = errors.New("too short identifier")
)
