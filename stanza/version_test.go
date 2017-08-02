package stanza

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	// preparations
	v := NewVersion("1.0.0")

	// test
	assert.IsType(t, Version{}, *v)
	assert.False(t, v.IsGlobal)
	assert.Equal(t, "1.0.0", v.Value)
	assert.Equal(t, "1.0.0", v.String())
}
