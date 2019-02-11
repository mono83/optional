package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool_OrFalse(t *testing.T) {
	assert.True(t, OfBool(true).OrFalse())
	assert.False(t, emptyBool.OrFalse())
}
