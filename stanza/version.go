package stanza

// A Version represents a version cask stanza.
type Version struct {
	BaseStanza

	// Value specifies the stanza value.
	Value string
}

// NewVersion creates a new Version instance and returns its pointer. Requires
// Version.Value to be passed as argument.
func NewVersion(value string) *Version {
	return &Version{
		Value: value,
	}
}

// String returns a string representation of the Version struct which is the
// Version.Value.
func (v Version) String() string {
	return v.Value
}
