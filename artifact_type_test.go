package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifactTypeString(t *testing.T) {
	assert.Equal(t, "App", ArtifactApp.String())
	assert.Equal(t, "Pkg", ArtifactPkg.String())
	assert.Equal(t, "Binary", ArtifactBinary.String())
}
