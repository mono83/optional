package optional

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestDuration_FilterZero(t *testing.T) {
	assert.Equal(t, emptyDuration, OfDuration(time.Duration(0)).FilterZero())
	assert.Equal(t, OfDuration(time.Duration(1)), OfDuration(time.Duration(1)).FilterZero())
}
