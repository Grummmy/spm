# SPM -- Simple Project Manager
---

I always struggled with convenient project names and project placement. If I want to develop a minecraft mod or plugin, than I should uniquely name it or create a whole new directory tree: `minecraft/mods/mymod`, since I do many minecraft-related things. So, I came up with a solution - Simple Project Manager. It makes creating, viewing, and showing of project stats very convenient (or should at some point).

## Features
- Easily create projects with `spm init` (inspired by `npm init`)
- List projects and some stats, like programming lang. usage
- Fully delete or archive projects with one command
- Look at your work history: project creation, deletion, opening
- Show detailed info about project like modification or last opening time
- Open projects in your default code editor and change your terminal wd
- Configurable project defaults
- Configurable scripts!
- Basically, cross platform
- Other ideas accepted!

## Usage
When you first run the executable, it will try to find your default projects directory. It looks for paths like `~/Projects` (any case: `projects`, `PrOjEcTs` etc). If project dir wasnt guessed, the program will prompt you to choose it. Then, spm creates default config. It includes your user name as default author, default license, initial version, and a bunch of default language extentions mapped to language names and colors.
On Linux, config dir is `$XDG_CONFIG_HOME/spm/` (most likely `~/.config/spm/config.toml`). On Darwin it is `$HOME/Library/Application Support/spm/` and for Windows it is `%AppData%/spm/`. If `SPM_RUN_PORTABLE` environment variable is set (anithing except `""`), than executable dir is used as config directory.

#### Create project
```bash
spm init [name]
```
