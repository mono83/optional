package optional

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var textSource = `{
  "b1": true,
  "b2": false,
  "b3": null,
  "s1": "some",
  "s2": "",
  "s3": null,
  "i1": 58,
  "i2": -9001,
  "i3": 0,
  "i4": null,
  "f1": 0.2,
  "f2": -376.11,
  "f3": 0,
  "f4": null
}`

func TestJSONMarshalling(t *testing.T) {
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
	}

	if err := json.Unmarshal([]byte(textSource), &target); assert.NoError(t, err) {
		assert.True(t, target.B1.IsPresent())
		assert.True(t, target.B2.IsPresent())
		assert.False(t, target.B3.IsPresent())
		assert.True(t, target.S1.IsPresent())
		assert.True(t, target.S2.IsPresent())
		assert.False(t, target.S3.IsPresent())
		assert.True(t, target.I1.IsPresent())
		assert.True(t, target.I2.IsPresent())
		assert.True(t, target.I3.IsPresent())
		assert.False(t, target.I4.IsPresent())
		assert.True(t, target.F1.IsPresent())
		assert.True(t, target.F2.IsPresent())
		assert.True(t, target.F3.IsPresent())
		assert.False(t, target.F4.IsPresent())

		assert.True(t, target.B1.OrElse(false))
		assert.False(t, target.B2.OrElse(true))
		assert.Equal(t, "some", target.S1.OrElse("foo"))
		assert.Equal(t, "", target.S2.OrElse("foo"))
		assert.Equal(t, 58, target.I1.OrElse(1))
		assert.Equal(t, -9001, target.I2.OrElse(1))
		assert.Equal(t, 0, target.I3.OrElse(1))
		assert.Equal(t, 0.2, target.F1.OrElse(1))
		assert.Equal(t, -376.11, target.F2.OrElse(1))
		assert.Equal(t, float64(0), target.F3.OrElse(1))

		if bts, err := json.MarshalIndent(target, "", "  "); assert.NoError(t, err) {
			assert.Equal(t, textSource, string(bts))
		}
	}
}
