package cask

import (
	"fmt"
	"regexp"
	"strings"
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

// InterpolateIntoString interpolates existing version into the provided string
// with Ruby interpolation syntax.
func (v Version) InterpolateIntoString(str string) (result string) {
	result = str

	regexInterpolations := regexp.MustCompile(`(#{version})|(#{version\.[^}]*.[^{]*})`)
	regexAllMethods := regexp.MustCompile(`(?:^#{version\.)(.*)}`)

	// find all version interpolations
	matches := regexInterpolations.FindAllStringSubmatch(str, -1)

	// for every version interpolation
	for _, m := range matches {
		match := m[0]

		// extract all methods
		methodsAll := regexAllMethods.FindAllStringSubmatch(match, -1)
		if len(methodsAll) < 1 {
			// when no methods, then it's just a version replace
			re := regexp.MustCompile(regexp.QuoteMeta(match))
			result = re.ReplaceAllString(result, v.Value)
			continue
		}

		methods := strings.Split(methodsAll[0][1], ".")

		// for every method
		part := v.Value
		for _, method := range methods {
			switch method {
			case "major":
				r, err := NewVersion(part).Major()
				if err == nil {
					part = r
				}
				break
			case "minor":
				r, err := NewVersion(part).Minor()
				if err == nil {
					part = r
				}
				break
			case "patch":
				r, err := NewVersion(part).Patch()
				if err == nil {
					part = r
				}
				break
			case "major_minor":
				r, err := NewVersion(part).MajorMinor()
				if err == nil {
					part = r
				}
				break
			case "major_minor_patch":
				r, err := NewVersion(part).MajorMinorPatch()
				if err == nil {
					part = r
				}
				break
			case "before_comma":
				r, err := NewVersion(part).BeforeComma()
				if err == nil {
					part = r
				}
				break
			case "after_comma":
				r, err := NewVersion(part).AfterComma()
				if err == nil {
					part = r
				}
				break
			case "before_colon":
				r, err := NewVersion(part).BeforeColon()
				if err == nil {
					part = r
				}
				break
			case "after_colon":
				r, err := NewVersion(part).AfterColon()
				if err == nil {
					part = r
				}
				break
			case "no_dots":
				r, err := NewVersion(part).NoDots()
				if err == nil {
					part = r
				}
				break
			case "dots_to_underscores":
				r, err := NewVersion(part).DotsToUnderscores()
				if err == nil {
					part = r
				}
				break
			case "dots_to_hyphens":
				r, err := NewVersion(part).DotsToHyphens()
				if err == nil {
					part = r
				}
				break
			default:
				// if one of the methods is unknown, then return full string without any replacements
				return result
			}
		}

		re := regexp.MustCompile(regexp.QuoteMeta(match))
		result = re.ReplaceAllString(result, part)
	}

	return result
}

// String returns a string representation of the Version struct which is the
// Version.Value.
func (v Version) String() string {
	return v.Value
}
