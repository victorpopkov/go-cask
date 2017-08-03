package cask

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

func TestNewName(t *testing.T) {
	// preparations
	n := NewName("test")

	// test
	assert.IsType(t, Name{}, *n)
	assert.False(t, n.IsGlobal)
	assert.Equal(t, "test", n.Value)
	assert.Equal(t, "test", n.String())
}

func TestNewVersion(t *testing.T) {
	// preparations
	v := NewVersion("1.0.0")

	// test
	assert.IsType(t, Version{}, *v)
	assert.False(t, v.IsGlobal)
	assert.Equal(t, "1.0.0", v.Value)
	assert.Equal(t, "1.0.0", v.String())
}
