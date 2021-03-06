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
	if v.Version != nil {
		return *(v.Version)
	}

	return Version{}
}

// GetSHA256 returns the SHA256 struct from the existing Variant.SHA256 struct
// pointer.
func (v *Variant) GetSHA256() SHA256 {
	if v.SHA256 != nil {
		return *(v.SHA256)
	}

	return SHA256{}
}

// GetURL returns the URL struct from the existing Variant.URL struct pointer
// and interpolates the version into the Variant.URL.Value if available.
func (v *Variant) GetURL() (u URL) {
	if v.URL != nil {
		u = *(v.URL)

		if v.Version != nil && v.Version.HasVersionStringInterpolation(u.Value) {
			u.Value = v.Version.InterpolateIntoString(u.Value)
		}

		return u
	}

	return URL{}
}

// GetAppcast returns the Appcast struct from the existing Variant.Appcast
// struct pointer and interpolates the version into the Variant.Appcast.URL if
// available.
func (v *Variant) GetAppcast() (a Appcast) {
	if v.Appcast != nil {
		a = *(v.Appcast)

		if v.Version != nil && v.Version.HasVersionStringInterpolation(a.URL) {
			a.URL = v.Version.InterpolateIntoString(a.URL)
		}

		return a
	}

	return Appcast{}
}

// GetNames returns the []Name slice from the existing []Variant.Names slice
// pointer and interpolates the version into each name if available.
func (v *Variant) GetNames() (n []Name) {
	for _, name := range v.Names {
		newName := *name

		if v.Version != nil && v.Version.HasVersionStringInterpolation(name.Value) {
			newName.Value = v.Version.InterpolateIntoString(name.Value)
		}

		n = append(n, newName)
	}

	return n
}

// GetHomepage returns the Homepage struct from the existing Variant.Homepage
// struct pointer and interpolates the version into the Variant.Homepage.Value
// if available.
func (v *Variant) GetHomepage() (h Homepage) {
	if v.Homepage != nil {
		h = *(v.Homepage)

		if v.Version != nil && v.Version.HasVersionStringInterpolation(h.Value) {
			h.Value = v.Version.InterpolateIntoString(h.Value)
		}

		return h
	}

	return Homepage{}
}

// GetArtifacts returns the []Artifacts slice from the existing
// []Variant.Artifacts slice pointer and interpolates the version into each
// artifact value if available.
func (v *Variant) GetArtifacts() (a []Artifact) {
	for _, artifact := range v.Artifacts {
		newArtifact := *artifact

		if v.Version != nil && v.Version.HasVersionStringInterpolation(artifact.Value) {
			newArtifact.Value = v.Version.InterpolateIntoString(artifact.Value)
		}

		a = append(a, newArtifact)
	}

	return a
}
