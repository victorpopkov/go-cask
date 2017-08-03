package cask

// A Variant represents a single cask variant.
type Variant struct {
	// Version specifies the version stanza.
	Version *Version

	// SHA256 specifies the SHA256 checksum for the downloaded Variant.URL file.
	SHA256 *SHA256

	// URL specifies the url stanza.
	URL *URL

	// Appcast specifies the appcast stanza.
	Appcast *Appcast

	// Names specify the application names. Each cask can have multiple names.
	Names []*Name

	// Homepage specifies the application vendor homepage stanza.
	Homepage string

	// Artifacts specify artifact stanzas.
	Artifacts []*Artifact

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

// AddName adds a new *Name to the Variant.Names slice.
func (v *Variant) AddName(name *Name) {
	v.Names = append(v.Names, name)
}

// AddArtifact adds a new Artifact pointer to the Variant.Artifacts slice.
func (v *Variant) AddArtifact(artifact *Artifact) {
	v.Artifacts = append(v.Artifacts, artifact)
}
