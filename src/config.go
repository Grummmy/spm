package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

var APPDIR string
var CONFIG Config

type Config struct {
	ProjectsDir     string              `toml:"projects_dir"`
	Languages       map[string][]string `toml:"languages"`
	ProjectDefaults Project             `toml:"project_defaults"`
}

func checkAppDir() {
	if os.Getenv("SPM_RUN_PORTABLE") != "" {
		logger.Debug("Running in portable mode")
		APPDIR = exeDir()
	} else if dir, err := os.UserConfigDir(); err != nil {
		logger.Error("Failed to get user config dir:", err)
		APPDIR = exeDir()
	} else {
		APPDIR = filepath.Join(dir, "spm")
		if err := os.Mkdir(APPDIR, 0755); err != nil && !os.IsExist(err) {
			logger.Error("Failed to create app directory ("+APPDIR+"):", err)
		}
	}

	configPath := filepath.Join(APPDIR, "config.toml")
	if fileExists(configPath) {
		bs, err := os.ReadFile(configPath)
		if err != nil {
			logger.Error("Failed to read config ("+configPath+"):", err)
			os.Exit(1)
		}

		err = toml.Unmarshal(bs, &CONFIG)
		if err != nil {
			logger.Error("Failed to parse config ("+configPath+"):", err)
			os.Exit(1)
		}
	} else {
		logger.Info("No config found, creating new one with default values.")
		fillDefaultConfig(&DEFAULT_CONFIG)
		updateConfig(DEFAULT_CONFIG)

		CONFIG = DEFAULT_CONFIG
	}
}

func updateConfig(new Config) {
	data, err := toml.Marshal(new)
	if err != nil {
		logger.Error("Failed to marshal config:", err)
		return
	}

	configPath := filepath.Join(APPDIR, "config.toml")
	err = os.WriteFile(configPath, data, 0755)
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
