package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Trim(t *testing.T) {
	assert.Equal(t, OfString("foo"), OfString(" \t foo\r\n").Trim())
	assert.Equal(t, OfString(""), OfString("  ").Trim())
}

func TestString_FilterEmpty(t *testing.T) {
	assert.Equal(t, OfString(" foo"), OfString(" foo").FilterEmpty())
	assert.Equal(t, OfString(" "), OfString(" ").FilterEmpty())
	assert.Equal(t, emptyString, OfString("").FilterEmpty())
}
