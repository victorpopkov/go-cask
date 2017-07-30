package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppcast(t *testing.T) {
	// preparations
	a := NewAppcast(
		"https://example.com/sparkle/#{version.major}/appcast.xml",
		"8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d",
	)

	// test
	assert.IsType(t, Appcast{}, *a)
	assert.Equal(t, "https://example.com/sparkle/#{version.major}/appcast.xml", a.URL)
	assert.Equal(t, "8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d", a.Checkpoint)
	assert.Equal(t, "https://example.com/sparkle/#{version.major}/appcast.xml", a.String())
}
