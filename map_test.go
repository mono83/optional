package optional

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var mapDataEmpty = []struct {
	Expected interface{}
	Given    interface{}
}{
	{emptyString, emptyBool.MapToString(func(bool) string { return "foo" })},
	{emptyString, emptyInt.MapToString(func(int) string { return "foo" })},
	{emptyString, emptyFloat64.MapToString(func(float64) string { return "foo" })},
	{emptyString, emptyMixed.MapToString(func(interface{}) string { return "foo" })},
	{emptyString, emptyString.Map(func(string) string { return "foo" })},

	{emptyInt, emptyBool.MapToInt(func(bool) int { return 5 })},
	{emptyInt, emptyString.MapToInt(func(string) int { return 5 })},
	{emptyInt, emptyFloat64.MapToInt(func(float64) int { return 5 })},
	{emptyInt, emptyMixed.MapToInt(func(interface{}) int { return 5 })},
	{emptyInt, emptyInt.Map(func(int) int { return 5 })},

	{emptyBool, emptyInt.MapToBool(func(int) bool { return true })},
	{emptyBool, emptyString.MapToBool(func(string) bool { return true })},
	{emptyBool, emptyFloat64.MapToBool(func(float64) bool { return true })},
	{emptyBool, emptyMixed.MapToBool(func(interface{}) bool { return true })},
	{emptyBool, emptyBool.Map(func(bool) bool { return true })},

	{emptyFloat64, emptyInt.MapToFloat64(func(int) float64 { return .1 })},
	{emptyFloat64, emptyString.MapToFloat64(func(string) float64 { return .1 })},
	{emptyFloat64, emptyBool.MapToFloat64(func(bool) float64 { return .1 })},
	{emptyFloat64, emptyMixed.MapToFloat64(func(interface{}) float64 { return .1 })},
	{emptyFloat64, emptyFloat64.Map(func(float64) float64 { return .1 })},

	{emptyMixed, emptyMixed.Map(func(interface{}) interface{} { return "foo" })},

	{emptyTime, emptyTime.Map(func(time.Time) time.Time { return time.Now() })},

	{emptyDuration, emptyDuration.Map(func(time.Duration) time.Duration { return time.Hour })},
}

func TestMapEmpty(t *testing.T) {
	for _, data := range mapDataEmpty {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given)
		})
	}
}

var mapData = []struct {
	Expected interface{}
	Given    interface{}
}{
	{OfBool(false), OfBool(true).Map(func(b bool) bool {
		if b {
			return false
		}
		return true
	})},
	{OfInt(10), OfBool(true).MapToInt(func(b bool) int {
		if b {
			return 10
		}
		return 5
	})},
	{OfString("bar"), OfBool(true).MapToString(func(b bool) string {
		if b {
			return "bar"
		}
		return "foo"
	})},
	{OfFloat64(3), OfBool(true).MapToFloat64(func(b bool) float64 {
		if b {
			return 3
		}
		return 2
	})},

	{OfInt(20), OfInt(10).Map(func(i int) int {
		if i == 10 {
			return 20
		}
		return 30
	})},
	{OfBool(true), OfInt(10).MapToBool(func(i int) bool {
		return i == 10
	})},
	{OfString("bar"), OfInt(10).MapToString(func(i int) string {
		if i == 10 {
			return "bar"
		}
		return "foo"
	})},
	{OfFloat64(3), OfInt(10).MapToFloat64(func(i int) float64 {
		if i == 10 {
			return 3
		}
		return 2
	})},

	{OfString("x"), OfString("one").Map(func(s string) string {
		if s == "one" {
			return "x"
		}
		return "y"
	})},
	{OfBool(true), OfString("one").MapToBool(func(s string) bool {
		return s == "one"
	})},
	{OfInt(55), OfString("one").MapToInt(func(s string) int {
		if s == "one" {
			return 55
		}
		return 1
	})},
	{OfFloat64(3), OfString("one").MapToFloat64(func(s string) float64 {
		if s == "one" {
			return 3
		}
		return 2
	})},

	{OfFloat64(3.2), OfFloat64(.1234).Map(func(f float64) float64 {
		if f == .1234 {
			return 3.2
		}
		return 2
	})},
	{OfBool(true), OfFloat64(.1234).MapToBool(func(f float64) bool {
		return f == .1234
	})},
	{OfInt(55), OfFloat64(.1234).MapToInt(func(f float64) int {
		if f == .1234 {
			return 55
		}
		return 1
	})},
	{OfString("foo"), OfFloat64(.1234).MapToString(func(f float64) string {
		if f == .1234 {
			return "foo"
		}
		return "bar"
	})},

	{OfMixed("x"), OfMixed("one").Map(func(i interface{}) interface{} {
		if i == "one" {
			return "x"
		}
		return "y"
	})},
	{OfBool(true), OfMixed("one").MapToBool(func(i interface{}) bool {
		return i == "one"
	})},
	{OfInt(55), OfMixed("one").MapToInt(func(i interface{}) int {
		if i == "one" {
			return 55
		}
		return 1
	})},
	{OfFloat64(3), OfMixed("one").MapToFloat64(func(i interface{}) float64 {
		if i == "one" {
			return 3
		}
		return 2
	})},
	{OfString("x"), OfMixed("one").MapToString(func(i interface{}) string {
		if i == "one" {
			return "x"
		}
		return "y"
	})},

	{OfTime(time.Unix(1, 2)), OfTime(time.Unix(3, 4)).Map(func(t time.Time) time.Time {
		if t == time.Unix(3, 4) {
			return time.Unix(1, 2)
		}

		return t
	})},

	{OfDuration(time.Hour), OfDuration(time.Second).Map(func(d time.Duration) time.Duration {
		if d == time.Second {
			return time.Hour
		}

		return d
	})},
}

func TestMap(t *testing.T) {
	for _, data := range mapData {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given)
		})
	}
}
