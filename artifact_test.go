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

func TestArtifactString(t *testing.T) {
	// app
	a := NewArtifact(ArtifactApp, "Test.app")
	assert.Equal(t, "app, Test.app", a.String())
	a.Target = "Target.app"
	assert.Equal(t, "app, Test.app => Target.app", a.String())

	// pkg
	a = NewArtifact(ArtifactPkg, "test.pkg")
	assert.Equal(t, "pkg, test.pkg", a.String())
	a.AllowUntrusted = true
	assert.Equal(t, "pkg, test.pkg, allow_untrusted: true", a.String())

	// binary
	a = NewArtifact(ArtifactBinary, "test")
	assert.Equal(t, "binary, test", a.String())
	a.Target = "target"
	assert.Equal(t, "binary, test => target", a.String())
}
