package stanza

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
