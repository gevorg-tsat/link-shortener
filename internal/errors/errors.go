package errors

import (
	"github.com/gevorg-tsat/link-shortener/internal/shorting"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

// Error statuses for gRPC with message
var (
	NotFound            = status.Error(5, "not found")
	BadURL              = status.Error(3, "doesn't look like url")
	TooLongIdentifier   = status.Error(3, "too long identifier")
	TooShortIdentifier  = status.Error(3, "too short identifier")
	InternalServerError = status.Error(13, "internal server error")
	WrongURL            = status.Error(3, "url is not link-shortener http-server url")
	InvalidIdentifier   = status.Errorf(3, "invalid charset for identifier. identifier only contains: %v", shorting.CharSet)
	HTTPCode            = map[error]int{
		NotFound:            http.StatusNotFound,
		BadURL:              http.StatusBadRequest,
		TooShortIdentifier:  http.StatusBadRequest,
		TooShortIdentifier:  http.StatusBadRequest,
		InternalServerError: http.StatusInternalServerError,
		InvalidIdentifier:   http.StatusBadRequest,
		WrongURL:            http.StatusBadRequest,
	}
)

// Write error response in body, or log it if error is unknown
func WriteResponse(w http.ResponseWriter, err error) {
	statusCode, ok := HTTPCode[err]
	if !ok {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Error(w, err.Error(), statusCode)
}
