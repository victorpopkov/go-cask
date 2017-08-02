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

// An Appcast represents an appcast cask stanza.
type Appcast struct {
	Stanza

	// Checkpoint specifies the checksum of the request response.
	Checkpoint string
}

// NewAppcast creates a new Appcast instance and returns its pointer. Requires
// both Appcast.Value and Appcast.Checkpoint to be passed as arguments.
func NewAppcast(value string, checkpoint string) *Appcast {
	return &Appcast{
		Stanza: Stanza{
			Value:  value,
			global: false,
		},
		Checkpoint: checkpoint,
	}
}
