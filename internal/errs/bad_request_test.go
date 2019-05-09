package errs

import (
	"errors"
	"net/http"
	"testing"
)

func TestBadRequest(t *testing.T) {
	t.Parallel()
	err := errors.New("this is a failure")
	customErr := BadRequest(err.Error())
	if err.Error() != customErr.Error() {
		t.Errorf("Expecting %s, got %s", err.Error(), customErr.Error())
	}
	if customErr.Code() != http.StatusBadRequest {
		t.Errorf("Expecting %d, got %d", http.StatusBadRequest, customErr.Code())
	}
}
