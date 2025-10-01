package main

import (
	"fmt"
	"os"

	"cosine-cli/internal/auth"
	"cosine-cli/internal/browserlogin"
	"cosine-cli/internal/search"
)

func usage() {
	fmt.Println("Cosine CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cosine login              Store a token for cosine.sh (manual)")
	fmt.Println("  cosine login-chrome       Import cookies from Chrome for cosine.sh")
	fmt.Println("  cosine whoami             Show stored account info")
	fmt.Println("  cosine search <query>     Interactive search using ripgrep + fzf")
}

func cmdLogin() int {
	cfg, err := auth.PromptLogin(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "login aborted: %v\n", err)
		return 1
	}
	cfg.Source = "manual"
	if err := auth.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		return 1
	}
	dir, _ := auth.EnsureDir()
	fmt.Printf("Saved credentials to %s\n", dir)
	return 0
}

func cmdLoginChrome() int {
	header, err := browserlogin.CookieHeaderForCosine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read Chrome cookies: %v\n", err)
		return 1
	}
	cfg := auth.Config{
		CookieHeader: header,
		Source:       "chrome",
	}
	if err := auth.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		return 1
	}
	fmt.Println("Imported cookies from Chrome for cosine.sh")
	return 0
}

func cmdWhoAmI() int {
	cfg, err := auth.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "not logged in: %v\n", err)
		return 1
	}
	fmt.Println("cosine.sh")
	if cfg.Email != "" {
		fmt.Printf("  Email: %s\n", cfg.Email)
	}
	if cfg.Token != "" {
		fmt.Printf("  Token: %s... (%d chars)\n", cfg.Token[:min(4, len(cfg.Token))], len(cfg.Token))
	}
	if cfg.CookieHeader != "" {
		fmt.Printf("  Cookies: %d bytes from %s\n", len(cfg.CookieHeader), cfg.Source)
	}
	if cfg.Source != "" {
		fmt.Printf("  Source: %s\n", cfg.Source)
	}
	return 0
}

func cmdSearch(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "missing query")
		return 1
	}
	query := args[0]
	cwd, _ := os.Getwd()
	res, err := search.SearchInteractive(query, cwd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "search: %v\n", err)
		return 1
	}
	line, err := search.Open(res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
		return 1
	}
	fmt.Println(line)
	return 0
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "login":
		os.Exit(cmdLogin())
	case "login-chrome":
		os.Exit(cmdLoginChrome())
	case "whoami":
		os.Exit(cmdWhoAmI())
	case "search":
		os.Exit(cmdSearch(os.Args[2:]))
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}