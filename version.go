package cask

import (
	"fmt"
	"regexp"
)

// A Version represents a version cask stanza.
type Version struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewVersion creates a new Version instance and returns its pointer. Requires
// Version.Value to be passed as argument.
func NewVersion(value string) *Version {
	return &Version{
		Value: value,
	}
}

// Major extracts the major semantic version part from Version.Value and returns
// the result string.
func (v Version) Major() (string, error) {
	re := regexp.MustCompile(`^\d`)
	if re.MatchString(v.Value) {
		return re.FindAllString(v.Value, -1)[0], nil
	}
	return "", fmt.Errorf(`version "%s": no Major() match`, v.Value)
}

// Minor extracts the minor semantic version part from Version.Value and returns
// the result string.
func (v Version) Minor() (string, error) {
	re := regexp.MustCompile(`^(?:\d)\.(\d)`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no Minor() match`, v.Value)
}

// Patch extracts the patch semantic version part from Version.Value and returns
// the result string.
func (v Version) Patch() (string, error) {
	re := regexp.MustCompile(`^(?:\d)\.(?:\d)\.(\d)`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no Patch() match`, v.Value)
}

// MajorMinor extracts the major and minor semantic version parts from
// Version.Value and returns the result string.
func (v Version) MajorMinor() (string, error) {
	re := regexp.MustCompile(`^((?:\d)\.(?:\d))`)
	if re.MatchString(v.Value) {
		return re.FindAllString(v.Value, -1)[0], nil
	}
	return "", fmt.Errorf(`version "%s": no MajorMinor() match`, v.Value)
}

// MajorMinorPatch extracts the major, minor and patch semantic version parts
// from Version.Value and returns the result string.
func (v Version) MajorMinorPatch() (string, error) {
	re := regexp.MustCompile(`^((?:\d)\.(?:\d)\.(?:\d))`)
	if re.MatchString(v.Value) {
		return re.FindAllString(v.Value, -1)[0], nil
	}
	return "", fmt.Errorf(`version "%s": no MajorMinorPatch() match`, v.Value)
}

// BeforeComma extracts the Version.Value part before comma and returns the
// result string.
func (v Version) BeforeComma() (string, error) {
	re := regexp.MustCompile(`(^.*)\,`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no BeforeComma() match`, v.Value)
}

// AfterComma extracts the Version.Value part after comma and returns the result
// string.
func (v Version) AfterComma() (string, error) {
	re := regexp.MustCompile(`\,(.*$)`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no AfterComma() match`, v.Value)
}

// BeforeColon extracts the Version.Value part before colon and returns the
// result string.
func (v Version) BeforeColon() (string, error) {
	re := regexp.MustCompile(`(^.*)\:`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no BeforeColon() match`, v.Value)
}

// AfterColon extracts the Version.Value part after colon and returns the result
// string.
func (v Version) AfterColon() (string, error) {
	re := regexp.MustCompile(`\:(.*$)`)
	if re.MatchString(v.Value) {
		return re.FindAllStringSubmatch(v.Value, -1)[0][1], nil
	}
	return "", fmt.Errorf(`version "%s": no AfterColon() match`, v.Value)
}

// NoDots removes all Version.Value dots and returns the result string.
func (v Version) NoDots() (string, error) {
	re := regexp.MustCompile(`\.`)
	if re.MatchString(v.Value) {
		return re.ReplaceAllString(v.Value, ""), nil
	}
	return "", fmt.Errorf(`version "%s": no NoDots() match`, v.Value)
}

// DotsToUnderscores converts all Version.Value dots to underscores and returns
// the result string.
func (v Version) DotsToUnderscores() (string, error) {
	re := regexp.MustCompile(`\.`)
	if re.MatchString(v.Value) {
		return re.ReplaceAllString(v.Value, "_"), nil
	}
	return "", fmt.Errorf(`version "%s": no DotsToUnderscores() match`, v.Value)
}

// DotsToHyphens converts all Version.Value dots to hyphens and returns the
// result string.
func (v Version) DotsToHyphens() (string, error) {
	re := regexp.MustCompile(`\.`)
	if re.MatchString(v.Value) {
		return re.ReplaceAllString(v.Value, "-"), nil
	}
	return "", fmt.Errorf(`version "%s": no DotsToHyphens() match`, v.Value)
}

// String returns a string representation of the Version struct which is the
// Version.Value.
func (v Version) String() string {
	return v.Value
}
