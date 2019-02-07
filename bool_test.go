package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolTrue(t *testing.T) {
	b := OfBool(true)

	assert.Equal(t, true, b.OrElse(false))
	assert.Equal(t, true, b.OrElse(true))
	assert.True(t, b.IsPresent())
	assert.True(t, b.Filter(func(bool) bool { return true }).IsPresent())
	assert.False(t, b.Filter(func(bool) bool { return false }).IsPresent())
	assert.False(t, b.Filter(nil).IsPresent())
	ref := false
	b.IfPresent(func(b bool) { ref = b })
	assert.Equal(t, true, ref)
	assert.Equal(t, OfInt(555), b.MapToInt(func(b bool) int {
		assert.True(t, b)
		return 555
	}))
	assert.Equal(t, OfString("foo"), b.MapToString(func(b bool) string {
		assert.True(t, b)
		return "foo"
	}))
	assert.Equal(t, OfFloat64(12.34), b.MapToFloat64(func(b bool) float64 {
		assert.True(t, b)
		return 12.34
	}))
}

func TestBoolFalse(t *testing.T) {
	b := OfBool(false)

	assert.Equal(t, false, b.OrElse(false))
	assert.Equal(t, false, b.OrElse(true))
	assert.True(t, b.IsPresent())
	assert.True(t, b.Filter(func(bool) bool { return true }).IsPresent())
	assert.False(t, b.Filter(func(bool) bool { return false }).IsPresent())
	assert.False(t, b.Filter(nil).IsPresent())
	ref := true
	b.IfPresent(func(b bool) { ref = b })
	assert.Equal(t, false, ref)
	assert.Equal(t, OfInt(444), b.MapToInt(func(b bool) int {
		assert.False(t, b)
		return 444
	}))
	assert.Equal(t, OfString("bar"), b.MapToString(func(b bool) string {
		assert.False(t, b)
		return "bar"
	}))
	assert.Equal(t, OfFloat64(5.6), b.MapToFloat64(func(b bool) float64 {
		assert.False(t, b)
		return 5.6
	}))
}

func TestBoolEmpty(t *testing.T) {
	b := Bool{}

	assert.Equal(t, false, b.OrElse(false))
	assert.Equal(t, true, b.OrElse(true))
	assert.False(t, b.IsPresent())
	assert.False(t, b.Filter(func(bool) bool { return true }).IsPresent())
	assert.False(t, b.Filter(func(bool) bool { return false }).IsPresent())
	assert.False(t, b.Filter(nil).IsPresent())
	ref := false
	b.IfPresent(func(bool) { ref = true })
	assert.Equal(t, false, ref)
	assert.Equal(t, Int{}, b.MapToInt(func(bool) int {
		return 33
	}))
	assert.Equal(t, String{}, b.MapToString(func(b bool) string {
		return "baz"
	}))
	assert.Equal(t, Float64{}, b.MapToFloat64(func(b bool) float64 {
		return 7.8
	}))
}
