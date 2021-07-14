package golang_error_handling_test

import (
	"api-examples/golang_error_handling"
	"errors"
	"testing"
)

func TestLogError(t *testing.T) {
	err := errors.New("example error message")
	golang_error_handling.LogError(err)
}
