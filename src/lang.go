package main

type Lang struct {
	Name string   `toml:"name,omitempty"`
	Ext  []string `toml:"ext"`
	// fg \x1b[38;5;Nm ; bg \x1b[48;5;Nm
	Color int `toml:"color"`
}
