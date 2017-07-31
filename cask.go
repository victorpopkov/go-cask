package cask

// A Cask represents the cask used in Homebrew-Cask.
type Cask struct {
	// Token specifies the cask token.
	Token string

	// Content specifies the string content of the loaded cask.
	Content string

	// Variants specifies all cask variants represented as an array of Variant
	// structs.
	Variants []Variant

	// parser specifies the Parser to be used for parsing the cask.
	parser *Parser
}

// NewCask creates a new Cask instance and returns its pointer.
func NewCask(content string) *Cask {
	c := new(Cask)
	c.Content = content

	l := NewLexer(c.Content)
	c.parser = NewParser(l)
	c.parser.cask = c

	return c
}

// Parse parses the cask.
func (c *Cask) Parse() error {
	return c.parser.ParseCask(c)
}

// AddVariant adds a new Variant to the Cask.Variants array.
func (c *Cask) AddVariant(variant Variant) {
	c.Variants = append(c.Variants, variant)
}

// String returns a string representation of the Cask struct which is the cask
// Token.
func (c Cask) String() string {
	return c.Token
}
