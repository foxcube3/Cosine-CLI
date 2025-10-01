package browserlogin

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/chrome"
	_ "github.com/browserutils/kooky/browser/firefox"
)

// CookieHeaderForCosine gathers non-expired cookies for cosine.sh from Chrome and formats
// them as a Cookie header string. It searches common Chrome profiles on the system.
// On some Windows setups with recent Chrome versions, cookie decryption can panic inside
// the kooky Chrome backend. We defensively recover and return a helpful error instead.
func CookieHeaderForCosine() (string, error) {
	header, err := cookieHeaderForCosineChromeSafe()
	if err != nil {
		return "", err
	}
	return header, nil
}

// CookieHeaderForCosineFirefox gathers cookies from Firefox instead of Chrome.
// Useful as a fallback when Chrome cookie decryption is unavailable on the system.
func CookieHeaderForCosineFirefox() (string, error) {
	return cookieHeaderForCosineWithFilters(kooky.DomainHasSuffix(".cosine.sh"), kooky.Valid, excludeCSRFFilter())
}

func excludeCSRFFilter() kooky.Filter {
	return kooky.FilterFunc(func(c *kooky.Cookie) bool {
		return c != nil && c.Name != "__Host-CSRF"
	})
}

// cookieHeaderForCosineChromeSafe wraps the Chrome traversal with a recover guard.
// This avoids crashing the CLI if the underlying library panics during decryption.
func cookieHeaderForCosineChromeSafe() (header string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Chrome cookie decryption failed (library panic): %v. Try `cosine login-firefox` or export cookies manually.", r)
		}
	}()
	return cookieHeaderForCosineWithFilters(kooky.DomainHasSuffix(".cosine.sh"), kooky.Valid, excludeCSRFFilter())
}

// cookieHeaderForCosineWithFilters is the internal logic to aggregate and format cookies.
func cookieHeaderForCosineWithFilters(filters ...kooky.Filter) (string, error) {
	seq := kooky.TraverseCookies(
		context.Background(),
		filters...,
	).OnlyCookies()

	var cookies []*kooky.Cookie
	for c := range seq {
		if c != nil {
			cookies = append(cookies, c)
		}
	}

	// Fallback for exact domain match if suffix yielded none
	if len(cookies) == 0 {
		seq2 := kooky.TraverseCookies(
			context.Background(),
			kooky.Valid,
			kooky.DomainHasSuffix("cosine.sh"),
		).OnlyCookies()
		for c := range seq2 {
			if c != nil {
				cookies = append(cookies, c)
			}
		}
	}

	if len(cookies) == 0 {
		return "", fmt.Errorf("no cosine.sh cookies found; ensure you're logged in to cosine.sh in the selected browser")
	}

	// Deduplicate by cookie name (keep latest expires)
	m := map[string]*kooky.Cookie{}
	for _, c := range cookies {
		if c == nil {
			continue
		}
		existing, ok := m[c.Name]
		if !ok || later(c, existing) {
			m[c.Name] = c
		}
	}
	names := make([]string, 0, len(m))
	for n := range m {
		names = append(names, n)
	}
	sort.Strings(names)

	parts := make([]string, 0, len(names))
	for _, n := range names {
		parts = append(parts, fmt.Sprintf("%s=%s", n, m[n].Value))
	}
	return strings.Join(parts, "; "), nil
}

func later(a, b *kooky.Cookie) bool {
	if a == nil || b == nil {
		return a != nil
	}
	ta := a.Expires
	tb := b.Expires
	if ta.IsZero() && tb.IsZero() {
		return true
	}
	if ta.IsZero() {
		return false
	}
	if tb.IsZero() {
		return true
	}
	return ta.After(tb)
}

// IsFreshHeader heuristically checks if the cookie header is likely valid (not empty, not expired-only).
func IsFreshHeader(header string) bool {
	return strings.Contains(header, "=") && !strings.Contains(header, "deleted")
}

// Expired returns true if the cookie is expired (helper, unused externally but kept for clarity)
func Expired(c *kooky.Cookie) bool {
	if c == nil {
		return true
	}
	if c.Expires.IsZero() {
		return false
	}
	return time.Now().After(c.Expires)
}