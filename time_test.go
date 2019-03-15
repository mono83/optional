package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_FilterZero(t *testing.T) {
	assert.Equal(t, Time{}, OfTime(time.Time{}).FilterZero())
	assert.Equal(t, OfTime(time.Unix(1, 2)), OfTime(time.Unix(1, 2)).FilterZero())
}

func TestOfTimeRef(t *testing.T) {
	var tRef *time.Time
	var iRef *int

	assert.False(t, OfTimeRef(tRef).IsPresent())
	assert.False(t, OfUnixSecondsRef(iRef).IsPresent())

	tt := time.Now()
	i := 1
	tRef = &tt
	iRef = &i

	assert.True(t, OfTimeRef(tRef).IsPresent())
	assert.True(t, OfUnixSecondsRef(iRef).IsPresent())
}
