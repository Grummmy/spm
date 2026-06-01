package main

import (
	"maps"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Project struct {
	Name        string    `toml:"name,omitempty"`
	Version     string    `toml:"version,omitempty"`
	Description string    `toml:"description,omitempty"`
	Author      string    `toml:"author,omitempty"`
	License     string    `toml:"license,omitempty"`
	Path        string    `toml:"path,omitempty"`
	Lang        string    `toml:"language,omitempty"`
	Created     time.Time `toml:"created,omitempty"`
	LastOpened  time.Time `toml:"lastOpened,omitempty"`
	Tags        []string  `toml:"tags,omitempty"`
	Favorite    bool      `toml:"favorite,omitempty"`
}

func formatProject(prj Project) string {
	var info strings.Builder

	if prj.Lang != "" {
		info.WriteString(" [\x1b[38;5;")
		info.WriteString(strconv.Itoa(CONFIG.Languages[prj.Lang].Color))
		info.WriteString("m")
		info.WriteString(prj.Lang)
		info.WriteString("\x1b[0m] ")
	}

	info.WriteString(prj.Name)
	if prj.Version != "" {
		info.WriteString(" (\x1b[38;5;3m")
		info.WriteString(prj.Version)
		info.WriteString("\x1b[0m)")
	}
	if prj.Favorite {
		info.WriteString(" \x1b[38;5;5m\x1b[1m⬤ fav\x1b[0m")
	}

	info.WriteString("\n     \x1b[38;5;8m")
	info.WriteString(prj.Path)
	info.WriteString("\n\x1b[0m")

	info.WriteString(toFixed(40, "Author: \x1b[38;5;2m\x1b[1m"+orDefault(prj.Author, "<nobody>"), " ", true, " "))
	info.WriteString("\x1b[0m")

	info.WriteString("License: ")
	info.WriteString(orDefault(prj.License, "<none>"))

	if prj.Description != "" {
		info.WriteString("\n\n")
		info.WriteString(prj.Description)
	}

	info.WriteString("\n\n")

	info.WriteString("Tags: ")
	info.WriteString(strings.Join(prj.Tags, "\x1b[38;5;8m,\x1b[0m "))
	info.WriteString("\n")

	info.WriteString(toFixed(40, "Created at \x1b[38;5;6m"+orDefaultTime(prj.Created, time.DateTime, "<unknown>"), " ", false, " "))
	info.WriteString(toFixed(40, "\x1b[0mLast opened at \x1b[38;5;6m"+orDefaultTime(prj.LastOpened, time.DateTime, "<unknown>"), " ", false, " "))
	info.WriteString("\x1b[0m")

	return info.String()
}

// eDirs - excluded dirs; eFiles - excluded eFiles
// separated for hight checking speed
func getLangStats(path string, eDirs Matcher, eFiles Matcher, result map[string]int, lines bool) map[string]int {
	if result == nil {
		result = make(map[string]int)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		logger.Error("Could not read dir '"+path+"':", err)
		return result
	}

	for _, f := range files {
		name := strings.ToLower(f.Name())
		var ext string

		if !f.IsDir() {
			ext = name[strings.LastIndex(name, ".")+1:]
			if !slices.Contains(slices.Collect(maps.Keys(CONFIG.Extentions)), ext) {
				continue
			}

			result[ext] += countFile(filepath.Join(path, name), lines)
			continue
		}
		result = getLangStats(filepath.Join(path, name), eDirs, eFiles, result, lines)
	}

	return result
}

func countFile(path string, lines bool) int {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error("Could not open file '"+path+"':", err)
		return 0
	}

	if lines {
		return strings.Count(string(data), "\n")
	}
	return utf8.RuneCount(data)
}

func langStatLine(stats map[string]int, length int) string {
	total := 0
	for _, n := range stats {
		total += n
	}

	var line strings.Builder
	for lang, amount := range stats {
		n := float64(float32(amount) / float32(total) * float32(length))
		logger.Debug(lang, amount, n)

		line.WriteString("\x1b[48;5;")
		line.WriteString(strconv.Itoa(CONFIG.Languages[lang].Color))
		line.WriteString("m")
		line.WriteString(strings.Repeat(" ", int(math.Round(n))))
	}

	line.WriteString("\x1b[0m")
	return line.String()
}
