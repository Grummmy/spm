package main

import (
	"strconv"
	"strings"
	"time"
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
		info.WriteString(" ★")
	}

	info.WriteString("\n     \x1b[38;5;8m")
	info.WriteString(prj.Path)
	info.WriteString("\n\x1b[0m")

	info.WriteString(toFixed(40, "Author: "+orDefault(prj.Author, "<nobody>"), " ", true, " "))

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
