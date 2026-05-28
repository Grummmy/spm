package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type Args struct {
	Command      string
	Name         string
	Description  string
	Lang         string
	Tags         []string
	Path         string
	Mkdir        bool
	Sort         string
	Des          bool
	CountLines   bool
	NoExceptions bool
	Reason       string
	Confirm      bool
	Permanent    bool
	Action       string
	New          string
}

var helpMessage string = `Simple Project Manager (spm)

Usage:
	spm <command> [options]

Commands:
	init [name]               Initialize new project.
	list                      List all tracked projects.
	delete <name> [reason]    Delete a project (moves it to trash by default).
	open <name>               Open a project in the default editor.
	info <name>               Show detailed information about a project.
	rename <old> <new>        Rename a project.
	history                   Show project creation and deletion history.

Run 'spm <command> -h' to get detailed help for each command.`

func getArgs() Args {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(2)
	}

	cla := Args{Command: os.Args[1]}
	sub := cla.Command

	switch sub {
	case "init":
		flag.Usage = func() {
			fmt.Println("Initialize new project and create .spm_project.json file.")
			fmt.Println("Usage: spm init [name]")
			flag.PrintDefaults()
		}
		lang := flag.StringP("lang", "l", "", "programming language that project mainly relies on")
		tags := flag.StringSliceP("tag", "t", nil, "project tags")
		path := flag.StringP("path", "p", "", "project directory path")
		mkdir := flag.BoolP("mkdir", "m", false, "whether to create new dir for a project or use currnet")
		flag.Parse()

		args := flag.Args()
		if len(args) >= 2 { // skip first one, because its a subcommand
			cla.Name = args[1]
		}

		cla.Lang, cla.Tags, cla.Path, cla.Mkdir = *lang, *tags, *path, *mkdir

	case "list":
		flag.Usage = func() {
			fmt.Println("List all tracked projects.")
			fmt.Println("Usage: spm list [options]")
			flag.PrintDefaults()
		}
		sort := flag.StringP("sort", "s", "opened", "sort projects by: date (creation date), lang, opened (last opened)")
		des := flag.Bool("des", false, "change sorting order to descending (default ascending)")
		countLines := flag.BoolP("lines", "L", false, "count lang use statistics by lines instead of characters")
		noExceptions := flag.BoolP("no-exceptions", "N", false, "dont exclude non-source code files when counting lang statistics")
		flag.Parse()

		cla.Sort, cla.Des, cla.CountLines, cla.NoExceptions = *sort, *des, *countLines, *noExceptions

	case "delete":
		flag.Usage = func() {
			fmt.Println("Delete a project (moves it to trash by default).")
			fmt.Println("Usage: spm delete <name> [reason] [options]")
			flag.PrintDefaults()
		}
		confirm := flag.BoolP("confirm", "y", false, "skip confirmation prompts, always accept")
		permanent := flag.BoolP("permanent", "P", false, "delete project permanently, instead of compressing and moving to spm trash")
		flag.Parse()

		args := flag.Args()
		if len(args) > 2 { // skip first one, because its a subcommand
			cla.Name = args[1]
			cla.Reason = args[2]
		} else if len(args) == 2 {
			cla.Name = args[1]
		}
		cla.Confirm, cla.Permanent = *confirm, *permanent

	case "history":
		flag.Usage = func() {
			fmt.Println("Show project creation and deletion history.")
			fmt.Println("Usage: spm history [options]")
			flag.PrintDefaults()
		}
		action := flag.StringP("action", "a", "all", "action to be listed: delete, create, all")
		flag.Parse()

		cla.Action = *action

	case "info":
		flag.Usage = func() {
			fmt.Println("Show detailed information about the project.")
			fmt.Println("Usage: spm info <name>")
			flag.PrintDefaults()
		}
		flag.Parse()

		args := flag.Args()
		if len(args) >= 2 { // skip first one, because its a subcommand
			cla.Name = args[1]
		}

	case "open":
		flag.Usage = func() {
			fmt.Println("Open code editor and new terminal with project directory.")
			fmt.Println("Usage: spm open <name>")
			flag.PrintDefaults()
		}
		flag.Parse()

		args := flag.Args()
		if len(args) >= 2 { // skip first one, because its a subcommand
			cla.Name = args[1]
		}

	case "rename":
		flag.Usage = func() {
			fmt.Println("Rename a project.")
			fmt.Println("Usage: spm rename <old> <new>")
			flag.PrintDefaults()
		}
		flag.Parse()

		args := flag.Args()
		if len(args) >= 3 { // skip first one, because its a subcommand
			cla.Name = args[1]
			cla.New = args[2]
		}

	default:
		fmt.Println(helpMessage)
		os.Exit(0)
	}

	return cla
}
