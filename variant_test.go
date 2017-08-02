package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVariant(t *testing.T) {
	// preparations
	v := NewVariant()

	// test
	assert.IsType(t, Variant{}, *v)
	assert.Empty(t, v.Version)
	assert.Empty(t, v.SHA256)
	assert.Empty(t, v.Appcast.Value)
	assert.Empty(t, v.Appcast.Checkpoint)
	assert.Empty(t, v.URL)
	assert.Len(t, v.Names, 0)
	assert.Empty(t, v.Homepage)
	assert.Len(t, v.Artifacts, 0)
	assert.Equal(t, MacOSHighSierra, v.MinimumSupportedMacOS)
	assert.Equal(t, MacOSHighSierra, v.MaximumSupportedMacOS)
}

func TestAddName(t *testing.T) {
	// preparations
	v := NewVariant()

	// test
	assert.Len(t, v.Names, 0)
	v.AddName("test")
	assert.Len(t, v.Names, 1)
}

func TestAddArtifact(t *testing.T) {
	// preparations
	v := NewVariant()

	// test
	assert.Len(t, v.Artifacts, 0)
	v.AddArtifact(*NewArtifact(ArtifactApp, "test"))
	assert.Len(t, v.Artifacts, 1)
}
