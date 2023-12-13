package errors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	w := httptest.NewRecorder()
	err := NotFound
	WriteResponse(w, err)
	if w.Code != HTTPCode[err] && w.Body.Len() != len(err.Error()) {
		t.Error("written wrong status code or text")
	}

	err = fmt.Errorf("some random error %v", "pipi")
	w = httptest.NewRecorder()
	WriteResponse(w, err)
	if w.Code != http.StatusInternalServerError && w.Body.Len() != len(InternalServerError.Error()) {
		t.Error("written wrong status code or text")
	}
}
