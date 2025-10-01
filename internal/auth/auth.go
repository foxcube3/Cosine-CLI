package auth

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	CookieHeader string `json:"cookie_header"` // aggregated cookies from Chrome for cosine.sh
	Source       string `json:"source"`        // "chrome"
}

func configDir() (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		return filepath.Join(xdg, "cosine"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "cosine"), nil
}

func configPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func EnsureDir() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	if _, statErr := os.Stat(dir); errors.Is(statErr, os.ErrNotExist) {
		if mkErr := os.MkdirAll(dir, 0o700); mkErr != nil {
			return "", mkErr
		}
	}
	return dir, nil
}

func Save(cfg Config) error {
	if _, err := EnsureDir(); err != nil {
		return err
	}
	path, err := configPath()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(&cfg)
}

func Load() (Config, error) {
	var cfg Config
	path, err := configPath()
	if err != nil {
		return cfg, err
	}
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}