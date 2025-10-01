# Cos CLI (cos)

This repository contains a decompiled overview of the Cosine CLI binary `cos` and a brief command reference based on observed behavior.

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

Running `cos` with no arguments attempts to start the interactive session and checks authentication status first.

## Minor UX polish note

- When not logged in, running `cos` prints:
  - `You are not logged in. Please run 'cosine login' first.`
- However, the binary name and documented command are `cos`, and the correct command is:
  - `cos login`
- Recommendation: Update the message to reference the installed binary name (e.g., `cos login`) or resolve via a symlink/alias if `cosine` is also supported. This avoids confusion for users who only have `cos` installed.