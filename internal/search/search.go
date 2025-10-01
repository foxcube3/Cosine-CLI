package search

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Result struct {
	File  string
	Line  int
	Text  string
	Raw   string
}

func rg(query string, dir string) (*bytes.Buffer, error) {
	cmd := exec.Command("rg", "--line-number", "--no-heading", "--hidden", "--glob", "!.git", query)
	cmd.Dir = dir
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		// If no matches, ripgrep returns exit code 1; treat as no results, not an error.
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 && out.Len() == 0 {
			return &bytes.Buffer{}, nil
		}
		return nil, fmt.Errorf("ripgrep error: %v: %s", err, stderr.String())
	}
	return &out, nil
}

func fzf(input *bytes.Buffer) (string, error) {
	cmd := exec.Command("fzf", "--ansi", "--no-sort", "--prompt", "rg> ")
	cmd.Stdin = bytes.NewReader(input.Bytes())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// likely ESC or no selection
			return "", errors.New("no selection")
		}
		return "", fmt.Errorf("fzf error: %v: %s", err, stderr.String())
	}
	return strings.TrimSpace(out.String()), nil
}

func parse(line string) (Result, error) {
	// Expect: file:line:text (ripgrep default with --line-number --no-heading)
	colon1 := strings.IndexByte(line, ':')
	if colon1 < 0 {
		return Result{}, errors.New("unrecognized line format")
	}
	colon2 := strings.IndexByte(line[colon1+1:], ':')
	if colon2 < 0 {
		return Result{}, errors.New("unrecognized line format")
	}
	colon2 += colon1 + 1
	file := line[:colon1]
	lnStr := line[colon1+1 : colon2]
	ln, err := strconv.Atoi(lnStr)
	if err != nil {
		return Result{}, err
	}
	text := line[colon2+1:]
	return Result{File: file, Line: ln, Text: text, Raw: line}, nil
}

func Open(result Result) (string, error) {
	abs, err := filepath.Abs(result.File)
	if err != nil {
		return "", err
	}
	f, err := os.Open(abs)
	if err != nil {
		return "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	cur := 0
	var snippet string
	for sc.Scan() {
		cur++
		if cur == result.Line {
			snippet = sc.Text()
			break
		}
	}
	return fmt.Sprintf("%s:%d: %s", abs, result.Line, snippet), nil
}

func SearchInteractive(query string, dir string) (Result, error) {
	out, err := rg(query, dir)
	if err != nil {
		return Result{}, err
	}
	if out.Len() == 0 {
		return Result{}, errors.New("no matches")
	}
	selected, err := fzf(out)
	if err != nil {
		return Result{}, err
	}
	return parse(selected)
}