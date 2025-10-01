package cosineapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"cosine-cli/internal/auth"
)

type ProbeResult struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	Status      int           `json:"status"`
	DurationMs  int64         `json:"duration_ms"`
	ContentType string        `json:"content_type,omitempty"`
	Size        int           `json:"size_bytes,omitempty"`
	Error       string        `json:"error,omitempty"`
	Snippet     string        `json:"snippet,omitempty"`
	Timestamp   time.Time     `json:"timestamp"`
}

// defaultAPIPaths is a conservative set of likely endpoints to probe.
var defaultAPIPaths = []string{
	"/",
	"/api/ask",
	"/api/ping",
	"/api/me",
	"/api/user",
	"/api/account",
	"/api/search",
	"/api/chat",
	"/api/login",
	"/api/logout",
	"/api/v1/ask",
	"/api/v1/ping",
	"/api/v1/me",
	"/api/v1/search",
}

// ProbeEndpoints probes the given paths with GET requests against baseURL.
// It attaches the stored Cookie header when available, and returns structured results.
func ProbeEndpoints(baseURL string, apiPaths []string) ([]ProbeResult, error) {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	// normalize base
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse baseURL: %w", err)
	}
	cfg, _ := auth.Load() // best effort; empty cookie if not logged in

	client := &http.Client{Timeout: 10 * time.Second}
	results := make([]ProbeResult, 0, len(apiPaths))
	for _, p := range apiPaths {
		full := *u
		full.Path = path.Join(full.Path, p)
		req, err := http.NewRequest(http.MethodGet, full.String(), nil)
		start := time.Now()
		res := ProbeResult{
			Method:    http.MethodGet,
			URL:       full.String(),
			Timestamp: start,
		}
		if err != nil {
			res.Error = err.Error()
			results = append(results, res)
			continue
		}
		req.Header.Set("User-Agent", "cosine-cli/0.1")
		if cfg.CookieHeader != "" {
			req.Header.Set("Cookie", cfg.CookieHeader)
		}
		resp, err := client.Do(req)
		dur := time.Since(start)
		res.DurationMs = dur.Milliseconds()
		if err != nil {
			res.Error = err.Error()
			results = append(results, res)
			continue
		}
		func() {
			defer resp.Body.Close()
			res.Status = resp.StatusCode
			res.ContentType = resp.Header.Get("Content-Type")
			body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
			res.Size = len(body)
			// keep a compact snippet (single-line)
			res.Snippet = strings.TrimSpace(string(bytes.TrimSpace(body)))
		}()
		results = append(results, res)
	}
	return results, nil
}

// ProbeDefault runs ProbeEndpoints with a default list of common API paths.
func ProbeDefault(baseURL string) ([]ProbeResult, error) {
	return ProbeEndpoints(baseURL, defaultAPIPaths)
}

// LogProbe appends results to a log file under the cosine config directory.
// Returns the log file path.
func LogProbe(results []ProbeResult) (string, error) {
	dir, err := auth.EnsureDir()
	if err != nil {
		return "", err
	}
	logPath := path.Join(dir, "endpoints.log")
	// Use JSON lines for easy parsing.
	var buf bytes.Buffer
	for _, r := range results {
		j, _ := json.Marshal(r)
		buf.Write(j)
		buf.WriteByte('\n')
	}
	// append
	f, err := openFileAppend(logPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(buf.Bytes()); err != nil {
		return "", err
	}
	return logPath, nil
}