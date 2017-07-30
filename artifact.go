package cask

// An Artifact represents the cask artifact stanza.
type Artifact struct {
	// Type specifies the artifact type.
	Type ArtifactType

	// Value specifies the artifact value.
	Value string
}

// NewArtifact creates a new Artifact instance and returns its pointer. Requires
// both Artifact.ArtifactType and Artifact.Value to be passed as arguments.
func NewArtifact(t ArtifactType, value string) *Artifact {
	return &Artifact{t, value}
}
