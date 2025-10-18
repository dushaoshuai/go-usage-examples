package interface_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_nil_error(t *testing.T) {
	var err error
	assert.Panic(t, func() { err.Error() }, "invalid memory address or nil pointer dereference")
}
