package cosineapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"cosine-cli/internal/auth"
)

const defaultBaseURL = "https://cosine.sh"

type askRequest struct {
	Question string `json:"question"`
}

// Ask sends a question to cosine.sh using the stored Cookie header and returns the raw reply body.
// It requires that the user is logged in (Cookie header present).
func Ask(question string) (string, error) {
	if question == "" {
		return "", errors.New("empty question")
	}
	cfg, err := auth.Load()
	if err != nil {
		return "", fmt.Errorf("load auth: %w", err)
	}
	if cfg.CookieHeader == "" {
		return "", errors.New("not logged in; run `cosine login-chrome`, `cosine login-firefox`, or `cosine login-cookie` first")
	}

	bodyBytes, _ := json.Marshal(askRequest{Question: question})
	req, err := http.NewRequest("POST", defaultBaseURL+"/api/ask", bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cfg.CookieHeader)
	req.Header.Set("User-Agent", "cosine-cli/0.1")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if len(respBody) > 0 {
			return "", fmt.Errorf("cosine.sh API error: %s", bytes.TrimSpace(respBody))
		}
		return "", fmt.Errorf("cosine.sh API status: %s", resp.Status)
	}
	return string(respBody), nil
}