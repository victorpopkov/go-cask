package cask

// An Appcast represents the cask appcast stanza.
type Appcast struct {
	// URL specifies the appcast URL.
	URL string

	// Checkpoint specifies the checksum of the request response.
	Checkpoint string
}

// NewAppcast creates a new Appcast instance and returns its pointer. Requires
// both Appcast.URL and Appcast.Checkpoint to be passed as arguments.
func NewAppcast(url string, checkpoint string) *Appcast {
	return &Appcast{url, checkpoint}
}

// String returns a string representation of the Appcast struct which is the
// appcast URL.
func (a Appcast) String() string {
	return a.URL
}
