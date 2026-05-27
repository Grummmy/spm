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
	Sort         string
	Des          bool
	CountLines   bool
	NoExceptions bool
	Reason       string
	Confirm      bool
	Permanent    bool
	Type         string
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

func _getArgs() Args {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	cla := Args{Command: os.Args[1]}

	fs := flag.NewFlagSet("spm", flag.ExitOnError)
	lang := fs.StringP("lang", "l", "", "programming language that project mainly relies on")
	tags := fs.StringSliceP("tag", "t", nil, "project tags")
	path := fs.StringP("path", "p", "", "project directory path")
	confirm := fs.BoolP("confirm", "y", false, "skip confirmation prompts, always accept")
	permanent := fs.BoolP("permanent", "P", false, "delete project permanently, instead of compressing and moving to spm trash")
	sort := fs.StringP("sort", "s", "opened", "sort projects by: date (creation date), lang, opened (last opened)")
	des := fs.Bool("des", false, "change sorting order to descending (default asceding)")
	countLines := fs.BoolP("lines", "L", false, "count lang use statistics by lines instead of characters")
	noExceptions := fs.BoolP("no-exceptions", "N", false, "dont exclude non source code files when counting lang statistics")
	type_ := fs.StringP("type", "T", "all", "log type to be listed: delete, create, all")

	fs.Parse(os.Args[2:])

	cla.Lang, cla.Tags, cla.Path, cla.Confirm = *lang, *tags, *path, *confirm
	cla.Permanent, cla.Sort, cla.Des = *permanent, *sort, *des
	cla.CountLines, cla.NoExceptions, cla.Type = *countLines, *noExceptions, *type_

	args := fs.Args()
	if len(args) > 1 {
		cla.Name = args[0]
		cla.Description = args[1]
	} else if len(args) == 1 {
		cla.Name = args[0]
	}

	return cla
}

func getArgs() Args {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	cla := Args{Command: os.Args[1]}
	sub := cla.Command

	switch sub {
	case "init":
		fs := flag.NewFlagSet("new", flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Initialize new project and create .spm_project.json file.")
			fmt.Println("Usage: spm init [name]")
			fs.PrintDefaults()
		}
		lang := fs.StringP("lang", "l", "", "programming language that project mainly relies on")
		tags := fs.StringSliceP("tag", "t", nil, "project tags")
		path := fs.StringP("path", "p", "", "project directory path")
		fs.Parse(os.Args[2:])

		args := fs.Args()
		if len(args) >= 1 {
			cla.Name = args[0]
		}

		cla.Lang, cla.Tags, cla.Path = *lang, *tags, *path

	case "list":
		fs := flag.NewFlagSet("list", flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("List all tracked projects.")
			fmt.Println("Usage: spm list [options]")
			fs.PrintDefaults()
		}
		sort := fs.StringP("sort", "s", "opened", "sort projects by: date (creation date), lang, opened (last opened)")
		des := fs.Bool("des", false, "change sorting order to descending (default ascending)")
		countLines := fs.BoolP("lines", "L", false, "count lang use statistics by lines instead of characters")
		noExceptions := fs.BoolP("no-exceptions", "N", false, "dont exclude non-source code files when counting lang statistics")
		fs.Parse(os.Args[2:])

		cla.Sort, cla.Des, cla.CountLines, cla.NoExceptions = *sort, *des, *countLines, *noExceptions

	case "delete":
		fs := flag.NewFlagSet("delete", flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Delete a project (moves it to trash by default).")
			fmt.Println("Usage: spm delete <name> [reason] [options]")
			fs.PrintDefaults()
		}
		confirm := fs.BoolP("confirm", "y", false, "skip confirmation prompts, always accept")
		permanent := fs.BoolP("permanent", "P", false, "delete project permanently, instead of compressing and moving to spm trash")
		fs.Parse(os.Args[2:])

		args := fs.Args()
		if len(args) > 1 {
			cla.Name = args[0]
			cla.Reason = args[1]
		} else if len(args) == 1 {
			cla.Name = args[0]
		}
		cla.Confirm, cla.Permanent = *confirm, *permanent

	case "history":
		fs := flag.NewFlagSet("history", flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Show project creation and deletion history.")
			fmt.Println("Usage: spm history [options]")
			fs.PrintDefaults()
		}
		type_ := fs.StringP("type", "T", "all", "log type to be listed: delete, create, all")
		fs.Parse(os.Args[2:])

		cla.Type = *type_

	case "info":
		fs := flag.NewFlagSet(sub, flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Show detailed information about the project.")
			fmt.Println("Usage: spm info <name>")
			fs.PrintDefaults()
		}
		fs.Parse(os.Args[2:])

		args := fs.Args()
		if len(args) >= 1 {
			cla.Name = args[0]
		}

	case "open":
		fs := flag.NewFlagSet(sub, flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Open code editor and new terminal with project directory.")
			fmt.Println("Usage: spm open <name>")
			fs.PrintDefaults()
		}
		fs.Parse(os.Args[2:])

		args := fs.Args()
		if len(args) >= 1 {
			cla.Name = args[0]
		}

	case "rename":
		fs := flag.NewFlagSet("rename", flag.ExitOnError)
		fs.Usage = func() {
			fmt.Println("Rename a project.")
			fmt.Println("Usage: spm rename <old> <new>")
			fs.PrintDefaults()
		}
		fs.Parse(os.Args[2:])

		args := fs.Args()
		if len(args) >= 2 {
			cla.Name = args[0]
			cla.New = args[1]
		}

	default:
		fmt.Println(helpMessage)
		os.Exit(0)
	}

	return cla
}
