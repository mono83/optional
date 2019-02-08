package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
