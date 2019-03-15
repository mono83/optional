package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_OrFalse(t *testing.T) {
	assert.True(t, OfBool(true).OrFalse())
	assert.False(t, Bool{}.OrFalse())
}

func TestOfBoolRef(t *testing.T) {
	var bRef *bool
	assert.False(t, OfBoolRef(bRef).IsPresent())
	b := false
	bRef = &b
	assert.True(t, OfBoolRef(bRef).IsPresent())
}
