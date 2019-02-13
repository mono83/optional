package optional

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

var dataTo = []struct {
	Expected interface{}
	Given    interface{}
}{
	{emptyString, emptyInt.ToString()},
	{OfString("72"), OfInt(72).ToString()},
	{DurationSeconds{}, emptyInt.ToDurationSeconds()},
	{OfSeconds(60), OfInt(60).ToDurationSeconds()},
	{emptyInt, emptyTime.ToUnixSeconds()},
	{OfInt(10), OfTime(time.Unix(10, 20000000)).ToUnixSeconds()},
	{emptyInt, emptyTime.ToUnixMillis()},
	{OfInt(10020), OfTime(time.Unix(10, 20000000)).ToUnixMillis()},
	{emptyFloat64, emptyDuration.ToSeconds()},
	{OfFloat64(60), OfDuration(time.Minute).ToSeconds()},
	{emptyInt, emptyDuration.ToMillis()},
	{OfInt(2000), OfDuration(time.Second * 2).ToMillis()},
}

func TestTo(t *testing.T) {
	for _, data := range dataTo {
		t.Run(fmt.Sprint(data.Expected, " ", data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given)
		})
	}
}
