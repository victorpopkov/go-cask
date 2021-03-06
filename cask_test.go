package cask

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testdataPath = "./testdata/"

// getWorkingDir returns a current working directory path. If it's not available
// prints an error to os.Stdout and exits with error status 1.
func getWorkingDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return pwd
}

// getTestdata returns a file content as a byte array from provided testdata
// filename. If file not found, prints an error to os.Stdout and exits with exit
// status 1.
func getTestdata(filename string) []byte {
	path := filepath.Join(getWorkingDir(), testdataPath, filename)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		os.Exit(1)
	}

	return content
}

func TestNewCask(t *testing.T) {
	// preparations
	c := NewCask("")

	// test
	assert.IsType(t, Cask{}, *c)
	assert.Empty(t, c.Token)
	assert.Empty(t, c.Content)
	assert.Len(t, c.Variants, 0)
}

func TestParse(t *testing.T) {
	testCases := map[string]Cask{
		"empty.rb": {
			Token:   "empty",
			Content: string(getTestdata("empty.rb")),
			Variants: []*Variant{
				{
					Version:   nil,
					SHA256:    nil,
					URL:       nil,
					Appcast:   nil,
					Names:     nil,
					Homepage:  nil,
					Artifacts: nil,
				},
			},
		},
		"if-global-sha256-last.rb": {
			Token:   "if-global-sha256-last",
			Content: string(getTestdata("if-global-sha256-last.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						Value: "1.0.0",
					},
					SHA256: &SHA256{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac32.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-sha256-last)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-sha256-last).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac64.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-sha256-last)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-sha256-last).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-global-version-last.rb": {
			Token:   "if-global-version-last",
			Content: string(getTestdata("if-global-version-last.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac32.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-version-last)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-version-last).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "9065ae8493fa73bfdf5d29ffcd0012cd343475cf3d550ae526407b9910eb35b7",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac64.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-version-last)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-version-last).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-global-version-first.rb": {
			Token:   "if-global-version-first",
			Content: string(getTestdata("if-global-version-first.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac32.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-version-first)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-version-first).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "9065ae8493fa73bfdf5d29ffcd0012cd343475cf3d550ae526407b9910eb35b7",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}_mac64.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-global-version-first)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-global-version-first).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-no-global.rb": {
			Token:   "if-no-global",
			Content: string(getTestdata("if-no-global.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						Value: "1.0.0",
					},
					SHA256: &SHA256{
						Value: "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305",
					},
					URL: &URL{
						Value: "https://example.com/app_mavericks_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version.major}/mavericks.xml",
					},
					Names: []*Name{
						{
							Value: "Example",
						},
						{
							Value: "Example (if-no-global)",
						},
					},
					Homepage: &Homepage{
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-no-global).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					},
					URL: &URL{
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							Value: "Example",
						},
						{
							Value: "Example (if-no-global)",
						},
					},
					Homepage: &Homepage{
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-no-global).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-six-versions-six-appcasts.rb": {
			Token:   "if-six-versions-six-appcasts",
			Content: string(getTestdata("if-six-versions-six-appcasts.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						Value: "0.1.0",
					},
					SHA256: &SHA256{
						Value: "6ad9613a455798d6d92e5f5f390ab4baa70596bc869ed6b17f5cdd2b28635f06",
					},
					URL: &URL{
						Value: "https://example.com/snowleopard/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version}/snowleopard.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "0.2.0",
					},
					SHA256: &SHA256{
						Value: "911fc0c48cb0c70601db5775a9bef1b740dc4cc9f9b46389b9f0563fe7eb94d7",
					},
					URL: &URL{
						Value: "https://example.com/lion/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version}/lion.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "0.3.0",
					},
					SHA256: &SHA256{
						Value: "550613537fc488f3b372af74a4001879f012c8465b816f1b85c6d3446b2cfb49",
					},
					URL: &URL{
						Value: "https://example.com/mountainlion/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version}/mountainlion.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "0.4.0",
					},
					SHA256: &SHA256{
						Value: "cd78534ed15ad46912b71339d1417d0d043d8309c2b94415f3ed1b9d1fdfaed0",
					},
					URL: &URL{
						Value: "https://example.com/mavericks/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version}/mavericks.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "0.5.0",
					},
					SHA256: &SHA256{
						Value: "d1f62539db82b51da84bda2f4885db5e847db8389183be41389efd0ae6edab94",
					},
					URL: &URL{
						Value: "https://example.com/yosemite/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version}/yosemite.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					},
					URL: &URL{
						Value: "https://example.com/elcapitan/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-six-versions-six-appcasts)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-six-versions-six-appcasts).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-three-versions-one-appcast.rb": {
			Token:   "if-three-versions-one-appcast",
			Content: string(getTestdata("if-three-versions-one-appcast.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						Value: "0.9.0",
					},
					SHA256: &SHA256{
						Value: "30c99e8b103eacbe6f6d6e1b54b06ca6d5f3164b4f50094334a517ae95ca8fba",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: nil,
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-three-versions-one-appcast)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-three-versions-one-appcast).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "1.0.0",
					},
					SHA256: &SHA256{
						Value: "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: nil,
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-three-versions-one-appcast)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-three-versions-one-appcast).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-three-versions-one-appcast)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-three-versions-one-appcast).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"if-two-versions-one-global-appcast.rb": {
			Token:   "if-two-versions-one-global-appcast",
			Content: string(getTestdata("if-two-versions-one-global-appcast.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						Value: "1.0.0",
					},
					SHA256: &SHA256{
						Value: "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-two-versions-one-global-appcast)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-two-versions-one-global-appcast).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
				{
					Version: &Version{
						Value: "2.0.0",
					},
					SHA256: &SHA256{
						Value: "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: &Appcast{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						URL: "https://example.com/sparkle/#{version.major}/appcast.xml",
					},
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (if-two-versions-one-global-appcast)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (if-two-versions-one-global-appcast).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-if",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
		"latest.rb": {
			Token:   "latest",
			Content: string(getTestdata("latest.rb")),
			Variants: []*Variant{
				{
					Version: &Version{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "latest",
					},
					SHA256: &SHA256{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "5e1e2bcac305958b27077ca136f35f0abae7cf38c9af678f7d220ed0cb51d4f8",
					},
					URL: &URL{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/app_#{version}.dmg",
					},
					Appcast: nil,
					Names: []*Name{
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example",
						},
						{
							BaseStanza: BaseStanza{
								IsGlobal: true,
							},
							Value: "Example (latest)",
						},
					},
					Homepage: &Homepage{
						BaseStanza: BaseStanza{
							IsGlobal: true,
						},
						Value: "https://example.com/",
					},
					Artifacts: []*Artifact{
						{
							Type:           ArtifactApp,
							Value:          "Example (latest).app",
							Target:         "Example.app",
							AllowUntrusted: false,
						},
						{
							Type:           ArtifactBinary,
							Value:          "#{appdir}/Example.app/Contents/MacOS/example-latest",
							Target:         "example",
							AllowUntrusted: false,
						},
					},
				},
			},
		},
	}

	for filename, expectedCask := range testCases {
		// preparations
		actualCask := NewCask(string(getTestdata(filename)))
		err := actualCask.Parse()

		// test
		assert.IsType(t, Cask{}, *actualCask)
		assert.Nil(t, err)
		assert.Equal(t, expectedCask.Token, actualCask.Token, fmt.Sprintf(
			"%s: expected and actual cask tokens doesn't match",
			filename,
		))
		assert.Equal(t, expectedCask.Content, actualCask.Content)

		// variants
		assert.Len(t, actualCask.Variants, len(expectedCask.Variants), fmt.Sprintf(
			"%s: expected (%d) and actual (%d) cask variants length doesn't match",
			filename,
			len(expectedCask.Variants),
			len(actualCask.Variants),
		))

		if len(expectedCask.Variants) == len(actualCask.Variants) {
			for keyVariant, actualVariant := range actualCask.Variants {
				expectedVariant := expectedCask.Variants[keyVariant]
				assert.IsType(t, &Variant{}, actualVariant, filename)
				assert.Equal(t, expectedVariant.GetVersion(), actualVariant.GetVersion(), filename)
				assert.Equal(t, expectedVariant.GetSHA256(), actualVariant.GetSHA256(), filename)
				assert.Equal(t, expectedVariant.GetURL(), actualVariant.GetURL(), filename)
				assert.Equal(t, expectedVariant.GetNames(), actualVariant.GetNames(), filename)
				assert.Equal(t, expectedVariant.GetHomepage(), actualVariant.GetHomepage(), filename)
				assert.Equal(t, expectedVariant.GetAppcast(), actualVariant.GetAppcast(), filename)
				assert.Equal(t, expectedVariant.GetArtifacts(), actualVariant.GetArtifacts(), filename)
			}
		}
	}
}

func TestAddVariant(t *testing.T) {
	// preparations
	c := NewCask("")

	// test
	assert.Len(t, c.Variants, 0)
	c.AddVariant(NewVariant())
	assert.Len(t, c.Variants, 1)
}

func TestCaskString(t *testing.T) {
	// preparations
	c := NewCask("")
	c.Token = "test"

	// test
	assert.Equal(t, "test", c.String())
}
