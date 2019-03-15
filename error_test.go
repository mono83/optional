package optional

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Panic(t *testing.T) {
	assert.Panics(t, func() {
		OfError(errors.New("foo")).Panic()
	})

	// No panic
	Error{}.Panic()
}
