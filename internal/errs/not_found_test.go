package errs

import (
	"errors"
	"net/http"
	"testing"
)

func TestNotFound(t *testing.T) {
	t.Parallel()
	err := errors.New("this is a failure")
	customErr := NotFound(err.Error())
	if err.Error() != customErr.Error() {
		t.Errorf("Expecting %s, got %s", err.Error(), customErr.Error())
	}
	if customErr.Code() != http.StatusNotFound {
		t.Errorf("Expecting %d, got %d", http.StatusNotFound, customErr.Code())
	}
}
