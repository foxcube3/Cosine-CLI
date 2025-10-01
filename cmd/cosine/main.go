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
	"cosine-cli/internal/cosineapi"
)

func usage() {
	fmt.Println("Cosine CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cosine login-chrome              Import cookies from Chrome for cosine.sh")
	fmt.Println("  cosine login-firefox             Import cookies from Firefox for cosine.sh")
	fmt.Println("  cosine login-cookie [header]     Save a raw Cookie header (arg or stdin)")
	fmt.Println("  cosine whoami                    Show stored account info")
	fmt.Println("  cosine search <query>            Interactive search using ripgrep + fzf")
	fmt.Println("  cosine list-search <query>       Print matches without fzf")
	fmt.Println("  cosine ask <question>            Ask a question to cosine.sh and print the reply")
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

func cmdDumpFirefoxCookies() int {
	// Uses the same underlying browser-agnostic collector; message tailored for Firefox.
	header, err := browserlogin.CookieHeaderForCosine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read Firefox cookies: %v\n", err)
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
		// Surface concise child stderr
		msg := trimChildError(stderr.String(), err)
		fmt.Fprintf(os.Stderr, "failed to read Chrome cookies (subprocess): %s\n", msg)

		// Interactive fallback: prompt user to paste cookie header
		fmt.Fprintln(os.Stderr, "\nInteractive fallback:")
		fmt.Fprintln(os.Stderr, "  1) In Chrome, open https://cosine.sh and ensure you are logged in.")
		fmt.Fprintln(os.Stderr, "  2) Open DevTools (F12) -> Network tab.")
		fmt.Fprintln(os.Stderr, "  3) Reload the page, click a request to cosine.sh, then in Headers")
		fmt.Fprintln(os.Stderr, "     copy the full value of the 'Cookie' request header.")
		fmt.Fprintln(os.Stderr, "Paste the Cookie header below (single line). Press Enter on an empty line to cancel.")
		fmt.Fprint(os.Stderr, "> ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		header := strings.TrimSpace(line)
		if header == "" {
			fmt.Fprintln(os.Stderr, "cancelled")
			return 1
		}
		if !strings.Contains(header, "=") {
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
	header := strings.TrimSpace(stdout.String())
	if header == "" {
		fmt.Fprintf(os.Stderr, "no cookies returned from Chrome\n")
		fmt.Fprintln(os.Stderr, "\nWorkaround: use manual import. See:")
		fmt.Fprintln(os.Stderr, "  cosine login-cookie \"<Cookie header>\"")
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

func cmdListSearch(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "missing query")
		return 1
	}
	query := args[0]
	cwd, _ := os.Getwd()
	results, err := search.SearchList(query, cwd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "list-search: %v\n", err)
		return 1
	}
	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "no matches")
		return 1
	}
	for _, r := range results {
		fmt.Printf("%s:%d: %s\n", r.File, r.Line, r.Text)
	}
	return 0
}

func cmdAsk(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "missing question")
		return 1
	}
	question := strings.Join(args, " ")
	reply, err := cosineapi.Ask(strings.TrimSpace(question))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ask: %v\n", err)
		return 1
	}
	// Print raw reply body
	fmt.Println(strings.TrimSpace(reply))
	return 0
}

func trimChildError(stderr string, runErr error) string {
	s := strings.TrimSpace(stderr)
	if s == "" {
		if runErr != nil {
			return runErr.Error()
		}
		return "unknown error"
	}
	// Prefer the 'panic:' line if present
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "panic:") {
			return line
		}
	}
	// Otherwise return first non-empty line
	sc = bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			return line
		}
	}
	if runErr != nil {
		return runErr.Error()
	}
	return "unknown error"
}

func cmdLoginFirefox() int {
	// Run cookie retrieval in a subprocess to isolate potential panics
	self, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "resolve executable: %v\n", err)
		return 1
	}
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(self, "dump-firefox-cookies")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		msg := trimChildError(stderr.String(), err)
		fmt.Fprintf(os.Stderr, "failed to read Firefox cookies (subprocess): %s\n", msg)

		// Interactive fallback
		fmt.Fprintln(os.Stderr, "\nInteractive fallback:")
		fmt.Fprintln(os.Stderr, "  1) In Firefox, open https://cosine.sh and ensure you are logged in.")
		fmt.Fprintln(os.Stderr, "  2) Open Developer Tools (F12) -> Network.")
		fmt.Fprintln(os.Stderr, "  3) Reload the page, select a request to cosine.sh, then in Headers copy")
		fmt.Fprintln(os.Stderr, "     the full value of the 'Cookie' request header.")
		fmt.Fprintln(os.Stderr, "Paste the Cookie header below (single line). Press Enter on an empty line to cancel.")
		fmt.Fprint(os.Stderr, "> ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		header := strings.TrimSpace(line)
		if header == "" {
			fmt.Fprintln(os.Stderr, "cancelled")
			return 1
		}
		if !strings.Contains(header, "=") {
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
	header := strings.TrimSpace(stdout.String())
	if header == "" {
		fmt.Fprintf(os.Stderr, "no cookies returned from Firefox\n")
		fmt.Fprintln(os.Stderr, "\nWorkaround: use manual import. See:")
		fmt.Fprintln(os.Stderr, "  cosine login-cookie \"<Cookie header>\"")
		return 1
	}

	cfg := auth.Config{
		CookieHeader: header,
		Source:       "firefox",
	}
	if err := auth.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save credentials: %v\n", err)
		return 1
	}
	fmt.Println("Imported cookies from Firefox for cosine.sh")
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
	case "login-firefox":
		os.Exit(cmdLoginFirefox())
	case "dump-chrome-cookies":
		os.Exit(cmdDumpChromeCookies())
	case "dump-firefox-cookies":
		os.Exit(cmdDumpFirefoxCookies())
	case "login-cookie":
		os.Exit(cmdLoginCookie(os.Args[2:]))
	case "whoami":
		os.Exit(cmdWhoAmI())
	case "search":
		os.Exit(cmdSearch(os.Args[2:]))
	case "list-search":
		os.Exit(cmdListSearch(os.Args[2:]))
	case "ask":
		os.Exit(cmdAsk(os.Args[2:]))
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}