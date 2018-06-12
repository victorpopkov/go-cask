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

func TestGetAppcast(t *testing.T) {
	// preparations
	v := NewVariant()
	v.Appcast = NewAppcast(
		"http://example.com/#{version}/appcast.xml",
		"2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1",
	)

	// test (without version)
	actual := v.GetAppcast()
	assert.IsType(t, &Appcast{}, v.Appcast)
	assert.IsType(t, Appcast{}, actual)
	assert.Equal(t, v.Appcast.URL, actual.URL)

	// test (with version)
	v.Version = NewVersion("2.0.0")
	actual = v.GetAppcast()
	assert.IsType(t, &Appcast{}, v.Appcast)
	assert.IsType(t, Appcast{}, actual)
	assert.Equal(t, "http://example.com/2.0.0/appcast.xml", actual.URL)
}

func TestGetNames(t *testing.T) {
	// preparations
	v := NewVariant()
	v.Names = append(v.Names, NewName("Name"))
	v.Names = append(v.Names, NewName("Name #{version}"))

	// test (without version)
	actual := v.GetNames()
	assert.Len(t, actual, len(v.Names))
	assert.IsType(t, []Name{}, actual)
	assert.Equal(t, v.Names[0].Value, actual[0].Value)
	assert.Equal(t, v.Names[1].Value, actual[1].Value)

	// test (with version)
	v.Version = NewVersion("2.0.0")
	actual = v.GetNames()
	assert.Len(t, actual, len(v.Names))
	assert.IsType(t, []Name{}, actual)
	assert.Equal(t, v.Names[0].Value, actual[0].Value)
	assert.Equal(t, "Name 2.0.0", actual[1].Value)
}

func TestGetHomepage(t *testing.T) {
	// preparations
	v := NewVariant()
	v.Homepage = NewHomepage("http://example.com/#{version}/")

	// test (without version)
	actual := v.GetHomepage()
	assert.IsType(t, &Homepage{}, v.Homepage)
	assert.IsType(t, Homepage{}, actual)
	assert.Equal(t, v.Homepage.Value, actual.Value)

	// test (with version)
	v.Version = NewVersion("2.0.0")
	actual = v.GetHomepage()
	assert.IsType(t, &Homepage{}, v.Homepage)
	assert.IsType(t, Homepage{}, actual)
	assert.Equal(t, "http://example.com/2.0.0/", actual.Value)
}

func TestGetArtifacts(t *testing.T) {
	// preparations
	v := NewVariant()
	v.Artifacts = append(v.Artifacts, NewArtifact(ArtifactApp, "Test.app"))
	v.Artifacts = append(v.Artifacts, NewArtifact(ArtifactApp, "Test #{version}.app"))

	// test (without version)
	actual := v.GetArtifacts()
	assert.Len(t, actual, len(v.Artifacts))
	assert.IsType(t, []Artifact{}, actual)
	assert.Equal(t, v.Artifacts[0].Value, actual[0].Value)
	assert.Equal(t, v.Artifacts[1].Value, actual[1].Value)

	// test (with version)
	v.Version = NewVersion("2.0.0")
	actual = v.GetArtifacts()
	assert.Len(t, actual, len(v.Artifacts))
	assert.IsType(t, []Artifact{}, actual)
	assert.Equal(t, v.Artifacts[0].Value, actual[0].Value)
	assert.Equal(t, "Test 2.0.0.app", actual[1].Value)
}
