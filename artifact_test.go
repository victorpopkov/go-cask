package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtifact(t *testing.T) {
	// preparations
	a := NewArtifact(ArtifactApp, "value")

	// test
	assert.IsType(t, Artifact{}, *a)
	assert.Equal(t, ArtifactApp, a.Type)
	assert.Equal(t, "value", a.Value)
	assert.Empty(t, a.Target)
	assert.False(t, a.AllowUntrusted)
}

func TestArtifactTypeString(t *testing.T) {
	assert.Equal(t, "app", ArtifactApp.String())
	assert.Equal(t, "pkg", ArtifactPkg.String())
	assert.Equal(t, "binary", ArtifactBinary.String())
}
