package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacOSName(t *testing.T) {
	assert.Equal(t, "macOS High Sierra", MacOSHighSierra.Name())
	assert.Equal(t, "macOS Sierra", MacOSSierra.Name())
	assert.Equal(t, "OS X El Capitan", MacOSElCapitan.Name())
	assert.Equal(t, "OS X Yosemite", MacOSYosemite.Name())
	assert.Equal(t, "OS X Mavericks", MacOSMavericks.Name())
	assert.Equal(t, "OS X Mountain Lion", MacOSMountainLion.Name())
	assert.Equal(t, "OS X Lion", MacOSLion.Name())
	assert.Equal(t, "Mac OS X Snow Leopard", MacOSSnowLeopard.Name())
	assert.Equal(t, "Mac OS X Leopard", MacOSLeopard.Name())
	assert.Equal(t, "Mac OS X Tiger", MacOSTiger.Name())
}

func TestMacOSVersion(t *testing.T) {
	assert.Equal(t, "10.13", MacOSHighSierra.Version())
	assert.Equal(t, "10.12", MacOSSierra.Version())
	assert.Equal(t, "10.11", MacOSElCapitan.Version())
	assert.Equal(t, "10.10", MacOSYosemite.Version())
	assert.Equal(t, "10.9", MacOSMavericks.Version())
	assert.Equal(t, "10.8", MacOSMountainLion.Version())
	assert.Equal(t, "10.7", MacOSLion.Version())
	assert.Equal(t, "10.6", MacOSSnowLeopard.Version())
	assert.Equal(t, "10.5", MacOSLeopard.Version())
	assert.Equal(t, "10.4", MacOSTiger.Version())
}

func TestMacOSString(t *testing.T) {
	assert.Equal(t, "macOS High Sierra (10.13)", MacOSHighSierra.String())
	assert.Equal(t, "macOS Sierra (10.12)", MacOSSierra.String())
	assert.Equal(t, "OS X El Capitan (10.11)", MacOSElCapitan.String())
	assert.Equal(t, "OS X Yosemite (10.10)", MacOSYosemite.String())
	assert.Equal(t, "OS X Mavericks (10.9)", MacOSMavericks.String())
	assert.Equal(t, "OS X Mountain Lion (10.8)", MacOSMountainLion.String())
	assert.Equal(t, "OS X Lion (10.7)", MacOSLion.String())
	assert.Equal(t, "Mac OS X Snow Leopard (10.6)", MacOSSnowLeopard.String())
	assert.Equal(t, "Mac OS X Leopard (10.5)", MacOSLeopard.String())
	assert.Equal(t, "Mac OS X Tiger (10.4)", MacOSTiger.String())
}
