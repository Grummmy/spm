package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/peterh/liner"
)

var loglevel string = "debug"

func init() {
	logger.logLevel = logLevelFromString(loglevel)

	checkAppDir()
}

func main() {
	cla := getArgs()

	switch cla.Command {
	case "init":
		initProject(cla)
	}

	updateConfig(CONFIG)
}

func initProject(cla Args) {
	var prj Project
	def := CONFIG.ProjectDefaults

	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	cwd, err := os.Getwd()
	if err != nil {
		logger.Error("Failed to get current working directory:", err)
	}

	// <->--- info ------------------------<->
	fmt.Println("This command will walk you through creating .spm.toml file, file describing your project.\n\nYou can declare any default values for that utility in config file, under \"project_defaults\" property.\nIt supports all fields that project has, however only asked values will be applied.\nSome default values are handled differently:\n\n * Tags: default tags will be added to the ones you type when asked.\n\nInstructions:\n - Default values are shown in (parentheses).\n - To add multiple tags, separate them by space.\n - Add - (minus) to delete any previously added tag (or default)\n - Press ^C to leave value empty, default value won't be used\n - Press ^D to exit early, changes won't be applied\n ")
	prompt := ""

	// ---- get project name ------------->>
	for {
		if cla.Name != "" { // if name was passed as argument when calling 'spm init'
			prompt = "Project name (" + cla.Name + "): "
			def.Name = cla.Name
		} else if cwd != "" { // if no name was passed, but cwd is available
			prompt = "Project name (" + filepath.Base(cwd) + "): "
			def.Name = filepath.Base(cwd)
		} else if def.Name != "" { // if cwd isnt available, but default name is set in config
			prompt = "Project name (" + def.Name + "): "
		} else { // if nothing above worked, just ask the name
			prompt = "Project name: "
		}
		scan, err := line.Prompt(prompt)
		if err != nil && err != liner.ErrPromptAborted {
			fmt.Println("\nAborting.")
			return
		}
		scan = strings.TrimSpace(scan)

		if scan != "" && testFilename(scan) { // typed name has priority over defautl and passed in args
			prj.Name = scan
			break
		} else if def.Name != "" && err != liner.ErrPromptAborted && testFilename(def.Name) { // if default value is set, but not ^C pressed - name cant be empty
			prj.Name = def.Name
			break
		} else {
			fmt.Println(" - Project name can't be empty and should be path-safe (available as filename)")
		}
	}
	// <<-- get project name ---------------

	// ---- get version ------------------>>
	if def.Version != "" { // if default version is set
		prompt = "Version (" + def.Version + "): "
	} else {
		prompt = "Version: "
	}
	scan, err := line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // typed version has priority over default one
			prj.Version = scan
		} else if def.Version != "" {
			prj.Version = def.Version
		}
	}
	// <<-- get version --------------------

	// ---- get project description ------>>
	if def.Description != "" { // if default description is set in config
		prompt = "Description (default desc.): "
	} else {
		prompt = "Description: "
	}
	scan, err = line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // typed description has priority over default
			prj.Description = scan
		} else {
			prj.Description = def.Description
		}
	}
	// <<-- get project description --------

	// ---- get project author ----------->>
	if def.Author != "" { // if default author is set
		prompt = "Author (" + def.Author + "): "
	} else {
		prompt = "Author: "
	}
	scan, err = line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // typed author has priority over default
			prj.Author = scan
		} else if def.Author != "" {
			prj.Author = def.Author
		}
	}
	// <<-- get project author -------------

	// ---- get project license ---------->>
	if def.License != "" { // if default license is set
		prompt = "License (" + def.License + "): "
	} else {
		prompt = "License: "
	}
	scan, err = line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // typed license has priority over default
			prj.License = scan
		} else if def.License != "" {
			prj.License = def.License
		}
	}
	// <<-- get project license ------------

	// ---- get project lang ------------->>
	if def.Lang != "" { // if default lang is set in config
		prompt = "Programming Language(" + def.Lang + "): "
	} else {
		prompt = "Programming language: "
	}
	scan, err = line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // typed lang has priority over default
			prj.Lang = scan
		} else if def.Lang != "" {
			prj.Lang = def.Lang
		}
	}
	// <<-- get project lang ---------------

	// ---- get tags --------------------->>
	if len(def.Tags) != 0 { // if there are any tags in defaults
		prompt = "Tags (" + strings.Join(def.Tags, " ") + "): "
	} else {
		prompt = "Tags: "
	}
	scan, err = line.Prompt(prompt)
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		tags := append(strings.Fields(scan), def.Tags...)

		for i := range tags { // iterate through each tag to apply "-" rules
			tag := strings.TrimSpace(tags[i])
			index := slices.Index(prj.Tags, tag[1:])

			// if tag starts with -, and it was added to tags earlier, than remove it
			if strings.HasPrefix(tag, "-") && index != -1 {
				prj.Tags = append(prj.Tags[:index], prj.Tags[index+1:]...)
			} else if !strings.HasPrefix(tag, "-") {
				prj.Tags = append(prj.Tags, tag)
			}
		}
		logger.Debug("Tags:", strings.Join(prj.Tags, " "))
	}
	// <<-- get tags -----------------------

	fmt.Println("\nAbout to write to " + filepath.Join(cwd, ".spm.toml") + ":")
	data, err := toml.Marshal(prj)
	if err != nil {
		logger.Error("Failed to marshal project:", err)
	} else {
		fmt.Println(string(data) + "\n")
	}

	scan, err = line.Prompt("Is this OK? (yes) ")
	scan = strings.TrimSpace(strings.ToLower(scan)) // trim + toLower the input
	if err == nil && (scan == "" || strings.HasPrefix(scan, "y") || strings.HasPrefix(scan, "+")) {
		prj.Created = time.Now()

		prj.Path, err = os.Getwd()
		if err != nil { // get wd to know where the project is
			logger.Error("Could not get wd:", err)
			return
		}

		if cla.Mkdir { // if mkdir, then add project name to the path
			prj.Path = path.Join(prj.Path, prj.Name)
		}

		// create dir if doesnt exist
		if err := os.MkdirAll(prj.Path, 0755); err != nil {
			logger.Error("Could not create directory '"+prj.Path+"':", err)
			return
		}

		if err := os.WriteFile(path.Join(prj.Path, ".spm.toml"), data, 0655); err != nil {
			logger.Error("Could not create file '"+path.Join(prj.Path, ".spm.toml")+"':", err)
		} else {
			logger.Debug("Successfully wrote '" + path.Join(prj.Path, ".spm.toml") + "'")
		}
	} else {
		fmt.Println("Project wasn't created.")
	}
}
