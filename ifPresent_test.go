package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBool_IfPresent(t *testing.T) {
	emptyBool.IfPresent(func(bool) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfBool(true).IfPresent(func(b bool) {
		assert.True(t, b)
	})
	assert.Equal(t, emptyBool, emptyBool.IfPresent(func(bool) {}))
	assert.Equal(t, OfBool(true), OfBool(true).IfPresent(func(bool) {}))
	assert.Equal(t, OfBool(true), OfBool(true).IfPresent(nil))
}

func TestInt_IfPresent(t *testing.T) {
	emptyInt.IfPresent(func(int) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfInt(33).IfPresent(func(i int) {
		assert.Equal(t, 33, i)
	})
	assert.Equal(t, emptyInt, emptyInt.IfPresent(func(int) {}))
	assert.Equal(t, OfInt(55), OfInt(55).IfPresent(func(int) {}))
	assert.Equal(t, OfInt(1), OfInt(1).IfPresent(nil))
}

func TestString_IfPresent(t *testing.T) {
	emptyString.IfPresent(func(string) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfString("foo").IfPresent(func(s string) {
		assert.Equal(t, "foo", s)
	})
	assert.Equal(t, emptyString, emptyString.IfPresent(func(string) {}))
	assert.Equal(t, OfString("bar"), OfString("bar").IfPresent(func(string) {}))
	assert.Equal(t, OfString(""), OfString("").IfPresent(nil))
}

func TestFloat64_IfPresent(t *testing.T) {
	emptyFloat64.IfPresent(func(float64) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfFloat64(0.123).IfPresent(func(f float64) {
		assert.Equal(t, 0.123, f)
	})
	assert.Equal(t, emptyFloat64, emptyFloat64.IfPresent(func(float64) {}))
	assert.Equal(t, OfFloat64(55), OfFloat64(55).IfPresent(func(float64) {}))
	assert.Equal(t, OfFloat64(1), OfFloat64(1).IfPresent(nil))
}

func TestTime_IfPresent(t *testing.T) {
	emptyTime.IfPresent(func(time.Time) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfTime(time.Unix(12, 34)).IfPresent(func(tt time.Time) {
		assert.Equal(t, time.Unix(12, 34), tt)
	})
	assert.Equal(t, emptyTime, emptyTime.IfPresent(func(time.Time) {}))
	assert.Equal(t, OfTime(time.Unix(12, 34)), OfTime(time.Unix(12, 34)).IfPresent(func(time.Time) {}))
	assert.Equal(t, OfTime(time.Unix(56, 78)), OfTime(time.Unix(56, 78)).IfPresent(nil))
}

func TestDuration_IfPresent(t *testing.T) {
	emptyDuration.IfPresent(func(time.Duration) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfDuration(time.Hour).IfPresent(func(d time.Duration) {
		assert.Equal(t, time.Hour, d)
	})
	assert.Equal(t, emptyDuration, emptyDuration.IfPresent(func(time.Duration) {}))
	assert.Equal(t, OfDuration(time.Minute), OfDuration(time.Minute).IfPresent(func(time.Duration) {}))
	assert.Equal(t, OfDuration(time.Second), OfDuration(time.Second).IfPresent(nil))
}

func TestMixed_IfPresent(t *testing.T) {
	emptyMixed.IfPresent(func(interface{}) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfMixed("foo").IfPresent(func(i interface{}) {
		assert.Equal(t, "foo", i)
	})
	assert.Equal(t, emptyMixed, emptyMixed.IfPresent(func(interface{}) {}))
	assert.Equal(t, OfMixed("bar"), OfMixed("bar").IfPresent(func(interface{}) {}))
	assert.Equal(t, OfMixed(7), OfMixed(7).IfPresent(nil))
}
