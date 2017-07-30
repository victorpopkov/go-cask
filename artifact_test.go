package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArtifact(t *testing.T) {
	a := NewArtifact(ArtifactApp, "value")
	assert.IsType(t, Artifact{}, *a)
	assert.Equal(t, ArtifactApp, a.Type)
	assert.Equal(t, "value", a.Value)
}
