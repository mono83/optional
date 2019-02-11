package optional

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type uniOptional interface {
	IsPresent() bool
	Scan(interface{}) error
	String() string
}

var scanData = []struct {
	Source  interface{}
	Target  uniOptional
	Present bool
	String  string
}{
	{nil, &Bool{}, false, "Optional.Empty"},
	{nil, &Int{}, false, "Optional.Empty"},
	{nil, &String{}, false, "Optional.Empty"},
	{nil, &Float64{}, false, "Optional.Empty"},
	{nil, &TimeUnixSeconds{}, false, "Optional.Empty"},
	{nil, &DurationSeconds{}, false, "Optional.Empty"},

	{[]uint8("true"), &Bool{}, true, "Optional[true]"},
	{[]uint8("1"), &Bool{}, true, "Optional[true]"},
	{[]uint8("false"), &Bool{}, true, "Optional[false]"},
	{[]uint8("0"), &Bool{}, true, "Optional[false]"},
	{[]uint8(""), &Bool{}, true, "Optional[false]"},
	{true, &Bool{}, true, "Optional[true]"},
	{false, &Bool{}, true, "Optional[false]"},

	{[]uint8("21"), &Int{}, true, "Optional[21]"},
	{[]uint8("-441"), &Int{}, true, "Optional[-441]"},
	{54, &Int{}, true, "Optional[54]"},
	{int64(-12), &Int{}, true, "Optional[-12]"},

	{[]uint8("foo"), &String{}, true, "Optional[foo]"},
	{"bar", &String{}, true, "Optional[bar]"},

	{[]uint8("12"), &Float64{}, true, "Optional[12.000000]"},
	{[]uint8("0.11"), &Float64{}, true, "Optional[0.110000]"},
	{[]uint8("-3.1415"), &Float64{}, true, "Optional[-3.141500]"},
	{65.112, &Float64{}, true, "Optional[65.112000]"},
	{3, &Float64{}, true, "Optional[3.000000]"},

	{[]uint8("12345"), &TimeUnixSeconds{}, true, "Optional[1970-01-01 03:25:45 +0000 UTC]"},
	{-432112312, &TimeUnixSeconds{}, true, "Optional[1956-04-22 16:48:08 +0000 UTC]"},
	{int64(22), &TimeUnixSeconds{}, true, "Optional[1970-01-01 00:00:22 +0000 UTC]"},
	{time.Unix(123456789, 2231).UTC(), &TimeUnixSeconds{}, true, "Optional[1973-11-29 21:33:09.000002231 +0000 UTC]"},

	{[]uint8("99"), &DurationSeconds{}, true, "Optional[1m39s]"},
	{-15, &DurationSeconds{}, true, "Optional[-15s]"},
	{int64(2), &DurationSeconds{}, true, "Optional[2s]"},
	{time.Second * 8, &DurationSeconds{}, true, "Optional[8s]"},

	{[]uint8("2500"), &DurationMillis{}, true, "Optional[2.5s]"},
	{-1000, &DurationMillis{}, true, "Optional[-1s]"},
	{int64(20000), &DurationMillis{}, true, "Optional[20s]"},
	{time.Second * 8, &DurationMillis{}, true, "Optional[8s]"},

	{[]uint8("8"), &DurationMinutes{}, true, "Optional[8m0s]"},
	{-10, &DurationMinutes{}, true, "Optional[-10m0s]"},
	{int64(5), &DurationMinutes{}, true, "Optional[5m0s]"},
	{time.Second * 8, &DurationMinutes{}, true, "Optional[8s]"},
}

func TestScan(t *testing.T) {
	for _, data := range scanData {
		t.Run(data.String, func(t *testing.T) {
			err := data.Target.Scan(data.Source)
			if assert.NoError(t, err) {
				assert.Equal(t, data.Present, data.Target.IsPresent())
				assert.Equal(t, data.String, data.Target.String())
			}
		})
	}
}
