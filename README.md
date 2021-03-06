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
- [x] String interpolations
  - [x] `#{version}`
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
- [ ] `depends_on`
- [ ] `conflicts_with`
- [ ] `caveats`
  - [ ] indented heredoc (`<<-EOS`)
  - [ ] "squiggly" heredoc (`<<~EOS`)
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
			fmt.Printf("%10s: %s\n", "version", v.GetVersion())
			fmt.Printf("%10s: %s\n", "sha256", v.GetSHA256())
			fmt.Printf("%10s: %s\n", "url", v.GetURL())
			fmt.Printf("%10s: %s\n", "appcast", v.GetAppcast().URL)
			fmt.Printf("%10s: %v\n", "names", v.GetNames())
			fmt.Printf("%10s: %s\n", "homepage", v.GetHomepage())

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.GetArtifacts()) > 0 {
				for i, a := range v.GetArtifacts() {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.GetArtifacts())
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
	//        url: https://example.com/app_2.0.0.dmg
	//    appcast: https://example.com/sparkle/2/appcast.xml
	//      names: [Example Example One]
	//   homepage: https://example.com/
	//  artifacts: app, Example 2.0.app => Example.app
	//             app, Example 2.0 Uninstaller.app
	//             binary, #{appdir}/Example 2.0.app/Contents/MacOS/example-one => example
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
			fmt.Printf("%10s: %s\n", "version", v.GetVersion())
			fmt.Printf("%10s: %s\n", "sha256", v.GetSHA256())
			fmt.Printf("%10s: %s\n", "url", v.GetURL())
			fmt.Printf("%10s: %s\n", "appcast", v.GetAppcast().URL)
			fmt.Printf("%10s: %v\n", "names", v.GetNames())
			fmt.Printf("%10s: %s\n", "homepage", v.GetHomepage())

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.GetArtifacts()) > 0 {
				for i, a := range v.GetArtifacts() {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.GetArtifacts())
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
	//        url: https://example.com/app_1.5.0.pkg
	//    appcast: https://example.com/sparkle/1/el_capitan.xml
	//      names: [Example Example Two]
	//   homepage: https://example.com/
	//  artifacts: pkg, app_1.5.0.pkg, allow_untrusted: true
	//      macOS: Mac OS X Tiger (10.4) [minimum]
	//             OS X El Capitan (10.11) [maximum]
	// Variant #2:
	//    version: 2.0.0
	//     sha256: f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261
	//        url: https://example.com/app_2.0.0.pkg
	//    appcast: https://example.com/sparkle/2/appcast.xml
	//      names: [Example Example Two]
	//   homepage: https://example.com/
	//  artifacts: pkg, app_2.0.0.pkg, allow_untrusted: true
	//      macOS: macOS High Sierra (10.13) [minimum]
	//             macOS High Sierra (10.13) [maximum]
}
```

## License

Released under the [MIT License](https://opensource.org/licenses/MIT).
