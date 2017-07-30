package cask

// A Variant represents a single cask variant.
type Variant struct {
	// Current specifies the current cask version string.
	Version string

	// SHA256 specifies the SHA256 version checksum.
	SHA256 string

	// Appcast specifies the appcast info represented as Appcast struct.
	Appcast Appcast

	// URL specifies the url stanza value.
	URL string

	// Names specify the application names. Each cask can have multiple names.
	Names []string

	// Homepage specifies the application vendor homepage stanza.
	Homepage string

	// Artifacts specify artifact stanzas.
	Artifacts []Artifact

	// MinimumSupportedMacOS specifies the minimum supported macOS release. By
	// default each cask uses the latest stable macOS release.
	MinimumSupportedMacOS MacOS

	// MaximumSupportedMacOS specifies the maximum supported macOS release. By
	// default each cask uses the latest stable macOS release.
	MaximumSupportedMacOS MacOS
}

// NewVariant returns a new Variant instance pointer.
func NewVariant() *Variant {
	return &Variant{}
}

// AddName adds a new name string to the Variant.Names array.
func (v *Variant) AddName(name string) {
	v.Names = append(v.Names, name)
}

// AddArtifact adds a new Artifact to the Variant.Artifacts array.
func (v *Variant) AddArtifact(artifact Artifact) {
	v.Artifacts = append(v.Artifacts, artifact)
}
