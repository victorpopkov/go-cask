package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	// preparations
	v := NewVersion("1.0.0")

	// test
	assert.IsType(t, Version{}, *v)
	assert.False(t, v.IsGlobal)
	assert.Equal(t, "1.0.0", v.Value)
	assert.Equal(t, "1.0.0", v.String())
}

func TestNewSHA256(t *testing.T) {
	// preparations
	s := NewSHA256("92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305")

	// test
	assert.IsType(t, SHA256{}, *s)
	assert.False(t, s.IsGlobal)
	assert.Equal(t, "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305", s.Value)
	assert.Equal(t, "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305", s.String())
}

func TestNewURL(t *testing.T) {
	// preparations
	u := NewURL("http://example.com/")

	// test
	assert.IsType(t, URL{}, *u)
	assert.False(t, u.IsGlobal)
	assert.Equal(t, "http://example.com/", u.Value)
	assert.Equal(t, "http://example.com/", u.String())
}

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

func TestNewHomepage(t *testing.T) {
	// preparations
	h := NewHomepage("http://example.com/")

	// test
	assert.IsType(t, Homepage{}, *h)
	assert.False(t, h.IsGlobal)
	assert.Equal(t, "http://example.com/", h.Value)
	assert.Equal(t, "http://example.com/", h.String())
}
