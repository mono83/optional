package optional

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var dataTo = []struct {
	Expected interface{}
	Given    interface{}
}{
	{String{}, Int{}.ToString()},
	{OfString("72"), OfInt(72).ToString()},
	{DurationSeconds{}, Int{}.ToDurationSeconds()},
	{OfSeconds(60), OfInt(60).ToDurationSeconds()},
	{Int{}, Time{}.ToUnixSeconds()},
	{OfInt(10), OfTime(time.Unix(10, 20000000)).ToUnixSeconds()},
	{Int{}, Time{}.ToUnixMillis()},
	{OfInt(10020), OfTime(time.Unix(10, 20000000)).ToUnixMillis()},
	{Float64{}, Duration{}.ToSeconds()},
	{OfFloat64(60), OfDuration(time.Minute).ToSeconds()},
	{Int{}, Duration{}.ToMillis()},
	{OfInt(2000), OfDuration(time.Second * 2).ToMillis()},
	{String{}, Error{}.ToString()},
	{OfString("hello, World"), OfError(errors.New("hello, World")).ToString()},
}

func TestTo(t *testing.T) {
	for _, data := range dataTo {
		t.Run(fmt.Sprint(data.Expected, " ", data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given)
		})
	}
}
