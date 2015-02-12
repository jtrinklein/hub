package utils

import (
	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"
	"testing"
)

func TestSearchBrowserLauncher(t *testing.T) {
	browser := searchBrowserLauncher("darwin")
	assert.Equal(t, "open", browser)

	browser = searchBrowserLauncher("windows")
	assert.Equal(t, "cmd /c start", browser)
}

func TestConcatPaths(t *testing.T) {
	assert.Equal(t, "foo/bar/baz", ConcatPaths("foo", "bar", "baz"))
}

func TestNewlineForOsIsCorrect(t *testing.T) {
	newline := getNewlineForOs("windows")

	assert.Equal(t, newline, "\r\n")

	newline = getNewlineForOs("darwin")

	assert.Equal(t, newline, "\n")
}
