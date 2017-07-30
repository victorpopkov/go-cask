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
}

func TestArtifactTypeString(t *testing.T) {
	assert.Equal(t, "App", ArtifactApp.String())
	assert.Equal(t, "Pkg", ArtifactPkg.String())
	assert.Equal(t, "Binary", ArtifactBinary.String())
}
