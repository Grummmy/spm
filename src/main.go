package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/peterh/liner"
)

var loglevel string = "info"

func init() {
	logger.logLevel = logLevelFromString(loglevel)

	checkAppDir()
}

func main() {
	cla := getArgs()

	switch cla.Command {
	case "new":
		newProject(cla)
	case "init":
		initProject(cla)
	}

	updateConfig(CONFIG)
}

func newProject(cla Args) {
	if cla.Name == "" { // exit if no name was passed
		fmt.Println("Not enough arguments. Name is required to use 'spm new <name> [description]'")
	}
	prj := Project{ // basic project configuration
		Name:        cla.Name,
		Description: cla.Description,
		Lang:        cla.Lang,
		Created:     time.Now(),
		Tags:        cla.Tags,
	}

	if CONFIG.ProjectsDir != "" { // try to create project folder if ProjectsDir is set
		prj.Path = filepath.Join(CONFIG.ProjectsDir, prj.Name)
		if err := os.Mkdir(APPDIR, 0755); err != nil && !os.IsExist(err) {
			logger.Error("Failed to create directory for '"+prj.Name+"':", err)
		}
	} else {
		logger.Warn("Could not create directory for '" + prj.Name + "', because projects_dir is not defined in config")
	}

	if CONFIG.ProjectDefaults.Author != "" { // try applying default author
		prj.Author = CONFIG.ProjectDefaults.Author
	}
	if CONFIG.ProjectDefaults.License != "" { // try applying default license
		prj.License = CONFIG.ProjectDefaults.License
	}
	if CONFIG.ProjectDefaults.Version != "" { // try applying default version
		prj.Version = CONFIG.ProjectDefaults.Version
	}

	CONFIG.Projects = append(CONFIG.Projects, prj)
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
	fmt.Println("This command will walk you through creating .spm_project.toml file, file describing your project.\nYour project also will be added to main config.toml file.\n\nYou can declare any default values for that utility in config file, in \"project_defaults\".\nIt supports all fields that project has, but only asked values will be applied.\nSome default values are handled differently:\n\n * Tags: default tags will be added to the ones you type when asked.\n\nInstructions:\n - Default values are shown in (parentheses).\n - To add multiple tags, separate them by space.\n - Press ^C to leave value empty, default value won't be used\n - Press ^D to exit early, changes won't be applied\n ")

	// ---- get project name ------------->>
	for {
		if cla.Name != "" { // if name was passed as argument when calling 'spm init'
			fmt.Print("Project name (" + cla.Name + "): ")
			def.Name = cla.Name
		} else if cwd != "" { // if no name was passed, but cwd is available
			fmt.Print("Project name (" + filepath.Base(cwd) + "): ")
			def.Name = filepath.Base(cwd)
		} else if def.Name != "" { // if cwd isnt available, but default name is set in config
			fmt.Print("Project name (" + def.Name + "): ")
		} else { // if nothing above worked, just ask the name
			fmt.Print("Project name: ")
		}
		scan, err := (line.Prompt(""))
		if err != nil && err != liner.ErrPromptAborted {
			fmt.Println("\nAborting.")
			return
		}
		scan = strings.TrimSpace(scan)

		if scan != "" && testFilename(scan) { // typed name has priority over defautl and passed in args
			prj.Name = scan
			break
		} else if def.Name != "" && err != liner.ErrPromptAborted && testFilename(scan) { // if default value is set, but not ^C pressed - name cant be empty
			prj.Name = def.Name
			break
		} else {
			fmt.Println(" - Project name can't be empty and should be path-safe (available as filename)")
		}
	}
	// <<-- get project name ---------------

	// ---- get version ------------------>>
	if def.Version != "" { // if default version is set
		fmt.Print("Version (" + def.Version + "): ")
	} else {
		fmt.Print("Version: ")
	}
	scan, err := line.Prompt("")
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
		fmt.Print("Description (default desc.): ")
	} else {
		fmt.Print("Description: ")
	}
	scan, err = line.Prompt("")
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
		fmt.Print("Author (" + def.Author + "): ")
	} else {
		fmt.Print("Author: ")
	}
	scan, err = line.Prompt("")
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
		fmt.Print("License (" + def.License + "): ")
	} else {
		fmt.Print("License: ")
	}
	scan, err = line.Prompt("")
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
		fmt.Print("Programming Language(" + def.Lang + "): ")
	} else {
		fmt.Print("Programming language: ")
	}
	scan, err = line.Prompt("")
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
		fmt.Print("Tags (" + strings.Join(def.Tags, " ") + "): ")
	} else {
		fmt.Print("Tags: ")
	}
	scan, err = line.Prompt("")
	if err != nil && err != liner.ErrPromptAborted {
		fmt.Println("\nAborting.")
		return
	}
	scan = strings.TrimSpace(scan)

	if err != liner.ErrPromptAborted { // leave empty if ^C pressed
		if scan != "" { // default tags are added to typed tags
			prj.Tags = append(strings.Split(scan, " "), def.Tags...)
		} else { // if no typed tags were added, default are used
			prj.Tags = def.Tags
		}
	}
	// <<-- get tags -----------------------

	fmt.Println("\nAbout to write to " + filepath.Join(cwd, ".spm.toml") + ":")
	data, err := toml.Marshal(prj)
	if err != nil {
		logger.Error("Failed to marshal config:", err)
	} else {
		fmt.Println(string(data))
	}

	scan, _ = line.Prompt("\nIs this OK? (yes) ")
	scan = strings.TrimSpace(strings.ToLower(scan)) // trim + toLower the input
	if scan == "" || strings.HasPrefix(scan, "y") || strings.HasPrefix(scan, "+") {
		prj.Created = time.Now()
		CONFIG.Projects = append(CONFIG.Projects, prj)
	} else {
		fmt.Println("Project wasn't added.")
	}
}
