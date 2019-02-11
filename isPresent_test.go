package optional

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var isPresentData = []struct {
	Expected bool
	Given    interface {
		IsPresent() bool
	}
}{
	{false, emptyBool},
	{true, OfBool(true)},
	{true, OfBool(false)},
	{false, emptyInt},
	{true, OfInt(76)},
	{true, OfInt(0)},
	{false, OfInt(0).FilterZero()},
	{true, OfInt(-22)},
	{false, emptyString},
	{true, OfString("foo")},
	{true, OfString("")},
	{false, OfString("").FilterEmpty()},
	{false, emptyFloat64},
	{true, OfFloat64(0.1)},
	{true, OfFloat64(0)},
	{false, OfFloat64(0).FilterZero()},
	{false, OfFloat64(0.0).FilterZero()},
	{false, emptyTime},
	{true, OfTime(time.Now())},
	{true, OfTime(time.Time{})},
	{true, OfTime(time.Time{})},
	{false, OfTime(time.Time{}).FilterZero()},
	{false, emptyDuration},
	{true, OfDuration(time.Second)},
	{true, OfDuration(0)},
	{false, OfDuration(0).FilterZero()},
	{false, emptyMixed},
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
