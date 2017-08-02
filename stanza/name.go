package stanza

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
