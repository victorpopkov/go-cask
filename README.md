# go-cask

[![Build Status](https://travis-ci.org/victorpopkov/go-cask.svg?branch=master)](https://travis-ci.org/victorpopkov/go-cask)
[![Coverage Status](https://coveralls.io/repos/github/victorpopkov/go-cask/badge.svg?branch=master)](https://coveralls.io/github/victorpopkov/go-cask?branch=master)
[![Report Card](https://goreportcard.com/badge/github.com/victorpopkov/go-cask)](https://goreportcard.com/report/github.com/victorpopkov/go-cask)
[![GoDoc](https://godoc.org/github.com/victorpopkov/go-cask?status.svg)](https://godoc.org/github.com/victorpopkov/go-cask)

**NOTICE:** Currently in development.

This library provides functionality for working with Homebrew-Cask casks.

- [What "cask" means?](#what-cask-means)
- [Why is this library needed and what it does?](#why-is-this-library-needed-and-what-it-does)
- [Supported stanzas](#supported-stanzas)

## What "cask" means?

The "cask" is a small Ruby block in a separate file that is used to describe the
application in [Homebrew-Cask](https://github.com/caskroom/homebrew-cask)
project. You can learn more about them by reading through the
[Homebrew-Cask "Synopsis"](https://github.com/caskroom/homebrew-cask/blob/master/doc/cask_language_reference/readme.md#synopsis).

## Why is this library needed and what it does?

This library attempts to provide a way of parsing and extracting basic
information from casks for later use in Go. It parses the cask Ruby block and
creates the corresponding `cask.Cask` struct.

## Supported stanzas

Below you can find a list of all supported cask stanzas that this library can
understand and recognize during the parsing phase. If the checkbox next to the
stanza is not ticked, that stanza is not supported yet.

> To learn more about all available cask stanzas check out the
[Homebrew-Cask "All stanzas"](https://github.com/caskroom/homebrew-cask/blob/master/doc/cask_language_reference/all_stanzas.md).

### Required

- [ ] `version`
- [ ] `sha256`
- [ ] `url`
- [ ] `name`
- [ ] `homepage`

### Artifacts

- [ ] `app`
- [ ] `pkg`
- [ ] `binary`
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
- [ ] `appcast`
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

## License

Released under the [MIT License](https://opensource.org/licenses/MIT).
