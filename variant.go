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
	Homepage *Homepage

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

// GetVersion returns the Version struct from the existing Variant.Version
// struct pointer.
func (v *Variant) GetVersion() Version {
	return *(v.Version)
}

// GetSHA256 returns the SHA256 struct from the existing Variant.SHA256 struct
// pointer.
func (v *Variant) GetSHA256() SHA256 {
	return *(v.SHA256)
}

// GetURL returns the URL struct from the existing Variant.URL struct pointer
// and interpolates the version into the Variant.URL.Value if available.
func (v *Variant) GetURL() (u URL) {
	u = *(v.URL)

	if v.Version != nil && v.Version.HasVersionStringInterpolation(u.Value) {
		u.Value = v.Version.InterpolateIntoString(u.Value)
	}

	return u
}

// GetAppcast returns the Appcast struct from the existing Variant.Appcast
// struct pointer and interpolates the version into the Variant.Appcast.URL if
// available.
func (v *Variant) GetAppcast() (a Appcast) {
	a = *(v.Appcast)

	if v.Version != nil && v.Version.HasVersionStringInterpolation(a.URL) {
		a.URL = v.Version.InterpolateIntoString(a.URL)
	}

	return a
}

// GetNames returns the []Name slice from the existing []Variant.Names slice
// pointer. and interpolates the version into each name if available.
func (v *Variant) GetNames() (n []Name) {
	for _, name := range v.Names {
		if v.Version != nil && v.Version.HasVersionStringInterpolation(name.Value) {
			name.Value = v.Version.InterpolateIntoString(name.Value)
		}

		n = append(n, *name)
	}

	return n
}
