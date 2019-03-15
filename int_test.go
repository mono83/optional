package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt_FilterZero(t *testing.T) {
	assert.Equal(t, Int{}, OfInt(0).FilterZero())
	assert.Equal(t, OfInt(1), OfInt(1).FilterZero())
}

func TestOfIntRef(t *testing.T) {
	var iRef *int
	assert.False(t, OfIntRef(iRef).IsPresent())
	i := 1
	iRef = &i
	assert.True(t, OfIntRef(iRef).IsPresent())
}
