Cosine CLI

A minimal CLI that:
- Logs in to cosine.sh by importing your authenticated Chrome session cookies
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
  # Import cookies from your authenticated Chrome session for cosine.sh
  ./cosine login-chrome

  # Show stored account information
  ./cosine whoami

  # Interactive code/content search using ripgrep + fzf
  ./cosine search <query>

Notes
- Authentication: cosine.sh does not provide a public programmatic login API. This CLI imports your existing authenticated Chrome session cookies for cosine.sh (using github.com/zellyn/kooky).
- Chrome session import: You must already be logged in to cosine.sh in Chrome on this machine. On macOS/Windows/Linux, cookies are decrypted using system facilities supported by kooky. If no cookies are found, log in via Chrome first and retry.
- Dependencies: This CLI executes the external binaries rg (ripgrep) and fzf supplied by their respective repositories:
  - https://github.com/BurntSushi/ripgrep
  - https://github.com/junegunn/fzf
- The CLI requires these binaries to be installed and available on PATH.