package optional

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var orElseData = []struct {
	Expected interface{}
	Given    interface{}
}{
	{true, emptyBool.OrElse(true)},
	{false, emptyBool.OrElse(false)},
	{false, OfBool(false).OrElse(true)},
	{88, emptyInt.OrElse(88)},
	{21, OfInt(21).OrElse(12)},
	{"foo", emptyString.OrElse("foo")},
	{"foo", OfString("foo").OrElse("bar")},
	{0.12, emptyFloat64.OrElse(0.12)},
	{2.2, OfFloat64(2.2).OrElse(10)},
	{"boo", emptyMixed.OrElse("boo")},
	{"foo", OfMixed("foo").OrElse("bar")},
	{time.Unix(1, 2), emptyTime.OrElse(time.Unix(1, 2))},
	{time.Unix(3, 4), OfTime(time.Unix(3, 4)).OrElse(time.Now())},
	{time.Minute, emptyDuration.OrElse(time.Minute)},
	{time.Second, OfDuration(time.Second).OrElse(time.Minute)},
}

func TestOrElse(t *testing.T) {
	for _, data := range orElseData {
		t.Run(fmt.Sprint(data.Expected, data.Given), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Given)
		})
	}
}
