Main command - `spm` - Simple Project Manager

- new \<name> [description] [-l --lang] [-t --tag] [-p --path]
    - creates new project with specified `name` and optional `description`
        - cretaes new directory with project name and tries to cd in it
        - openes default code editor with directory argument
    - lang : specifies project's main language
    - tag : adds project tags
    - path : dont create new directory, record speciifed directory as project root

- list [-s --sort] [--asc / --des] [-L --lines] [--no-exceptions]
    - lists created projects with a bar representing languages project includes (percentage counted by character count in each file)
        - lang counter excludes common non-code directories like dependencies: node_modules, .git, dist, build
    - sort : sorts project list by: date (creation), lang, open (last opened, default sorting)
    - asc / des : ascending or descending sorting order, defaul is ascendiong
    - lines : cound lines instead of characters
    - no-exceptions : dont exclude non-code directories/files

- delete \<name> [reason] [-f --force] [-P --permanent]
    - prompts confirmation, than compresses code files (not dependencies) and moves to pm trash directory
    - saves project name, description, stats(lang bar, creation time and deletion time) and deletion reason to history file
    - confirm : forcefully deletes project directory, no confirmation prompt will be asked
    - permanent : deletes project directory instead of compressing and moving to trash

- history [-T --type]
    - logs projects history
    - type : log type to be printed: delete / create

- info \<name>
    - print name, description, lang bar and tags of project `name` provided

- rename \<old> \<new>
    - rename project with `old` name to `new` name

- open \<name>
    - opens specified project in default editor
