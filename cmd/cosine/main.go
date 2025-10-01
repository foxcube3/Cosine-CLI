package main

import (
	"fmt"
	"os"
	"path/filepath"

	"cosine-cli/internal/auth"
	"cosine-cli/internal/search"
)

func usage() {
	fmt.Println("Cosine CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cosine login           Store a token for cosine.sh")
	fmt.Println("  cosine whoami          Show stored account info")
	fmt.Println("  cosine search <query>  Interactive search using ripgrep + fzf")
}

func cmdLogin() int {
	cfg, err := auth.PromptLogin(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "login aborted: %v\n", err)
		return 1
	}
	if err := auth.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		return 1
	}
	dir, _ := auth.EnsureDir()
	fmt.Printf("Saved credentials to %s\n", dir)
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