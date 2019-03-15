package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDuration_FilterZero(t *testing.T) {
	assert.Equal(t, Duration{}, OfDuration(time.Duration(0)).FilterZero())
	assert.Equal(t, OfDuration(time.Duration(1)), OfDuration(time.Duration(1)).FilterZero())
}

func TestOfDurationRef(t *testing.T) {
	var dRef *time.Duration
	var iRef *int

	assert.False(t, OfDurationRef(dRef).IsPresent())
	assert.False(t, OfMillisecondsRef(iRef).IsPresent())
	assert.False(t, OfSecondsRef(iRef).IsPresent())
	assert.False(t, OfMinutesRef(iRef).IsPresent())

	d := time.Second
	dRef = &d
	i := 1
	iRef = &i

	assert.True(t, OfDurationRef(dRef).IsPresent())
	assert.True(t, OfMillisecondsRef(iRef).IsPresent())
	assert.True(t, OfSecondsRef(iRef).IsPresent())
	assert.True(t, OfMinutesRef(iRef).IsPresent())
}
