package optional

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

var getData = []struct {
	ExpectedValue interface{}
	Given         interface{}
}{
	{true, OfBool(true)},
	{false, OfBool(false)},
	{nil, Bool{}},
	{15, OfInt(15)},
	{-2, OfInt(-2)},
	{0, OfInt(0)},
	{nil, Int{}},
	{"", OfString("")},
	{"bar", OfString("bar")},
	{nil, String{}},
	{.1, OfFloat64(.1)},
	{-222.01, OfFloat64(-222.01)},
	{nil, Float64{}},
	{time.Unix(12345, 678), OfTime(time.Unix(12345, 678))},
	{time.Time{}, OfTime(time.Time{})},
	{nil, Time{}},
	{time.Minute, OfDuration(time.Minute)},
	{time.Duration(0), OfDuration(0)},
	{nil, Duration{}},
	{"foo", OfMixed("foo")},
	{0, OfMixed(0)},
	{nil, OfMixed(nil)},
}

func TestGet(t *testing.T) {
	for _, data := range getData {
		t.Run(fmt.Sprint(data.Given), func(t *testing.T) {
			method := reflect.ValueOf(data.Given).MethodByName("Get")
			if data.ExpectedValue == nil {
				assert.Panics(t, func() { method.Call(nil) })
			} else {
				response := method.Call(nil)
				if assert.Len(t, response, 1) {
					assert.Equal(t, data.ExpectedValue, response[0].Interface())
				}
			}
		})
	}
}
