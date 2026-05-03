package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var APPDIR string
var CONFIG Config

type Config struct {
	ProjectsDir     string              `json:"projects_dir"`
	Projects        []Project           `json:"projects"`
	Languages       map[string][]string `json:"languages"`
	ProjectDefaults Project             `json:"project_defaults"`
}

func checkAppDir() {
	if os.Getenv("SPM_RUN_PORTABLE") != "" {
		APPDIR = exeDir()
	} else if dir, err := os.UserConfigDir(); err != nil {
		logger.Error("Failed to get user config dir:", err)
		APPDIR = exeDir()
	} else {
		APPDIR = filepath.Join(dir, "SiPM")
		if err := os.Mkdir(APPDIR, 0755); err != nil && !os.IsExist(err) {
			logger.Error("Failed to create app directory ("+APPDIR+"):", err)
		}
	}

	configPath := filepath.Join(APPDIR, "config.json")
	if fileExists(configPath) {
		bs, err := os.ReadFile(configPath)
		if err != nil {
			logger.Error("Failed to read config ("+configPath+"):", err)
			os.Exit(1)
		}

		err = json.Unmarshal(bs, &CONFIG)
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
	data, err := json.MarshalIndent(new, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal config:", err)
		return
	}

	configPath := filepath.Join(APPDIR, "config.json")
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
