package optional

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var textSourceYAML = `b1: true
b2: false
b3: null
s1: some
s3: null
i1: 58
i2: -9001
i3: 0
i4: null
f1: 0.2
f2: -376.11
f3: 0
f4: null
ts1: 1384923976
ts2: null
d1: 42
d2: null
d3: 9000
d4: null
d5: 3
d6: null
`

func TestYAMLMarshalling(t *testing.T) {
	var target struct {
		B1 Bool `yaml:"b1"`
		B2 Bool `yaml:"b2"`
		B3 Bool `yaml:"b3"`

		S1 String `yaml:"s1"`
		S3 String `yaml:"s3"`

		I1 Int `yaml:"i1"`
		I2 Int `yaml:"i2"`
		I3 Int `yaml:"i3"`
		I4 Int `yaml:"i4"`

		F1 Float64 `yaml:"f1"`
		F2 Float64 `yaml:"f2"`
		F3 Float64 `yaml:"f3"`
		F4 Float64 `yaml:"f4"`

		TS1 TimeUnixSeconds `yaml:"ts1"`
		TS2 TimeUnixSeconds `yaml:"ts2"`

		D1 DurationSeconds `yaml:"d1"`
		D2 DurationSeconds `yaml:"d2"`
		D3 DurationMillis  `yaml:"d3"`
		D4 DurationMillis  `yaml:"d4"`
		D5 DurationMinutes `yaml:"d5"`
		D6 DurationMinutes `yaml:"d6"`
	}

	if err := yaml.Unmarshal([]byte(textSourceYAML), &target); assert.NoError(t, err) {
		assert.True(t, target.B1.IsPresent())
		assert.True(t, target.B2.IsPresent())
		assert.False(t, target.B3.IsPresent())
		assert.True(t, target.S1.IsPresent())
		assert.False(t, target.S3.IsPresent())
		assert.True(t, target.I1.IsPresent())
		assert.True(t, target.I2.IsPresent())
		assert.True(t, target.I3.IsPresent())
		assert.False(t, target.I4.IsPresent())
		assert.True(t, target.F1.IsPresent())
		assert.True(t, target.F2.IsPresent())
		assert.True(t, target.F3.IsPresent())
		assert.False(t, target.F4.IsPresent())
		assert.True(t, target.TS1.IsPresent())
		assert.False(t, target.TS2.IsPresent())
		assert.True(t, target.D1.IsPresent())
		assert.False(t, target.D2.IsPresent())
		assert.True(t, target.D3.IsPresent())
		assert.False(t, target.D4.IsPresent())
		assert.True(t, target.D5.IsPresent())
		assert.False(t, target.D6.IsPresent())

		assert.True(t, target.B1.OrElse(false))
		assert.False(t, target.B2.OrElse(true))
		assert.Equal(t, "some", target.S1.OrElse("foo"))
		assert.Equal(t, 58, target.I1.OrElse(1))
		assert.Equal(t, -9001, target.I2.OrElse(1))
		assert.Equal(t, 0, target.I3.OrElse(1))
		assert.Equal(t, 0.2, target.F1.OrElse(1))
		assert.Equal(t, -376.11, target.F2.OrElse(1))
		assert.Equal(t, float64(0), target.F3.OrElse(1))
		assert.Equal(t, time.Unix(1384923976, 0).UTC(), target.TS1.OrNow())
		assert.Equal(t, 42*time.Second, target.D1.OrElse(time.Second))
		assert.Equal(t, time.Second*9, target.D3.Get())
		assert.Equal(t, time.Second*180, target.D5.Get())

		if bts, err := yaml.Marshal(target); assert.NoError(t, err) {
			assert.Equal(t, textSourceYAML, string(bts))
		}
	}
}

var dataYAMLError = []struct {
	Source string
	Target interface {
		UnmarshalYAML(unmarshal func(interface{}) error) error
	}
}{
	{"1", &Bool{}},
	{"\"1\"", &Int{}},
	{"\"1\"", &Float64{}},
	{"\"1\"", &TimeUnixSeconds{}},
	{"\"1\"", &DurationSeconds{}},
	{"\"1\"", &DurationMillis{}},
	{"\"1\"", &DurationMinutes{}},
}

func TestYAMLError(t *testing.T) {
	for _, data := range dataYAMLError {
		t.Run(fmt.Sprint(data.Source, " ", data.Target), func(t *testing.T) {
			assert.Error(t, data.Target.UnmarshalYAML(func(i interface{}) error {
				return yaml.Unmarshal([]byte(data.Source), i)
			}))
		})
	}
}

var dataYAMLErrorDelivery = []interface {
	UnmarshalYAML(unmarshal func(interface{}) error) error
}{
	&Bool{},
	&Int{},
	&Float64{},
	&String{},
	&TimeUnixSeconds{},
	&DurationSeconds{},
	&DurationMillis{},
	&DurationMinutes{},
}

func TestYAMLErrorDelivery(t *testing.T) {
	err := errors.New("the error")
	errGenerator := func(interface{}) error {
		return err
	}

	for _, target := range dataYAMLErrorDelivery {
		t.Run(fmt.Sprint(target), func(t *testing.T) {
			assert.Equal(t, err, target.UnmarshalYAML(errGenerator))
		})
	}
}
