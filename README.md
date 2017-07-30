# go-cask

[![Build Status](https://travis-ci.org/victorpopkov/go-cask.svg?branch=master)](https://travis-ci.org/victorpopkov/go-cask)
[![Coverage Status](https://coveralls.io/repos/github/victorpopkov/go-cask/badge.svg?branch=master)](https://coveralls.io/github/victorpopkov/go-cask?branch=master)
[![Report Card](https://goreportcard.com/badge/github.com/victorpopkov/go-cask)](https://goreportcard.com/badge/github.com/victorpopkov/go-cask)
[![GoDoc](https://godoc.org/github.com/victorpopkov/go-cask?status.svg)](https://godoc.org/github.com/victorpopkov/go-cask)

**NOTICE:** Currently in development.

This library provides functionality for working with Homebrew-Cask casks.

- [What "cask" means?](#what-cask-means)
- [Why is this library needed and what it does?](#why-is-this-library-needed-and-what-it-does)
- [Supported providers](#supported-providers)

## What "cask" means?

The "cask" is a small Ruby block in a separate file that is used to describe the
application in [Homebrew-Cask](https://github.com/caskroom/homebrew-cask)
project. You can learn more about them by reading through the
[Homebrew-Cask Synopsis](https://github.com/caskroom/homebrew-cask/blob/master/doc/cask_language_reference/readme.md#synopsis).

## Why is this library needed and what it does?

This library attempts to provide a way of parsing and extracting basic
information from casks for later use in Go. It parses the cask Ruby block and
creates the corresponding `cask.Cask` struct.

## License

Released under the [MIT License](https://opensource.org/licenses/MIT).
