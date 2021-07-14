package wrap_test

import (
	"errors"
	"fmt"
	"testing"

	gitErrors "github.com/pkg/errors"
)

// gitErrors.New() also records the stack trace at the point it was called,
// that's why I use official errors.New().
var err = errors.New("a test error")

func TestWrap(t *testing.T) {
	wrappedErr := gitErrors.Wrap(err, "an example context error")
	fmt.Printf("%+v\n", wrappedErr)
}
