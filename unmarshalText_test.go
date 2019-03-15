package optional

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var unmarshalTextData = []struct {
	Target interface {
		UnmarshalText(bts []byte) error
	}
	Expected interface{}
	Source   []byte
	Error    error
}{
	{&Bool{}, Bool{}, nil, nil},
	{&Bool{}, Bool{}, []byte(""), nil},
	{&Bool{}, OfBool(true), []byte("tRuE"), nil},
	{&Bool{}, OfBool(false), []byte("fALse"), nil},
	{&Bool{}, nil, []byte("foo"), errors.New("unable to produce boolean from foo")},

	{&Int{}, Int{}, nil, nil},
	{&Int{}, Int{}, []byte(""), nil},
	{&Int{}, OfInt(42), []byte("42"), nil},
	{&Int{}, OfInt(-2), []byte("-2"), nil},
	{&Int{}, nil, []byte("0xFF"), errors.New("strconv.ParseInt: parsing \"0xFF\": invalid syntax")},
	{&Int{}, OfInt(35), []byte("0035"), nil},

	{&Float64{}, Float64{}, nil, nil},
	{&Float64{}, Float64{}, []byte(""), nil},
	{&Float64{}, OfFloat64(10.), []byte("10"), nil},
	{&Float64{}, OfFloat64(-.21), []byte("-.21"), nil},
	{&Float64{}, nil, []byte("3,2"), errors.New("strconv.ParseFloat: parsing \"3,2\": invalid syntax")},
	{&Float64{}, nil, []byte("foo"), errors.New("strconv.ParseFloat: parsing \"foo\": invalid syntax")},

	{&String{}, String{}, nil, nil},
	{&String{}, String{}, []byte(""), nil},
	{&String{}, OfString("true"), []byte("true"), nil},
	{&String{}, OfString("foo"), []byte("foo"), nil},

	{&TimeUnixSeconds{}, TimeUnixSeconds{}, nil, nil},
	{&TimeUnixSeconds{}, TimeUnixSeconds{}, []byte(""), nil},
	{&TimeUnixSeconds{}, OfUnixSeconds(42), []byte("42"), nil},
	{&TimeUnixSeconds{}, OfUnixSeconds(-2), []byte("-2"), nil},
	{&TimeUnixSeconds{}, nil, []byte("0xFF"), errors.New("strconv.ParseInt: parsing \"0xFF\": invalid syntax")},
	{&TimeUnixSeconds{}, OfUnixSeconds(35), []byte("0035"), nil},

	{&DurationMillis{}, DurationMillis{}, nil, nil},
	{&DurationMillis{}, DurationMillis{}, []byte(""), nil},
	{&DurationMillis{}, OfMilliseconds(432), []byte("432"), nil},
	{&DurationMillis{}, OfMilliseconds(-800), []byte("-800"), nil},
	{&DurationMillis{}, nil, []byte("0xAA"), errors.New("strconv.ParseInt: parsing \"0xAA\": invalid syntax")},
	{&DurationMillis{}, OfMilliseconds(35), []byte("0035"), nil},

	{&DurationSeconds{}, DurationSeconds{}, nil, nil},
	{&DurationSeconds{}, DurationSeconds{}, []byte(""), nil},
	{&DurationSeconds{}, OfSeconds(31), []byte("31"), nil},
	{&DurationSeconds{}, OfSeconds(-999), []byte("-999"), nil},
	{&DurationSeconds{}, nil, []byte("0xAB"), errors.New("strconv.ParseInt: parsing \"0xAB\": invalid syntax")},
	{&DurationSeconds{}, OfSeconds(35), []byte("0035"), nil},

	{&DurationMinutes{}, DurationMinutes{}, nil, nil},
	{&DurationMinutes{}, DurationMinutes{}, []byte(""), nil},
	{&DurationMinutes{}, OfMinutes(221), []byte("221"), nil},
	{&DurationMinutes{}, OfMinutes(-4), []byte("-4"), nil},
	{&DurationMinutes{}, nil, []byte("0xCE"), errors.New("strconv.ParseInt: parsing \"0xCE\": invalid syntax")},
	{&DurationMinutes{}, OfMinutes(35), []byte("0035"), nil},
}

func Test_UnmarshalText(t *testing.T) {
	for _, data := range unmarshalTextData {
		t.Run(fmt.Sprintf("%v %v", data.Expected, data.Source), func(t *testing.T) {
			err := data.Target.UnmarshalText(data.Source)
			if data.Error != nil {
				if assert.NotNil(t, err, "Error expected") {
					assert.Equal(t, data.Error.Error(), err.Error())
				}
			} else {
				if assert.NoError(t, err) {
					sub := reflect.ValueOf(data.Target).Elem().Interface()
					assert.Equal(t, data.Expected, sub)
				}
			}
		})
	}
}
