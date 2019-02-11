package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilterBool(t *testing.T) {
	OfBool(true).Filter(func(b bool) bool {
		assert.Equal(t, true, b)
		return true
	}).Filter(func(b bool) bool {
		assert.Equal(t, true, b)
		return false
	}).Filter(func(bool) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})

	OfBool(false).Filter(func(b bool) bool {
		assert.Equal(t, false, b)
		return true
	}).Filter(func(b bool) bool {
		assert.Equal(t, false, b)
		return false
	}).Filter(func(bool) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterInt(t *testing.T) {
	OfInt(4455).Filter(func(i int) bool {
		assert.Equal(t, 4455, i)
		return true
	}).Filter(func(i int) bool {
		assert.Equal(t, 4455, i)
		return false
	}).Filter(func(int) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterString(t *testing.T) {
	OfString("foo").Filter(func(s string) bool {
		assert.Equal(t, "foo", s)
		return true
	}).Filter(func(s string) bool {
		assert.Equal(t, "foo", s)
		return false
	}).Filter(func(string) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterFloat64(t *testing.T) {
	OfFloat64(.54321).Filter(func(f float64) bool {
		assert.Equal(t, .54321, f)
		return true
	}).Filter(func(f float64) bool {
		assert.Equal(t, .54321, f)
		return false
	}).Filter(func(float64) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterTime(t *testing.T) {
	OfTime(time.Unix(123, 456)).Filter(func(tt time.Time) bool {
		assert.Equal(t, time.Unix(123, 456), tt)
		return true
	}).Filter(func(tt time.Time) bool {
		assert.Equal(t, time.Unix(123, 456), tt)
		return false
	}).Filter(func(time.Time) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterDuration(t *testing.T) {
	OfDuration(time.Hour).Filter(func(d time.Duration) bool {
		assert.Equal(t, time.Hour, d)
		return true
	}).Filter(func(d time.Duration) bool {
		assert.Equal(t, time.Hour, d)
		return false
	}).Filter(func(time.Duration) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}

func TestFilterMixed(t *testing.T) {
	OfMixed(int16(5)).Filter(func(i interface{}) bool {
		assert.Equal(t, int16(5), i)
		return true
	}).Filter(func(i interface{}) bool {
		assert.Equal(t, int16(5), i)
		return false
	}).Filter(func(interface{}) bool {
		assert.Fail(t, "Should not enter this func")
		return false
	})
}
