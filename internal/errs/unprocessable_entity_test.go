package errs

import (
	"errors"
	"net/http"
	"testing"
)

func TestUnprocessableEntity(t *testing.T) {
	t.Parallel()
	err := errors.New("this is a failure")
	customErr := UnprocessableEntity(err.Error())
	if err.Error() != customErr.Error() {
		t.Errorf("Expecting %s, got %s", err.Error(), customErr.Error())
	}
	if customErr.Code() != http.StatusUnprocessableEntity {
		t.Errorf("Expecting %d, got %d", http.StatusUnprocessableEntity, customErr.Code())
	}
}
