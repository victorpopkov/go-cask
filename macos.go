package cask

import "fmt"

// A MacOS represents the available macOS versions.
type MacOS int

// Different macOS releases.
const (
	MacOSHighSierra MacOS = iota
	MacOSSierra
	MacOSElCapitan
	MacOSYosemite
	MacOSMavericks
	MacOSMountainLion
	MacOSLion
	MacOSSnowLeopard
	MacOSLeopard
	MacOSTiger
)

var macOSNames = [...]string{
	"macOS High Sierra",
	"macOS Sierra",
	"OS X El Capitan",
	"OS X Yosemite",
	"OS X Mavericks",
	"OS X Mountain Lion",
	"OS X Lion",
	"Mac OS X Snow Leopard",
	"Mac OS X Leopard",
	"Mac OS X Tiger",
}

var macOSVersion = [...]string{
	"10.13",
	"10.12",
	"10.11",
	"10.10",
	"10.9",
	"10.8",
	"10.7",
	"10.6",
	"10.5",
	"10.4",
}

// Name returns the MacOS release name.
func (m MacOS) Name() string {
	return macOSNames[m]
}

// Version returns the MacOS release version.
func (m MacOS) Version() string {
	return macOSVersion[m]
}

// String returns the string representation of the MacOS release.
func (m MacOS) String() string {
	return fmt.Sprintf("%s (%s)", macOSNames[m], macOSVersion[m])
}
