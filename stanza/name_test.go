package stanza

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	// preparations
	n := NewName("test")

	// test
	assert.IsType(t, Name{}, *n)
	assert.False(t, n.IsGlobal)
	assert.Equal(t, "test", n.Value)
	assert.Equal(t, "test", n.String())
}
