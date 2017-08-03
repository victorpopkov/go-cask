package cask

// Stanza is the interface which represents stanza.
type Stanza interface {
	String() string
}

// A BaseStanza represents a base for all stanzas. Shouldn't be used as is, but
// inherited by type specific stanzas.
type BaseStanza struct {
	// IsGlobal specifies if the appcast belongs to all Cask.Variants. If the
	// stanza wasn't found inside if statement, the stanza should be considered as
	// global and this value should be true. By default, this value is "false".
	IsGlobal bool
}

// An Appcast represents an appcast cask stanza.
type Appcast struct {
	BaseStanza

	// URL specifies the appcast URL.
	URL string

	// Checkpoint specifies the checksum of the request response.
	Checkpoint string
}

// A Name represents a name cask stanza.
type Name struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// A Version represents a version cask stanza.
type Version struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewAppcast creates a new Appcast instance and returns its pointer. Requires
// both Appcast.URL and Appcast.Checkpoint to be passed as arguments.
func NewAppcast(url string, checkpoint string) *Appcast {
	return &Appcast{
		URL:        url,
		Checkpoint: checkpoint,
	}
}

// NewName creates a new Name instance and returns its pointer. Requires
// Name.Value to be passed as argument.
func NewName(value string) *Name {
	return &Name{
		Value: value,
	}
}

// NewVersion creates a new Version instance and returns its pointer. Requires
// Version.Value to be passed as argument.
func NewVersion(value string) *Version {
	return &Version{
		Value: value,
	}
}

// String returns a string representation of the Appcast struct which is the
// Appcast.URL.
func (a Appcast) String() string {
	return a.URL
}

// String returns a string representation of the Name struct which is the
// Name.Value.
func (n Name) String() string {
	return n.Value
}

// String returns a string representation of the Version struct which is the
// Version.Value.
func (v Version) String() string {
	return v.Value
}
