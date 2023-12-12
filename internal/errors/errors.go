package errors

import (
	"google.golang.org/grpc/status"
)

var (
	NotFound            = status.Error(4, "not found")
	BadURL              = status.Error(3, "doesn't look like url")
	TooLongIdentifier   = status.Error(3, "too long identifier")
	TooShortIdentifier  = status.Error(3, "too short identifier")
	InternalServerError = status.Error(13, "internal server error")
)
