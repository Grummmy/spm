package main

import (
	"os"
	"runtime"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func testFilename(name string) bool {
	if name == "" {
		return false
	}

	switch runtime.GOOS {
	case "windows":
		return !strings.ContainsAny(name, `<>:"/\|?*`)
	default: // unix-like
		return !strings.Contains(name, "/") && !strings.ContainsRune(name, 0)
	}
}
