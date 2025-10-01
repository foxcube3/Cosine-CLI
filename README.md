Cosine CLI

A minimal CLI that:
- Logs in to cosine.sh by storing a user-provided token locally.
- Uses ripgrep (rg) and fzf to provide a fast interactive search across your working directory.

Requirements
- Go 1.21+
- ripgrep (rg) installed and in PATH
- fzf installed and in PATH

Install ripgrep and fzf
- macOS (Homebrew):
  brew install ripgrep fzf
  /opt/homebrew/opt/fzf/install
- Ubuntu/Debian:
  sudo apt-get update
  sudo apt-get install ripgrep fzf
- Windows (Chocolatey):
  choco install ripgrep
  choco install fzf

Build
  go build ./cmd/cosine

Usage
  # First-time login: store a token and optional email
  ./cosine login

  # Show stored account information
  ./cosine whoami

  # Interactive code/content search using ripgrep + fzf
  ./cosine search <query>

Notes
- Authentication: cosine.sh does not provide a public programmatic login API. The CLI stores a token you provide (e.g., a session or API token) for future use. No network authentication flow is performed.
- Dependencies: This CLI executes the external binaries rg (ripgrep) and fzf supplied by their respective repositories:
  - https://github.com/BurntSushi/ripgrep
  - https://github.com/junegunn/fzf
- The CLI requires these binaries to be installed and available on PATH.