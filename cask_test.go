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
