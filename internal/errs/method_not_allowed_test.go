package errs

import (
	"errors"
	"net/http"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	t.Parallel()
	err := errors.New("this is a failure")
	customErr := MethodNotAllowed(err.Error())
	if err.Error() != customErr.Error() {
		t.Errorf("Expecting %s, got %s", err.Error(), customErr.Error())
	}
	if customErr.Code() != http.StatusMethodNotAllowed {
		t.Errorf("Expecting %d, got %d", http.StatusMethodNotAllowed, customErr.Code())
	}
}
