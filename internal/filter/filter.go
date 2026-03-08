package filter

import (
	"path/filepath"
	"strings"
)

var ignoredDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"bin":          true,
	"build":        true,
	"dist":         true,
	".hotreload":   true,
}

var ignoredExtensions = map[string]bool{
	".tmp": true,
	".swp": true,
	".swo": true,
	".log": true,
}

func ShouldIgnore(path string) bool {

	base := filepath.Base(path)

	if ignoredDirs[base] {
		return true
	}

	ext := filepath.Ext(base)

	if ignoredExtensions[ext] {
		return true
	}

	if strings.HasPrefix(base, ".#") {
		return true
	}

	if strings.HasSuffix(base, "~") {
		return true
	}

	return false
}
