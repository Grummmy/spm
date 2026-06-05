package main

import (
	"os/user"
)

var defaultConfig Config = Config{
	Languages: map[string]Lang{
		"ada":         {Name: "Ada", Ext: []string{"adb", "ads"}, Color: 178},
		"assembly":    {Name: "Assembly", Ext: []string{"asm", "s"}, Color: 244},
		"bash":        {Name: "Bash", Ext: []string{"sh", "bash", "zsh"}, Color: 34},
		"brainfuck":   {Name: "Brainfuck", Ext: []string{"bf", "b"}, Color: 196},
		"c":           {Name: "C", Ext: []string{"c", "h"}, Color: 39},
		"clojure":     {Name: "Clojure", Ext: []string{"clj", "cljs", "cljc"}, Color: 70},
		"cobol":       {Name: "COBOL", Ext: []string{"cob", "cbl"}, Color: 220},
		"cpp":         {Name: "C++", Ext: []string{"cpp", "cc", "cxx", "hpp", "hh", "hxx"}, Color: 45},
		"crystal":     {Name: "Crystal", Ext: []string{"cr"}, Color: 201},
		"csharp":      {Name: "C#", Ext: []string{"cs"}, Color: 99},
		"css":         {Name: "CSS", Ext: []string{"css", "scss", "sass"}, Color: 33},
		"dart":        {Name: "Dart", Ext: []string{"dart"}, Color: 32},
		"delphi":      {Name: "Delphi", Ext: []string{"pas", "dpr"}, Color: 166},
		"elixir":      {Name: "Elixir", Ext: []string{"ex", "exs"}, Color: 129},
		"erlang":      {Name: "Erlang", Ext: []string{"erl", "hrl"}, Color: 160},
		"fortran":     {Name: "Fortran", Ext: []string{"f", "f90", "f95", "f03"}, Color: 214},
		"fsharp":      {Name: "F#", Ext: []string{"fs", "fsi", "fsx"}, Color: 75},
		"go":          {Name: "Go", Ext: []string{"go"}, Color: 44},
		"groovy":      {Name: "Groovy", Ext: []string{"groovy", "gradle"}, Color: 148},
		"haskell":     {Name: "Haskell", Ext: []string{"hs"}, Color: 135},
		"html":        {Name: "HTML", Ext: []string{"html", "htm"}, Color: 208},
		"java":        {Name: "Java", Ext: []string{"java"}, Color: 202},
		"javascript":  {Name: "JavaScript", Ext: []string{"js", "mjs", "cjs"}, Color: 226},
		"json":        {Name: "JSON", Ext: []string{"json", "jsonc"}, Color: 220},
		"julia":       {Name: "Julia", Ext: []string{"jl"}, Color: 141},
		"kotlin":      {Name: "Kotlin", Ext: []string{"kt", "kts"}, Color: 208},
		"lua":         {Name: "Lua", Ext: []string{"lua"}, Color: 63},
		"markdown":    {Name: "Markdown", Ext: []string{"md", "markdown"}, Color: 111},
		"matlab":      {Name: "MATLAB", Ext: []string{"m"}, Color: 166},
		"nim":         {Name: "Nim", Ext: []string{"nim"}, Color: 220},
		"objective-c": {Name: "Objective-C", Ext: []string{"mm"}, Color: 81},
		"ocaml":       {Name: "OCaml", Ext: []string{"ml", "mli"}, Color: 214},
		"perl":        {Name: "Perl", Ext: []string{"pl", "pm"}, Color: 69},
		"php":         {Name: "PHP", Ext: []string{"php", "phtml"}, Color: 99},
		"powershell":  {Name: "PowerShell", Ext: []string{"ps1", "psm1", "psd1"}, Color: 33},
		"prolog":      {Name: "Prolog", Ext: []string{"pro"}, Color: 172},
		"python":      {Name: "Python", Ext: []string{"py", "pyw", "pyi"}, Color: 220},
		"r":           {Name: "R", Ext: []string{"r"}, Color: 68},
		"racket":      {Name: "Racket", Ext: []string{"rkt"}, Color: 124},
		"ruby":        {Name: "Ruby", Ext: []string{"rb"}, Color: 160},
		"rust":        {Name: "Rust", Ext: []string{"rs"}, Color: 208},
		"scala":       {Name: "Scala", Ext: []string{"scala", "sc"}, Color: 196},
		"scheme":      {Name: "Scheme", Ext: []string{"scm", "ss"}, Color: 172},
		"smalltalk":   {Name: "Smalltalk", Ext: []string{"st"}, Color: 214},
		"solidity":    {Name: "Solidity", Ext: []string{"sol"}, Color: 40},
		"sql":         {Name: "SQL", Ext: []string{"sql"}, Color: 27},
		"swift":       {Name: "Swift", Ext: []string{"swift"}, Color: 208},
		"toml":        {Name: "TOML", Ext: []string{"toml"}, Color: 173},
		"typescript":  {Name: "TypeScript", Ext: []string{"ts", "tsx"}, Color: 39},
		"txt":         {Name: "Text", Ext: []string{"txt"}, Color: 250},
		"vala":        {Name: "Vala", Ext: []string{"vala", "vapi"}, Color: 111},
		"verilog":     {Name: "Verilog", Ext: []string{"v"}, Color: 178},
		"vhdl":        {Name: "VHDL", Ext: []string{"vhd", "vhdl"}, Color: 177},
		"xml":         {Name: "XML", Ext: []string{"xml", "xsd", "xsl", "xslt", "svg"}, Color: 214},
		"yaml":        {Name: "YAML", Ext: []string{"yaml", "yml"}, Color: 71},
		"zig":         {Name: "Zig", Ext: []string{"zig"}, Color: 214},
	},
	StatsExcludeDirs: Matcher{
		Full: []string{
			// VCS
			".git", ".hg", ".svn",
			// JS/TS
			"node_modules", "bower_components", "jspm_packages",
			// Rust
			"target",
			// Python
			"__pycache__", "venv", ".venv", "env", "ENV",
			".gradle", // Java / JVM
			// Ruby
			".bundle", "gemfile.lock",
			// .NET
			"packages", "obj", "bin",
			// JS frameworks / build
			".next", ".nuxt", ".svelte-kit", ".parcel-cache",
			// coverage / caches
			"coverage", ".nyc_output", ".cache",
			// IDEs
			".idea", ".vscode",
			// common build outputs
			"dist", "build", "out", "release", "debug",
			// temp
			"tmp", "temp",

			"vendor",
		},
		Start: []string{
			// versioned build dirs
			"build-", "dist-", "target-",
		},
		End: []string{
			// vendored / external code
			"-vendor", "-vendors",
			"-external", "-third-party",
			"-deps",
		},
	},
	StatsExcludeFiles: Matcher{
		Full: []string{
			// JS/TS
			"package-lock.json", "npm-shrinkwrap.json", "yarn.lock",
			"pnpm-lock.yaml", "bun.lockb",
			"composer.lock", // PHP
			"go.sum",        // GO
			// Rust
			"cargo.lock",
			// Python
			"pipfile.lock", "poetry.lock", "uv.lock",
			"requirements.txt",
			// spm conf
			".spm.toml",
		},
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
