# go-cask

[![Build Status](https://travis-ci.org/victorpopkov/go-cask.svg?branch=master)](https://travis-ci.org/victorpopkov/go-cask)
[![Coverage Status](https://coveralls.io/repos/github/victorpopkov/go-cask/badge.svg?branch=master)](https://coveralls.io/github/victorpopkov/go-cask?branch=master)
[![Report Card](https://goreportcard.com/badge/github.com/victorpopkov/go-cask)](https://goreportcard.com/report/github.com/victorpopkov/go-cask)
[![GoDoc](https://godoc.org/github.com/victorpopkov/go-cask?status.svg)](https://godoc.org/github.com/victorpopkov/go-cask)

**NOTICE:** Currently in development.

This library provides functionality for working with [Homebrew-Cask](https://github.com/caskroom/homebrew-cask)
casks.

- [What "cask" means?](#what-cask-means)
- [Why is this library needed and what it does?](#why-is-this-library-needed-and-what-it-does)
- [Supported stanzas](#supported-stanzas)
- [Examples](#examples)

## What "cask" means?

The "cask" is a small Ruby block in a separate file that is used to describe the
application in [Homebrew-Cask](https://github.com/caskroom/homebrew-cask)
project. You can learn more about them by reading through the
[Homebrew-Cask "Synopsis"](https://github.com/caskroom/homebrew-cask/blob/master/doc/cask_language_reference/readme.md#synopsis).

## Why is this library needed and what it does?

This library attempts to provide a way of parsing and extracting basic
information from casks for later use in Go. It parses the cask Ruby block and
creates the corresponding `cask.Cask` struct.

### Features

- [x] Conditional statements
  - [x] MacOS.version
- [ ] Language blocks
- [ ] String interpolations
  - [ ] `#{version}`
  - [ ] `#{language}`

## Supported stanzas

Below you can find a list of all supported cask stanzas that this library can
understand and recognize during the parsing phase. If the checkbox next to the
stanza is not ticked, that stanza is not supported yet.

> To learn more about all available cask stanzas check out the
[Homebrew-Cask "All stanzas"](https://github.com/caskroom/homebrew-cask/blob/master/doc/cask_language_reference/all_stanzas.md).

### Required

- [x] `version`
- [x] `sha256`
- [x] `url`
- [x] `name`
- [x] `homepage`

### Artifacts

- [x] `app`
  - [x] `target:`
- [x] `pkg`
  - [x] `allow_untrusted:`
- [x] `binary`
  - [x] `target:`
- [ ] `colorpicker`
- [ ] `dictionary`
- [ ] `font`
- [ ] `input_method`
- [ ] `internet_plugin`
- [ ] `prefpane`
- [ ] `qlplugin`
- [ ] `screen_saver`
- [ ] `service`
- [ ] `audio_unit_plugin`
- [ ] `vst_plugin`
- [ ] `vst3_plugin`
- [ ] `suite`
- [ ] `artifact`
- [ ] `installer`
- [ ] `stage_only`

### Optional

- [ ] `uninstall`
- [ ] `zap`
- [x] `appcast`
  - [x] `checkpoint:`
- [ ] `depends_on`
- [ ] `conflicts_with`
- [ ] `caveats`
- [ ] `preflight`
- [ ] `postflight`
- [ ] `uninstall_preflight`
- [ ] `uninstall_postflight`
- [ ] `language`
- [ ] `accessibility_access`
- [ ] `container nested:`
- [ ] `container type:`
- [ ] `gpg`
- [ ] `auto_updates`

## Examples

### First example

For the first example we will parse the [example-one.rb](https://github.com/victorpopkov/go-cask/blob/master/testdata/example-one.rb)
cask from our testdata directory.

```go
package main

import (
	"fmt"

	"github.com/victorpopkov/go-cask"
)

func main() {
	// for this example we will load the cask from our testdata directory
	content := string(getTestdata("example-one.rb"))

	// example
	c := cask.NewCask(content)
	err := c.Parse()

	if err == nil {
		fmt.Println("Token:", c.Token)
		for i, v := range c.Variants {
			fmt.Printf("Variant #%d:\n", i+1)
			fmt.Printf("%10s: %s\n", "version", v.Version)
			fmt.Printf("%10s: %s\n", "sha256", v.SHA256)
			fmt.Printf("%10s: %s\n", "url", v.URL)
			fmt.Printf("%10s: %s\n", "appcast", v.Appcast.URL)
			fmt.Printf("%12s%s\n", "", v.Appcast.Checkpoint)
			fmt.Printf("%10s: %v\n", "names", v.Names)
			fmt.Printf("%10s: %s\n", "homepage", v.Homepage)

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.Artifacts) > 0 {
				for i, a := range v.Artifacts {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.Artifacts)
			}

			// macOS
			fmt.Printf("%10s: %s [minimum]\n", "macOS", v.MinimumSupportedMacOS)
			fmt.Printf("%12s%s [maximum]\n", "", v.MaximumSupportedMacOS)
		}
	}

	// Output:
	// Token: example-one
	// Variant #1:
	//    version: 2.0.0
	//     sha256: f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261
	//        url: https://example.com/app_#{version}.dmg
	//    appcast: https://example.com/sparkle/#{version.major}/appcast.xml
	//             8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d
	//      names: [Example Example One]
	//   homepage: https://example.com/
	//  artifacts: app, Example One.app => Example.app
	//             app, Example One Uninstaller.app
	//             binary, #{appdir}/Example.app/Contents/MacOS/example-one => example
	//      macOS: macOS High Sierra (10.13) [minimum]
	//             macOS High Sierra (10.13) [maximum]
}
```

### Second example

For the second example we will parse the [example-two.rb](https://github.com/victorpopkov/go-cask/blob/master/testdata/example-two.rb)
cask from our testdata directory.

```go
package main

import (
	"fmt"

	"github.com/victorpopkov/go-cask"
)

func main() {
	// for this example we will load the cask from our testdata directory
	content := string(getTestdata("example-two.rb"))

	// example
	c := cask.NewCask(content)
	err := c.Parse()

	if err == nil {
		fmt.Println("Token:", c.Token)
		for i, v := range c.Variants {
			fmt.Printf("Variant #%d:\n", i+1)
			fmt.Printf("%10s: %s\n", "version", v.Version)
			fmt.Printf("%10s: %s\n", "sha256", v.SHA256)
			fmt.Printf("%10s: %s\n", "url", v.URL)
			fmt.Printf("%10s: %s\n", "appcast", v.Appcast.URL)
			fmt.Printf("%12s%s\n", "", v.Appcast.Checkpoint)
			fmt.Printf("%10s: %v\n", "names", v.Names)
			fmt.Printf("%10s: %s\n", "homepage", v.Homepage)

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.Artifacts) > 0 {
				for i, a := range v.Artifacts {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.Artifacts)
			}

			// macOS
			fmt.Printf("%10s: %s [minimum]\n", "macOS", v.MinimumSupportedMacOS)
			fmt.Printf("%12s%s [maximum]\n", "", v.MaximumSupportedMacOS)
		}
	}

	// Output:
	// Token: example-two
	// Variant #1:
	//    version: 1.5.0
	//     sha256: 1f4dc096d58f7d21e3875671aee6f29b120ab84218fa47db2cb53bc9eb5b4dac
	//        url: https://example.com/app_#{version}.pkg
	//    appcast: https://example.com/sparkle/#{version}/el_capitan.xml
	//             93ef3101ca730028d70524f71b7f6f17cbdb8d26906299f90c38b7079e1d03a4
	//      names: [Example Example Two]
	//   homepage: https://example.com/
	//  artifacts: pkg, app_#{version}.pkg, allow_untrusted: true
	//      macOS: Mac OS X Tiger (10.4) [minimum]
	//             OS X El Capitan (10.11) [maximum]
	// Variant #2:
	//    version: 2.0.0
	//     sha256: f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261
	//        url: https://example.com/app_#{version}.pkg
	//    appcast: https://example.com/sparkle/#{version.major}/appcast.xml
	//             57956bd3fb23a5673e30dc83ed19d51b43e5a9235756887f3ed90662e6c68fb7
	//      names: [Example Example Two]
	//   homepage: https://example.com/
	//  artifacts: pkg, app_#{version}.pkg, allow_untrusted: true
	//      macOS: macOS High Sierra (10.13) [minimum]
	//             macOS High Sierra (10.13) [maximum]
}
```

## License

Released under the [MIT License](https://opensource.org/licenses/MIT).
