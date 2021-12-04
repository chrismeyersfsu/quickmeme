package main

import (
	"path/filepath"
)

func GetGifPaths(dir string) []string {
	files, err := filepath.Glob(filepath.Join(dir, "*.gif"))
	panicIf(err)
	return files
}
