package main

import (
	"time"
)

type Project struct {
	Name        string    `json:"name,omitempty"`
	Version     string    `json:"version,omitempty"`
	Description string    `json:"description,omitempty"`
	Author      string    `json:"author,omitempty"`
	License     string    `json:"license,omitempty"`
	Path        string    `json:"path,omitempty"`
	Lang        string    `json:"language,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastOpened  time.Time `json:"lastOpened,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Favorite    bool      `json:"favorite,omitempty"`
}
