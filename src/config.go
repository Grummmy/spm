package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	APPDIR string
	CONFIG Config
)

type Config struct {
	Path              string            `toml:"-"`
	ProjectsDir       string            `toml:"projects_dir"`
	Languages         map[string]Lang   `toml:"languages,omitempty"`
	Extentions        map[string]string `toml:"-"`
	StatsExcludeDirs  Matcher           `toml:"stats_exclude_dirs,omitempty"`
	StatsExcludeFiles Matcher           `toml:"stats_exclude_files,omitempty"`
	ProjectDefaults   Project           `toml:"project_defaults"`
}

func checkAppDir() {
	if os.Getenv("SPM_RUN_PORTABLE") != "" {
		logger.Info("Running in portable mode")
		APPDIR = exeDir()
	} else if dir, err := os.UserConfigDir(); err != nil {
		logger.Error("Failed to get user config dir:", err)
		APPDIR = exeDir()
	} else {
		APPDIR = filepath.Join(dir, "spm")
		if err := os.Mkdir(APPDIR, 0o755); err != nil && !os.IsExist(err) {
			logger.Error("Failed to create app directory ("+APPDIR+"):", err)
		}
	}

	CONFIG = loadConfig(filepath.Join(APPDIR, "config.toml"))
}

func loadConfig(path string) Config {
	var cfg Config
	if fileExists(path) {
		bs, err := os.ReadFile(path)
		if err != nil {
			logger.Error("Failed to read config ("+path+"):", err)
			os.Exit(1)
		}

		err = toml.Unmarshal(bs, &cfg)
		if err != nil {
			logger.Error("Failed to parse config ("+path+"):", err)
			os.Exit(1)
		}
	} else {
		logger.Info("No config found, creating new one with default values.")
		fillDefaultConfig(&defaultConfig)
		updateConfig(defaultConfig)

		cfg = defaultConfig
	}

	if cfg.Extentions == nil {
		cfg.Extentions = make(map[string]string)
	}
	for name, lang := range cfg.Languages {
		for _, ext := range lang.Ext {
			cfg.Extentions[ext] = name
		}
	}

	cfg.Path = path
	return cfg
}

func updateConfig(new Config) {
	data, err := toml.Marshal(new)
	if err != nil {
		logger.Error("Failed to marshal config:", err)
		return
	}

	configPath := filepath.Join(APPDIR, "config.toml")
	err = os.WriteFile(configPath, data, 0o755)
	if err != nil {
		logger.Error("Failed to write config ("+configPath+"):", err)
	}
}

func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		logger.Error("Failed to get executable path:", err)
		os.Exit(1)
	}

	return filepath.Dir(exe)
}

func guessProjectsDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		logger.Warn("Could not get home dir to guess projects dir (canceling):", err)
		return ""
	}

	files, err := os.ReadDir(home)
	if err != nil {
		logger.Warn("Could not read home dir (cancleing project dir guess):", err)
		return ""
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		if n := strings.ToLower(f.Name()); n == "projects" || n == "prjs" || n == "project" || n == "prj" {
			return filepath.Join(home, f.Name())
		}
	}

	logger.Warn("Could not guess projects dir, please set it manually in config.")
	return ""
}
