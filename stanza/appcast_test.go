package stanza

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppcast(t *testing.T) {
	// preparations
	a := NewAppcast(
		"http://example.com/appcast.xml",
		"2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1",
	)

	// test
	assert.IsType(t, Appcast{}, *a)
	assert.False(t, a.IsGlobal)
	assert.Equal(t, "http://example.com/appcast.xml", a.URL)
	assert.Equal(t, "2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1", a.Checkpoint)
	assert.Equal(t, "http://example.com/appcast.xml", a.String())
}
