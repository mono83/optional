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

func TestError_Then(t *testing.T) {
	i := 0
	OfError(nil).Then(func() error {
		i++
		return nil
	}).Then(func() error {
		i++
		return errors.New("foo")
	}).Then(func() error {
		assert.Fail(t, "This func should not be invoked")
		return nil
	})

	assert.Equal(t, 2, i)
}
