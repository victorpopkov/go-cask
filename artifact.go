package cask

import "fmt"

// An ArtifactType represents a known artifact stanza type.
type ArtifactType int

// An Artifact represents the cask artifact stanza itself.
type Artifact struct {
	// Type specifies the artifact type.
	Type ArtifactType

	// Value specifies the artifact value.
	Value string

	// Target specifies the "target:" value. By default, it's empty string.
	Target string

	// AllowUntrusted specifies the "allow_untrusted:" value. By default, it's
	// false. This should be true only if the Artifact.Type is ArtifactPkg.
	AllowUntrusted bool
}

// Different supported artifact types.
const (
	ArtifactApp ArtifactType = iota
	ArtifactPkg
	ArtifactBinary
)

var artifactTypeNames = [...]string{
	"app",
	"pkg",
	"binary",
}

// NewArtifact creates a new Artifact instance and returns its pointer. Requires
// both Artifact.Type and Artifact.Value to be passed as arguments.
func NewArtifact(t ArtifactType, value string) *Artifact {
	return &Artifact{t, value, "", false}
}

// String returns the string representation of the ArtifactType.
func (t ArtifactType) String() string {
	return artifactTypeNames[t]
}

// String returns the string representation of the Artifact.
func (a Artifact) String() (result string) {
	switch a.Type {
	case ArtifactApp:
		result = fmt.Sprintf("%s, %s", a.Type.String(), a.Value)
		if a.Target != "" {
			result += fmt.Sprintf(" => %s", a.Target)
		}
	case ArtifactPkg:
		result = fmt.Sprintf("%s, %s", a.Type.String(), a.Value)
		if a.AllowUntrusted {
			result += ", allow_untrusted: true"
		}
	case ArtifactBinary:
		result = fmt.Sprintf("%s, %s", a.Type.String(), a.Value)
		if a.Target != "" {
			result += fmt.Sprintf(" => %s", a.Target)
		}
	}

	return result
}
