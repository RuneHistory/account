package errs

import (
	"errors"
	"net/http"
	"testing"
)

func TestInternalServer(t *testing.T) {
	t.Parallel()
	err := errors.New("this is a failure")
	customErr := InternalServer(err.Error())
	if err.Error() != customErr.Error() {
		t.Errorf("Expecting %s, got %s", err.Error(), customErr.Error())
	}
	if customErr.Code() != http.StatusInternalServerError {
		t.Errorf("Expecting %d, got %d", http.StatusInternalServerError, customErr.Code())
	}
}
