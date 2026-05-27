package main

import (
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
