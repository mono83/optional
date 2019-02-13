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
