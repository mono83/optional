package optional

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBool_IfPresent(t *testing.T) {
	Bool{}.IfPresent(func(bool) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfBool(true).IfPresent(func(b bool) {
		assert.True(t, b)
	})
	assert.Equal(t, Bool{}, Bool{}.IfPresent(func(bool) {}))
	assert.Equal(t, OfBool(true), OfBool(true).IfPresent(func(bool) {}))
	assert.Equal(t, OfBool(true), OfBool(true).IfPresent(nil))
}

func TestInt_IfPresent(t *testing.T) {
	Int{}.IfPresent(func(int) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfInt(33).IfPresent(func(i int) {
		assert.Equal(t, 33, i)
	})
	assert.Equal(t, Int{}, Int{}.IfPresent(func(int) {}))
	assert.Equal(t, OfInt(55), OfInt(55).IfPresent(func(int) {}))
	assert.Equal(t, OfInt(1), OfInt(1).IfPresent(nil))
}

func TestString_IfPresent(t *testing.T) {
	String{}.IfPresent(func(string) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfString("foo").IfPresent(func(s string) {
		assert.Equal(t, "foo", s)
	})
	assert.Equal(t, String{}, String{}.IfPresent(func(string) {}))
	assert.Equal(t, OfString("bar"), OfString("bar").IfPresent(func(string) {}))
	assert.Equal(t, OfString(""), OfString("").IfPresent(nil))
}

func TestFloat64_IfPresent(t *testing.T) {
	Float64{}.IfPresent(func(float64) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfFloat64(0.123).IfPresent(func(f float64) {
		assert.Equal(t, 0.123, f)
	})
	assert.Equal(t, Float64{}, Float64{}.IfPresent(func(float64) {}))
	assert.Equal(t, OfFloat64(55), OfFloat64(55).IfPresent(func(float64) {}))
	assert.Equal(t, OfFloat64(1), OfFloat64(1).IfPresent(nil))
}

func TestTime_IfPresent(t *testing.T) {
	Time{}.IfPresent(func(time.Time) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfTime(time.Unix(12, 34)).IfPresent(func(tt time.Time) {
		assert.Equal(t, time.Unix(12, 34), tt)
	})
	assert.Equal(t, Time{}, Time{}.IfPresent(func(time.Time) {}))
	assert.Equal(t, OfTime(time.Unix(12, 34)), OfTime(time.Unix(12, 34)).IfPresent(func(time.Time) {}))
	assert.Equal(t, OfTime(time.Unix(56, 78)), OfTime(time.Unix(56, 78)).IfPresent(nil))
}

func TestDuration_IfPresent(t *testing.T) {
	Duration{}.IfPresent(func(time.Duration) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfDuration(time.Hour).IfPresent(func(d time.Duration) {
		assert.Equal(t, time.Hour, d)
	})
	assert.Equal(t, Duration{}, Duration{}.IfPresent(func(time.Duration) {}))
	assert.Equal(t, OfDuration(time.Minute), OfDuration(time.Minute).IfPresent(func(time.Duration) {}))
	assert.Equal(t, OfDuration(time.Second), OfDuration(time.Second).IfPresent(nil))
}

func TestMixed_IfPresent(t *testing.T) {
	Mixed{}.IfPresent(func(interface{}) {
		assert.Fail(t, "This method should not be invoked")
	})
	OfMixed("foo").IfPresent(func(i interface{}) {
		assert.Equal(t, "foo", i)
	})
	assert.Equal(t, Mixed{}, Mixed{}.IfPresent(func(interface{}) {}))
	assert.Equal(t, OfMixed("bar"), OfMixed("bar").IfPresent(func(interface{}) {}))
	assert.Equal(t, OfMixed(7), OfMixed(7).IfPresent(nil))
}

func TestError_IfPresent(t *testing.T) {
	Error{}.IfPresent(func(error) {
		assert.Fail(t, "This method should not be invoked")
	})

	OfError(errors.New("bar")).IfPresent(func(e error) {
		assert.Equal(t, errors.New("bar"), e)
	})

	assert.Equal(t, Error{}, Error{}.IfPresent(func(error) {}))
	assert.Equal(t, OfError(errors.New("baz")), OfError(errors.New("baz")).IfPresent(func(error) {}))
	assert.Equal(t, OfError(errors.New("baz")), OfError(errors.New("baz")).IfPresent(nil))
}
