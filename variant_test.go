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
	assert.Nil(t, v.Appcast)
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
	v.AddName(NewName("test"))
	assert.Len(t, v.Names, 1)
}

func TestAddArtifact(t *testing.T) {
	// preparations
	v := NewVariant()

	// test
	assert.Len(t, v.Artifacts, 0)
	v.AddArtifact(NewArtifact(ArtifactApp, "test"))
	assert.Len(t, v.Artifacts, 1)
}

func TestGetVersion(t *testing.T) {
	// preparations
	v := NewVariant()
	v.Version = NewVersion("2.0.0")

	// test
	actual := v.GetVersion()
	assert.IsType(t, &Version{}, v.Version)
	assert.IsType(t, Version{}, actual)
	assert.Equal(t, v.Version.Value, actual.Value)
}

func TestGetSHA256(t *testing.T) {
	// preparations
	v := NewVariant()
	v.SHA256 = NewSHA256("92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305")

	// test
	actual := v.GetSHA256()
	assert.IsType(t, &SHA256{}, v.SHA256)
	assert.IsType(t, SHA256{}, actual)
	assert.Equal(t, v.SHA256.Value, actual.Value)
}

func TestGetURL(t *testing.T) {
	// preparations
	v := NewVariant()
	v.URL = NewURL("http://example.com/#{version}.dmg")

	// test (without version)
	actual := v.GetURL()
	assert.IsType(t, &URL{}, v.URL)
	assert.IsType(t, URL{}, actual)
	assert.Equal(t, v.URL.Value, actual.Value)

	// test (with version)
	v.Version = NewVersion("2.0.0")
	actual = v.GetURL()
	assert.IsType(t, &URL{}, v.URL)
	assert.IsType(t, URL{}, actual)
	assert.Equal(t, "http://example.com/2.0.0.dmg", actual.Value)
}
