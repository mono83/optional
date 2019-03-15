package optional

import (
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var textSourceTOML = `
  "b1" = "true"
  "b2" = "false"
  "s1" = "some"
  "s2" = ""
  "i1" = "58"
  "i2" = "-9001"
  "i3" = "0"
  "f1" = 0.2
  "f2" = -376.11
  "f3" = 0
  "ts1" = 1384923976
  "d1" = 42
  "d3" = 9000
  "d5" = 3
`

func TestTOMLMarshalling(t *testing.T) {
	var target struct {
		B1 Bool `json:"b1"`
		B2 Bool `json:"b2"`
		B3 Bool `json:"b3"`

		S1 String `json:"s1"`
		S2 String `json:"s2"`
		S3 String `json:"s3"`

		I1 Int `json:"i1"`
		I2 Int `json:"i2"`
		I3 Int `json:"i3"`
		I4 Int `json:"i4"`

		F1 Float64 `json:"f1"`
		F2 Float64 `json:"f2"`
		F3 Float64 `json:"f3"`
		F4 Float64 `json:"f4"`

		TS1 TimeUnixSeconds `json:"ts1"`
		TS2 TimeUnixSeconds `json:"ts2"`

		D1 DurationSeconds `json:"d1"`
		D2 DurationSeconds `json:"d2"`
		D3 DurationMillis  `json:"d3"`
		D4 DurationMillis  `json:"d4"`
		D5 DurationMinutes `json:"d5"`
		D6 DurationMinutes `json:"d6"`
	}

	if err := toml.Unmarshal([]byte(textSourceTOML), &target); assert.NoError(t, err) {
		assert.True(t, target.B1.IsPresent())
		assert.True(t, target.B2.IsPresent())
		assert.False(t, target.B3.IsPresent())
		assert.True(t, target.S1.IsPresent())
		assert.False(t, target.S2.IsPresent())
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
	}
}
