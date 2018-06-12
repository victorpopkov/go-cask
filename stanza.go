package cask

// A Stanza represents the interface that each stanza Type specific stanza
// should implement.
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

// A SHA256 represents a sha256 cask stanza.
type SHA256 struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewSHA256 creates a new SHA256 instance and returns its pointer. Requires
// SHA256.Value to be passed as argument.
func NewSHA256(value string) *SHA256 {
	return &SHA256{
		Value: value,
	}
}

// String returns a string representation of the SHA256 struct which is the
// SHA256.Value.
func (s SHA256) String() string {
	return s.Value
}

// An URL represents an url cask stanza.
type URL struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewURL creates a new URL instance and returns its pointer. Requires URL.Value
// to be passed as argument.
func NewURL(value string) *URL {
	return &URL{
		Value: value,
	}
}

// String returns a string representation of the URL struct which is the
// URL.Value.
func (u URL) String() string {
	return u.Value
}

// An Appcast represents an appcast cask stanza.
type Appcast struct {
	BaseStanza

	// URL specifies the appcast URL.
	URL string
}

// NewAppcast creates a new Appcast instance and returns its pointer. Requires
// both Appcast.URL and Appcast.Checkpoint to be passed as arguments.
func NewAppcast(url string, checkpoint string) *Appcast {
	return &Appcast{
		URL: url,
	}
}

// String returns a string representation of the Appcast struct which is the
// Appcast.URL.
func (a Appcast) String() string {
	return a.URL
}

// A Name represents a name cask stanza.
type Name struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewName creates a new Name instance and returns its pointer. Requires
// Name.Value to be passed as argument.
func NewName(value string) *Name {
	return &Name{
		Value: value,
	}
}

// String returns a string representation of the Name struct which is the
// Name.Value.
func (n Name) String() string {
	return n.Value
}

// A Homepage represents a homepage cask stanza.
type Homepage struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewHomepage creates a new Homepage instance and returns its pointer. Requires
// Homepage.Value to be passed as argument.
func NewHomepage(value string) *Homepage {
	return &Homepage{
		Value: value,
	}
}

// String returns a string representation of the Homepage struct which is the
// Homepage.Value.
func (h Homepage) String() string {
	return h.Value
}
