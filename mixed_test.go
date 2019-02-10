package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMixedPresentOrElse(t *testing.T) {
	// Is present
	assert.False(t, emptyMixed.IsPresent())
	assert.False(t, OfMixed(nil).IsPresent())
	ref := 5
	assert.True(t, OfMixed(&ref).IsPresent())
	assert.True(t, OfMixed(5).IsPresent())

	// If present
	OfMixed(1).IfPresent(func(i interface{}) {
		assert.Equal(t, 1, i)
	})
	OfMixed(nil).IfPresent(func(interface{}) {
		assert.Fail(t, "Should not be invoked")
	})
	OfMixed("foo").IfPresent(func(i interface{}) {
		assert.Equal(t, "foo", i)
	})

	// Or else
	assert.Equal(t, 9, OfMixed(9).OrElse(1))
	assert.Equal(t, "foo", OfMixed(nil).OrElse("foo"))
}

func TestMixedFilter(t *testing.T) {
	f := func(in interface{}) bool {
		f64, ok := in.(float64)
		return ok && f64 == 42.
	}

	assert.True(t, OfMixed(42.).Filter(f).IsPresent())
	assert.False(t, OfMixed(42).Filter(f).IsPresent())
	assert.False(t, OfMixed("42").Filter(f).IsPresent())
}
