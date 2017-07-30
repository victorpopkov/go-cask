package cask

// ArtifactType represents a known token type.
type ArtifactType int

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

// String returns the string representation of the MacOS release.
func (t ArtifactType) String() string {
	return artifactTypeNames[t]
}
