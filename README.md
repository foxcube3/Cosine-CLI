# Cos CLI (cos)

This repository contains a decompiled overview of the Cosine CLI binary `cos` and a complete command reference based on observed behavior and runtime help output.

## Command Reference (observed)

- login — Login to the Cosine platform
- logout — Logout from the Cosine platform
- serve — Expose a directory to Genie
  - Flags: `--cwd string="."`, `--project string`, `--verbose`
- start — Start an interactive CLI session
  - Flags: `--cwd string="."`, `--project string`, `--log-file string="~/.cosine/cli/debug.log"`, `--simple, -s`
- diff — Open a side-by-side diff viewer for a unified diff
  - Flags: `--patch, -p string`, `--task string`, `--cwd string`
- terminal-setup — Configure VS Code terminal so Shift+Enter inserts a newline
  - Flags: `--dry-run`, `--print`
- daemon (hidden) — Serve a grpc server for Genie
  - Flags: `--cwd string`, `--token int=0`, `--port int=7001`

Global flags: `--help, -h`, `--version, -v`  
Version observed: `1.0.7`

## Default Behavior

Running `cos` with no arguments attempts to run the default command (observed: starts the interactive session) and checks authentication status first.

Example output when not logged in:
```
You are not logged in. Please run 'cosine login' first.
```

## Full help output (captured)

Top-level:
```
NAME:
   cos - Work with the Cosine platform and pair with Genie

USAGE:
   cos [global options] [command [command options]]

VERSION:
   1.0.7

COMMANDS:
   login           Login to the Cosine platform
   logout          Logout from the Cosine platform
   serve           Expose a directory to Genie
   start           Start an interactive CLI session
   diff            Open a side-by-side diff viewer for a unified diff
   terminal-setup  Configure VS Code terminal so Shift+Enter inserts a newline
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

start:
```
NAME:
   cos start - Start an interactive CLI session

USAGE:
   cos start [options]

OPTIONS:
   --cwd string       Directory to expose to Genie. Defaults to current directory. (default: ".")
   --project string   Cosine project ID.
   --log-file string  Path to log file (default: ~/.cosine/cli/debug.log)
   --simple, -s       Enable simple scrollable UI mode (stdout, prompt at bottom) (default: false)
   --help, -h         show help
```

serve:
```
NAME:
   cos serve - Expose a directory to Genie

USAGE:
   cos serve [options]

OPTIONS:
   --cwd string      Directory to expose to Genie. Defaults to current directory. (default: ".")
   --project string  Cosine project ID.
   --verbose         Enable verbose logging. (default: false)
   --help, -h        show help
```

daemon (hidden but help-visible):
```
NAME:
   cos daemon - Serve a grpc server for Genie

USAGE:
   cos daemon [options]

OPTIONS:
   --cwd string  Directory to expose to Genie. Defaults to current directory.
   --token int   Cosine authentication token. (default: 0)
   --port int    Port to serve the grpc server on. (default: 7001)
   --help, -h    show help
```

diff:
```
NAME:
   cos diff - Open a side-by-side diff viewer for a unified diff

USAGE:
   cos diff [options]

OPTIONS:
   --patch string, -p string  Path to a unified diff (use '-' to read from stdin)
   --task string              Task ID to fetch and view the full diff for
   --cwd string               Working directory for resolving relative paths (default: current directory)
   --help, -h                 show help
```

terminal-setup:
```
NAME:
   cos terminal-setup - Configure VS Code terminal so Shift+Enter inserts a newline

USAGE:
   cos terminal-setup [options]

OPTIONS:
   --dry-run   Print what would be changed but do not modify files (default: false)
   --print     Print the VS Code keybinding JSON entries (default: false)
   --help, -h  show help
```

## Minor UX polish note

- When not logged in, running `cos` prints:
  - `You are not logged in. Please run 'cosine login' first.`
- However, the installed binary and documented command are `cos`, and the correct command is:
  - `cos login`
- Recommendation: Update the message to reference the installed binary name (e.g., `cos login`) or support `cosine` via symlink/alias, to avoid confusion when only `cos` is installed.