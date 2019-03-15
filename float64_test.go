package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64_FilterZero(t *testing.T) {
	assert.Equal(t, Float64{}, OfFloat64(0).FilterZero())
	assert.Equal(t, OfFloat64(1), OfFloat64(1).FilterZero())
}

func TestOfFloat64Ref(t *testing.T) {
	var fRef *float64
	assert.False(t, OfFloat64Ref(fRef).IsPresent())
	f := 1.1
	fRef = &f
	assert.True(t, OfFloat64Ref(fRef).IsPresent())
}
