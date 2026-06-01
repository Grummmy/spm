# SPM -- Simple Project Manager
⚠️ Notice: The project is in early alpha, check TODO for already implemented features

I always struggled with convenient project names and project placement. If I want to develop a minecraft mod or plugin, than I should uniquely name it or create a whole new directory tree: `minecraft/mods/mymod`, since I do many minecraft-related things. So, I came up with a solution - Simple Project Manager. It makes creating, viewing, and showing of project stats very convenient (or should at some point).

## Features
- Easily create projects with `spm init` (inspired by `npm init`)
- List projects and some stats, like programming lang. usage
- Fully delete or archive projects with one command
- Look at your work history: project creation, deletion, opening
- Show detailed info about project like modification or last opening time
- Open projects in your default code editor and change your terminal wd
- Configurable project defaults
- Configurable scripts! (inspired by `npm run`)
- Basically, cross platform
- Other ideas accepted!

## Usage
When you first run the executable, it will try to find your default projects directory. It looks for paths like `~/Projects` (any case: `projects`, `PrOjEcTs` etc). If project dir wasnt guessed, the program will prompt you to choose it. Then, spm creates default config. It includes your user name as default author, default license (`GPL v3.0`), initial version (`0.1.0-alpha`), and a bunch of default language extentions mapped to language names and colors.

On Linux, config dir is `$XDG_CONFIG_HOME/spm/` (most likely `~/.config/spm/config.toml`). On Darwin (MacOs) it is `$HOME/Library/Application Support/spm/`, and for Windows it is `%AppData%/spm/`. If `SPM_RUN_PORTABLE` environment variable is set (anithing except ` `, not blank), than executable dir is used as config directory.

#### Create project
```bash
spm init
```

You will be prompted several questions like project name, description, tags etc for your new porject. At the end, it will print settings it collect and ask if these are okay. If you proceed, `.spm.toml` file will be created where you ran the command.

If you are running the command inside you projects directory (`~/Projects`), you probably would want to create a new directory for you project, and then run `spm init` in it. And I've got you covered, saving you 2 commands! Just run `spm init -m`, and spm will init your project in a new directory.


#### List projects
```bash
spm list
```

Spm will go through your projects folder and parse through found `.spm.toml`. It will list projects' name, description, path tags, favorites and programming language stats! Stats about programming lang is similar lika it is on github - a colorful line, with language percentages listed below. There are only 2 sizes: 100x1 and 50x1.




## TODO
- [ ] spm init
  - [x] base functionality
  - [ ] cla options for setting question answers (--author, --license etc)
  - [x] tag negation
- [ ] spm info
  - [x] base functionality
  - [x] color
  - [ ] short form
  - [ ] long form
  - [ ] format option (like in `date +%Y/%d/%m`)
  - [x] language stats
    - [x] ignore special non-source code dirs, like deps (`node_modules` etc)
    - [ ] exclude configs (.json, .toml, .ini etc)
- [ ] spm list (uses spm info)
  - [ ] base functionality
  - [ ] sorting options
- [ ] spm history
  - [ ] base functionality
  - [ ] setup a log file at alls
- [ ] spm delete
  - [ ] base functionality
  - [ ] compress + archive except of delete
- [ ] spm rename
  - [ ] base functionality
- [ ] spm open
  - [ ] open project in default code editor
  - [ ] cd into root project dir
    - [ ] bash function
    - [ ] zsh function
    - [ ] fish function
    - [ ] ps1 function
- [ ] spm run
  - [ ] base functionality
- [ ] misc
  - [ ] create projects under tag directories
  - [ ] scripts system
