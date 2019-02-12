package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64_FilterZero(t *testing.T) {
	assert.Equal(t, emptyFloat64, OfFloat64(0).FilterZero())
	assert.Equal(t, OfFloat64(1), OfFloat64(1).FilterZero())
}
