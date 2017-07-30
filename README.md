# go-cask

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
