package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"cosine-cli/internal/auth"
	"cosine-cli/internal/browserlogin"
	"cosine-cli/internal/search"
)

func usage() {
	fmt.Println("Cosine CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cosine login-chrome              Import cookies from Chrome for cosine.sh")
	fmt.Println("  cosine login-cookie [header]     Save a raw Cookie header (arg or stdin)")
	fmt.Println("  cosine whoami                    Show stored account info")
	fmt.Println("  cosine search <query>            Interactive search using ripgrep + fzf")
}

func cmdDumpChromeCookies() int {
	header, err := browserlogin.CookieHeaderForCosine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read Chrome cookies: %v\n", err)
		return 1
	}
	fmt.Print(header)
	return 0
}

func cmdLoginChrome() int {
	// Run cookie retrieval in a subprocess to isolate potential panics
	self, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "resolve executable: %v\n", err)
		return 1
	}
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(self, "dump-chrome-cookies")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		// Surface child stderr to aid debugging
		msg := stderr.String()
		if msg == "" {
			msg = err.Error()
		}
		fmt.Fprintf(os.Stderr, "failed to read Chrome cookies (subprocess): %s\n", msg)
		return 1
	}
	header := strings.TrimSpace(stdout.String())
	if header == "" {
		fmt.Fprintf(os.Stderr, "no cookies returned from Chrome\n")
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

func cmdLoginCookie(args []string) int {
	var header string
	if len(args) > 0 {
		header = strings.TrimSpace(strings.Join(args, " "))
	} else {
		// Read from stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Fprintln(os.Stderr, "no cookie header provided; pass as argument or pipe via stdin")
			return 1
		}
		b := &strings.Builder{}
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			b.WriteString(line)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "read stdin: %v\n", err)
				return 1
			}
		}
		header = strings.TrimSpace(b.String())
	}
	if header == "" || !strings.Contains(header, "=") {
		fmt.Fprintln(os.Stderr, "invalid cookie header")
		return 1
	}
	cfg := auth.Config{
		CookieHeader: header,
		Source:       "manual",
	}
	if err := auth.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		return 1
	}
	fmt.Println("Saved cookie header for cosine.sh (manual)")
	return 0
}

func cmdWhoAmI() int {
	cfg, err := auth.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "not logged in: %v\n", err)
		return 1
	}
	fmt.Println("cosine.sh")
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
	case "login-chrome":
		os.Exit(cmdLoginChrome())
	case "dump-chrome-cookies":
		os.Exit(cmdDumpChromeCookies())
	case "login-cookie":
		os.Exit(cmdLoginCookie(os.Args[2:]))
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