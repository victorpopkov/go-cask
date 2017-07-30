package cask

// An ArtifactType represents a known artifact stanza type.
type ArtifactType int

// An Artifact represents the cask artifact stanza itself.
type Artifact struct {
	// Type specifies the artifact type.
	Type ArtifactType

	// Value specifies the artifact value.
	Value string
}

// Different supported artifact types.
const (
	ArtifactApp ArtifactType = iota
	ArtifactPkg
	ArtifactBinary
)

var artifactTypeNames = [...]string{
	"App",
	"Pkg",
	"Binary",
}

// NewArtifact creates a new Artifact instance and returns its pointer. Requires
// both Artifact.Type and Artifact.Value to be passed as arguments.
func NewArtifact(t ArtifactType, value string) *Artifact {
	return &Artifact{t, value}
}

// String returns the string representation of the ArtifactType.
func (t ArtifactType) String() string {
	return artifactTypeNames[t]
}
