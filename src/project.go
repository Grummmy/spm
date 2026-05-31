package main

import (
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
		info.WriteString(" [")
		info.WriteString(prj.Lang)
		info.WriteString("] ")
	}

	info.WriteString(prj.Name)
	if prj.Version != "" {
		info.WriteString(" (")
		info.WriteString(prj.Version)
		info.WriteString(")")
	}
	if prj.Favorite {
		info.WriteString(" ★")
	}

	info.WriteString("\n     ")
	info.WriteString(prj.Path)
	info.WriteString("\n")

	info.WriteString(toFixed(40, "Author: "+orDefault(prj.Author, "<nobody>"), " ", true, " "))

	info.WriteString("License: ")
	info.WriteString(orDefault(prj.License, "<none>"))

	if prj.Description != "" {
		info.WriteString("\n\n")
		info.WriteString(prj.Description)
	}

	info.WriteString("\n\n")

	info.WriteString("Tags: ")
	info.WriteString(strings.Join(prj.Tags, ", "))
	info.WriteString("\n")

	info.WriteString(toFixed(40, "Created at "+orDefaultTime(prj.Created, time.DateTime, "<unknown>"), " ", false, " "))
	info.WriteString(toFixed(40, "Last opened at "+orDefaultTime(prj.LastOpened, time.DateTime, "<unknown>"), " ", false, " "))

	return info.String()
}
