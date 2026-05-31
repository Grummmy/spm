package main

import (
	"os"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
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

func orDefault[T comparable](v T, def T) T {
	var zero T

	if v == zero {
		return def
	}
	return v
}

func orDefaultTime(t time.Time, format string, def string) string {
	if t.IsZero() {
		return def
	}

	return t.Format(format)
}

func toFixed(length int, text string, fill string, strict bool, end string) string {
	d := utf8.RuneCountInString(text) - length
	if d > 0 {
		if strict {
			return text[:length]
		} else if end != "" {
			return text[:length-1] + end
		} else {
			return text
		}
	} else if d < 0 {
		return text + strings.Repeat(fill, d*-1)
	} else {
		return text
	}
}
