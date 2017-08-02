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
		"if-six-versions-six-appcasts.rb": {
			Token:   "if-six-versions-six-appcasts",
			Content: string(getTestdata("if-six-versions-six-appcasts.rb")),
			Variants: []Variant{
				{
					Version: "0.1.0",
					SHA256:  "6ad9613a455798d6d92e5f5f390ab4baa70596bc869ed6b17f5cdd2b28635f06",
					URL:     "https://example.com/snowleopard/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version}/snowleopard.xml",
						},
						Checkpoint: "a93e9e53c90ab95e1ce83cbc1cbd76102da1bce5330b649872dbd95a1793a03e",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "0.2.0",
					SHA256:  "911fc0c48cb0c70601db5775a9bef1b740dc4cc9f9b46389b9f0563fe7eb94d7",
					URL:     "https://example.com/lion/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version}/lion.xml",
						},
						Checkpoint: "13dfb3758d65d265e4c12336815b2db327683ad38b2a1162cc88ab3579bbfaa1",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "0.3.0",
					SHA256:  "550613537fc488f3b372af74a4001879f012c8465b816f1b85c6d3446b2cfb49",
					URL:     "https://example.com/mountainlion/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version}/mountainlion.xml",
						},
						Checkpoint: "00af55f25d0c6e53f017a972b77fe4def95f9bb4ec4dc217c520e875fa0071a9",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "0.4.0",
					SHA256:  "cd78534ed15ad46912b71339d1417d0d043d8309c2b94415f3ed1b9d1fdfaed0",
					URL:     "https://example.com/mavericks/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version}/mavericks.xml",
						},
						Checkpoint: "9cbe5cfd22b0eb5f159ae634acf615d9c8032699b5a79d37a3046bdaf5677c84",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "0.5.0",
					SHA256:  "d1f62539db82b51da84bda2f4885db5e847db8389183be41389efd0ae6edab94",
					URL:     "https://example.com/yosemite/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version}/yosemite.xml",
						},
						Checkpoint: "f309466aea57120e04b214292d54a9d5e32d018582344b3a62021a91ed8dd69d",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "2.0.0",
					SHA256:  "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					URL:     "https://example.com/elcapitan/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version.major}/appcast.xml",
						},
						Checkpoint: "57956bd3fb23a5673e30dc83ed19d51b43e5a9235756887f3ed90662e6c68fb7",
					},
					Names: []string{
						"Example",
						"Example (if-six-versions-six-appcasts)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
		"if-three-versions-one-appcast.rb": Cask{
			Token:   "if-three-versions-one-appcast",
			Content: string(getTestdata("if-three-versions-one-appcast.rb")),
			Variants: []Variant{
				{
					Version: "0.9.0",
					SHA256:  "30c99e8b103eacbe6f6d6e1b54b06ca6d5f3164b4f50094334a517ae95ca8fba",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: nil,
					Names: []string{
						"Example",
						"Example (if-three-versions-one-appcast)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "1.0.0",
					SHA256:  "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: nil,
					Names: []string{
						"Example",
						"Example (if-three-versions-one-appcast)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "2.0.0",
					SHA256:  "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value: "https://example.com/sparkle/#{version.major}/appcast.xml",
						},
						Checkpoint: "8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d",
					},
					Names: []string{
						"Example",
						"Example (if-three-versions-one-appcast)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
		"if-two-versions-one-global-appcast.rb": Cask{
			Token:   "if-two-versions-one-global-appcast",
			Content: string(getTestdata("if-two-versions-one-global-appcast.rb")),
			Variants: []Variant{
				{
					Version: "1.0.0",
					SHA256:  "92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value:  "https://example.com/sparkle/#{version.major}/appcast.xml",
							global: true,
						},
						Checkpoint: "8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d",
					},
					Names: []string{
						"Example",
						"Example (if-two-versions-one-global-appcast)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
					Version: "2.0.0",
					SHA256:  "f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: &Appcast{
						Stanza: Stanza{
							Value:  "https://example.com/sparkle/#{version.major}/appcast.xml",
							global: true,
						},
						Checkpoint: "8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d",
					},
					Names: []string{
						"Example",
						"Example (if-two-versions-one-global-appcast)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
			Variants: []Variant{
				{
					Version: "latest",
					SHA256:  "5e1e2bcac305958b27077ca136f35f0abae7cf38c9af678f7d220ed0cb51d4f8",
					URL:     "https://example.com/app_#{version}.dmg",
					Appcast: nil,
					Names: []string{
						"Example",
						"Example (latest)",
					},
					Homepage: "https://example.com/",
					Artifacts: []Artifact{
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
			for key, actualVariant := range actualCask.Variants {
				expectedVariant := expectedCask.Variants[key]
				assert.IsType(t, Variant{}, actualVariant, filename)
				assert.Equal(t, expectedVariant.Version, actualVariant.Version, filename)
				assert.Equal(t, expectedVariant.SHA256, actualVariant.SHA256, filename)
				assert.Equal(t, expectedVariant.URL, actualVariant.URL, filename)
				assert.Equal(t, expectedVariant.Names, actualVariant.Names, filename)
				assert.Equal(t, expectedVariant.Homepage, actualVariant.Homepage, filename)
				assert.Equal(t, expectedVariant.Appcast, actualVariant.Appcast, filename)
				assert.Equal(t, expectedVariant.Artifacts, actualVariant.Artifacts, filename)
			}
		}
	}
}

func TestAddVariant(t *testing.T) {
	// preparations
	c := NewCask("")

	// test
	assert.Len(t, c.Variants, 0)
	c.AddVariant(*NewVariant())
	assert.Len(t, c.Variants, 1)
}

func TestCaskString(t *testing.T) {
	// preparations
	c := NewCask("")
	c.Token = "test"

	// test
	assert.Equal(t, "test", c.String())
}
