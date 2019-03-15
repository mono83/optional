package optional

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var isPresentData = []struct {
	Expected bool
	Given    Optional
}{
	{false, Bool{}},
	{true, OfBool(true)},
	{true, OfBool(false)},
	{false, Int{}},
	{true, OfInt(76)},
	{true, OfInt(0)},
	{false, OfInt(0).FilterZero()},
	{true, OfInt(-22)},
	{false, String{}},
	{true, OfString("foo")},
	{true, OfString("")},
	{false, OfString("").FilterEmpty()},
	{false, Float64{}},
	{true, OfFloat64(0.1)},
	{true, OfFloat64(0)},
	{false, OfFloat64(0).FilterZero()},
	{false, OfFloat64(0.0).FilterZero()},
	{false, Time{}},
	{true, OfTime(time.Now())},
	{true, OfTime(time.Time{})},
	{true, OfTime(time.Time{})},
	{false, OfTime(time.Time{}).FilterZero()},
	{false, Duration{}},
	{true, OfDuration(time.Second)},
	{true, OfDuration(0)},
	{false, OfDuration(0).FilterZero()},
	{false, Mixed{}},
	{true, OfMixed(0)},
	{true, OfMixed(false)},
	{true, OfMixed("")},
	{true, OfMixed("bar")},
	{false, OfMixed(nil)},
}

func TestIsPresent(t *testing.T) {
	for _, data := range isPresentData {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given.IsPresent(), fmt.Sprintf("%T fails test", data.Given))
		})
	}
}

var isPresentRawData = []struct {
	Expected bool
	Given    interface{}
}{
	{false, nil},
	{true, false},
	{true, true},
	{true, ""},
	{true, "foo"},
}

func TestIsPresentFunc(t *testing.T) {
	// Testing optionals
	for _, data := range isPresentData {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, IsPresent(data.Given), fmt.Sprintf("%T fails test", data.Given))
		})
	}

	// Testing raw data
	for _, data := range isPresentRawData {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, IsPresent(data.Given), fmt.Sprintf("%T fails test", data.Given))
		})
	}

	// Empty pointer
	var s *string
	var i interface{}
	i = s
	assert.False(t, IsPresent(i), "Pointer empty reference check failed")

	// Filled pointer
	str := "foo"
	s = &str
	i = s
	assert.True(t, IsPresent(i), "Pointer reference check failed")
}
