package main

import (
	"os/user"
)

var DEFAULT_CONFIG Config = Config{
	Languages: map[string][]string{
		"ada":         {"adb", "ads"},
		"assembly":    {"asm", "s"},
		"bash":        {"sh", "bash", "zsh"},
		"brainfuck":   {"bf", "b"},
		"c":           {"c", "h"},
		"clojure":     {"clj", "cljs", "cljc"},
		"cobol":       {"cob", "cbl"},
		"cpp":         {"cpp", "cc", "cxx", "hpp", "hh", "hxx"},
		"crystal":     {"cr"},
		"csharp":      {"cs"},
		"css":         {"css", "scss", "sass"},
		"dart":        {"dart"},
		"delphi":      {"pas", "dpr"},
		"elixir":      {"ex", "exs"},
		"erlang":      {"erl", "hrl"},
		"fortran":     {"f", "f90", "f95", "f03"},
		"fsharp":      {"fs", "fsi", "fsx"},
		"go":          {"go", "mod", "sum"},
		"groovy":      {"groovy"},
		"haskell":     {"hs"},
		"html":        {"html", "htm"},
		"java":        {"java", "class", "jar"},
		"javascript":  {"js", "mjs", "cjs"},
		"julia":       {"jl"},
		"kotlin":      {"kt", "kts"},
		"lua":         {"lua"},
		"matlab":      {"m"},
		"nim":         {"nim"},
		"objective-c": {"m", "mm"},
		"ocaml":       {"ml", "mli"},
		"perl":        {"pl", "pm"},
		"php":         {"php", "phtml"},
		"powershell":  {"ps1", "psm1"},
		"prolog":      {"pl"},
		"python":      {"py", "pyw"},
		"r":           {"r"},
		"racket":      {"rkt"},
		"ruby":        {"rb"},
		"rust":        {"rs"},
		"scala":       {"scala", "sc"},
		"scheme":      {"scm", "ss"},
		"shell":       {"sh"},
		"smalltalk":   {"st"},
		"solidity":    {"sol"},
		"sql":         {"sql"},
		"swift":       {"swift"},
		"typescript":  {"ts", "tsx"},
		"vala":        {"vala", "vapi"},
		"verilog":     {"v"},
		"vhdl":        {"vhd", "vhdl"},
		"zig":         {"zig"},
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
