package main

import (
	"os/user"
)

var DEFAULT_CONFIG Config = Config{
	Languages: map[string]Lang{
		"ada":         {Ext: []string{"adb", "ads"}, Color: 178},
		"assembly":    {Ext: []string{"asm", "s"}, Color: 244},
		"bash":        {Ext: []string{"sh", "bash", "zsh"}, Color: 34},
		"brainfuck":   {Ext: []string{"bf", "b"}, Color: 196},
		"c":           {Ext: []string{"c", "h"}, Color: 39},
		"clojure":     {Ext: []string{"clj", "cljs", "cljc"}, Color: 70},
		"cobol":       {Ext: []string{"cob", "cbl"}, Color: 220},
		"cpp":         {Ext: []string{"cpp", "cc", "cxx", "hpp", "hh", "hxx"}, Color: 45},
		"crystal":     {Ext: []string{"cr"}, Color: 201},
		"csharp":      {Ext: []string{"cs"}, Color: 99},
		"css":         {Ext: []string{"css", "scss", "sass"}, Color: 33},
		"dart":        {Ext: []string{"dart"}, Color: 32},
		"delphi":      {Ext: []string{"pas", "dpr"}, Color: 166},
		"elixir":      {Ext: []string{"ex", "exs"}, Color: 129},
		"erlang":      {Ext: []string{"erl", "hrl"}, Color: 160},
		"fortran":     {Ext: []string{"f", "f90", "f95", "f03"}, Color: 214},
		"fsharp":      {Ext: []string{"fs", "fsi", "fsx"}, Color: 75},
		"go":          {Ext: []string{"go"}, Color: 44},
		"groovy":      {Ext: []string{"groovy"}, Color: 148},
		"haskell":     {Ext: []string{"hs"}, Color: 135},
		"html":        {Ext: []string{"html", "htm"}, Color: 208},
		"java":        {Ext: []string{"java", "class", "jar"}, Color: 202},
		"javascript":  {Ext: []string{"js", "mjs", "cjs"}, Color: 226},
		"julia":       {Ext: []string{"jl"}, Color: 141},
		"kotlin":      {Ext: []string{"kt", "kts"}, Color: 208},
		"lua":         {Ext: []string{"lua"}, Color: 63},
		"matlab":      {Ext: []string{"m"}, Color: 166},
		"nim":         {Ext: []string{"nim"}, Color: 220},
		"objective-c": {Ext: []string{"m", "mm"}, Color: 81},
		"ocaml":       {Ext: []string{"ml", "mli"}, Color: 214},
		"perl":        {Ext: []string{"pl", "pm"}, Color: 69},
		"php":         {Ext: []string{"php", "phtml"}, Color: 99},
		"powershell":  {Ext: []string{"ps1", "psm1"}, Color: 33},
		"prolog":      {Ext: []string{"pl"}, Color: 172},
		"python":      {Ext: []string{"py", "pyw"}, Color: 220},
		"r":           {Ext: []string{"r"}, Color: 68},
		"racket":      {Ext: []string{"rkt"}, Color: 124},
		"ruby":        {Ext: []string{"rb"}, Color: 160},
		"rust":        {Ext: []string{"rs"}, Color: 208},
		"scala":       {Ext: []string{"scala", "sc"}, Color: 196},
		"scheme":      {Ext: []string{"scm", "ss"}, Color: 172},
		"shell":       {Ext: []string{"sh"}, Color: 40},
		"smalltalk":   {Ext: []string{"st"}, Color: 214},
		"solidity":    {Ext: []string{"sol"}, Color: 40},
		"sql":         {Ext: []string{"sql"}, Color: 27},
		"swift":       {Ext: []string{"swift"}, Color: 208},
		"typescript":  {Ext: []string{"ts", "tsx"}, Color: 39},
		"vala":        {Ext: []string{"vala", "vapi"}, Color: 111},
		"verilog":     {Ext: []string{"v"}, Color: 178},
		"vhdl":        {Ext: []string{"vhd", "vhdl"}, Color: 177},
		"zig":         {Ext: []string{"zig"}, Color: 214},
	},
	ProjectDefaults: Project{
		License: "GPL v3.0",
		Version: "0.1.0-alpha",
	},
}

func fillDefaultConfig(cfg *Config) {
	cfg.ProjectsDir = guessProjectsDir()
	if cfg.ProjectsDir != "" {
		logger.Info("Guessed project directory to be '" + cfg.ProjectsDir + "'")
	}

	if usr, err := user.Current(); err == nil {
		cfg.ProjectDefaults.Author = usr.Username
		logger.Info("'" + cfg.ProjectDefaults.Author + "' was chosen as default projects' author")
	} else {
		logger.Error("Failed to get user's username:", err)
	}
}
