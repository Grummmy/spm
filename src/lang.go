package main

type Lang struct {
	Ext []string `toml:"ext"`
	// fg \x1b[38;5;Nm ; bg \x1b[48;5;Nm
	Color int `toml:"color"`
}
