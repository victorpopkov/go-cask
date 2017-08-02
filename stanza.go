package cask

// A Stanza represents a cask stanza itself. Shouldn't be used as is, but
// inherited by type specific stanzas.
type Stanza struct {
	// Value specifies the stanza value.
	Value string

	// global specifies if the appcast wasn't found inside if, which means that it
	// matches all Cask.Variants. By default, this value is "false".
	global bool
}

// A Version represents a version cask stanza.
type Version struct {
	Stanza
}

// A SHA256 represents a sha256 cask stanza.
type SHA256 struct {
	Stanza
}

// An URL represents an url cask stanza.
type URL struct {
	Stanza
}

// An Appcast represents an appcast cask stanza.
type Appcast struct {
	Stanza

	// Checkpoint specifies the checksum of the request response.
	Checkpoint string
}

// A Name represents a name cask stanza.
type Name struct {
	Stanza
}

// A Homepage represents a homepage cask stanza.
type Homepage struct {
	Stanza
}

// String returns a string representation of the Stanza struct which is the
// value.
func (s Stanza) String() string {
	return s.Value
}
