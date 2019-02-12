package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt_FilterZero(t *testing.T) {
	assert.Equal(t, emptyInt, OfInt(0).FilterZero())
	assert.Equal(t, OfInt(1), OfInt(1).FilterZero())
}
