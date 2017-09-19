package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestVersion() *Version {
	return NewVersion("1.2.3,1000:400")
}

func createTestVersionInvalid() *Version {
	return NewVersion("invalid")
}

func TestNewVersion(t *testing.T) {
	// preparations
	v := NewVersion("1.0.0")

	// test
	assert.IsType(t, Version{}, *v)
	assert.False(t, v.IsGlobal)
	assert.Equal(t, "1.0.0", v.Value)
	assert.Equal(t, "1.0.0", v.String())
}

func TestHasVersionStringInterpolation(t *testing.T) {
	var testCases = map[string]bool{
		"#{version}":       true,
		"#{version.major}": true,

		"#{version": false,
		"#version":  false,
		"invalid":   false,
	}

	for testCase, expected := range testCases {
		// preparations
		v := NewVersion("1.0.0")

		// test
		assert.Equal(t, expected, v.HasVersionStringInterpolation(testCase))
	}
}

func TestMajor(t *testing.T) {
	// test
	actual, err := createTestVersion().Major()
	assert.Equal(t, "1", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().Major()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no Major() match`, err.Error())
}

func TestMinor(t *testing.T) {
	// test
	actual, err := createTestVersion().Minor()
	assert.Equal(t, "2", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().Minor()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no Minor() match`, err.Error())
}

func TestPatch(t *testing.T) {
	// test
	actual, err := createTestVersion().Patch()
	assert.Equal(t, "3", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().Patch()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no Patch() match`, err.Error())
}

func TestMajorMinor(t *testing.T) {
	// test
	actual, err := createTestVersion().MajorMinor()
	assert.Equal(t, "1.2", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().MajorMinor()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no MajorMinor() match`, err.Error())
}

func TestMajorMinorPatch(t *testing.T) {
	// test
	actual, err := createTestVersion().MajorMinorPatch()
	assert.Equal(t, "1.2.3", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().MajorMinorPatch()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no MajorMinorPatch() match`, err.Error())
}

func TestBeforeComma(t *testing.T) {
	// test
	actual, err := createTestVersion().BeforeComma()
	assert.Equal(t, "1.2.3", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().BeforeComma()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no BeforeComma() match`, err.Error())
}

func TestAfterComma(t *testing.T) {
	// test
	actual, err := createTestVersion().AfterComma()
	assert.Equal(t, "1000:400", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().AfterComma()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no AfterComma() match`, err.Error())
}

func TestBeforeColon(t *testing.T) {
	// test
	actual, err := createTestVersion().BeforeColon()
	assert.Equal(t, "1.2.3,1000", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().BeforeColon()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no BeforeColon() match`, err.Error())
}

func TestAfterColon(t *testing.T) {
	// test
	actual, err := createTestVersion().AfterColon()
	assert.Equal(t, "400", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().AfterColon()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no AfterColon() match`, err.Error())
}

func TestNoDots(t *testing.T) {
	// test
	actual, err := createTestVersion().NoDots()
	assert.Equal(t, "123,1000:400", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().NoDots()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no NoDots() match`, err.Error())
}

func TestDotsToUnderscores(t *testing.T) {
	// test
	actual, err := createTestVersion().DotsToUnderscores()
	assert.Equal(t, "1_2_3,1000:400", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().DotsToUnderscores()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no DotsToUnderscores() match`, err.Error())
}

func TestDotsToHyphens(t *testing.T) {
	// test
	actual, err := createTestVersion().DotsToHyphens()
	assert.Equal(t, "1-2-3,1000:400", actual)
	assert.Nil(t, err)

	// test (error)
	actual, err = createTestVersionInvalid().DotsToHyphens()
	assert.Empty(t, actual)
	assert.Error(t, err)
	assert.Equal(t, `version "invalid": no DotsToHyphens() match`, err.Error())
}

func TestInterpolateIntoString(t *testing.T) {
	testCases := map[string]string{
		"#{version}": "1.2.3,1000:400",

		// semantic
		"#{version.major}":             "1",
		"#{version.minor}":             "2",
		"#{version.patch}":             "3",
		"#{version.major_minor}":       "1.2",
		"#{version.major_minor_patch}": "1.2.3",

		// before & after
		"#{version.before_comma}": "1.2.3",
		"#{version.after_comma}":  "1000:400",
		"#{version.before_colon}": "1.2.3,1000",
		"#{version.after_colon}":  "400",

		// dots
		"#{version.no_dots}":             "123,1000:400",
		"#{version.dots_to_underscores}": "1_2_3,1000:400",
		"#{version.dots_to_hyphens}":     "1-2-3,1000:400",

		// multiple
		"#{version.major} #{version.minor} #{version.patch}": "1 2 3",

		// chained
		"#{version.before_colon.before_comma.no_dots}": "123",

		// when unknown method (shouldn't change at all)
		"#{version.unknown}":                      "#{version.unknown}",
		"#{version.before_colon.unknown.no_dots}": "#{version.before_colon.unknown.no_dots}",
	}

	for content, interpolated := range testCases {
		actual := createTestVersion().InterpolateIntoString(content)
		assert.Equal(t, interpolated, actual)
	}
}
