package browserlogin

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/browserutils/kooky"
	"github.com/browserutils/kooky/browser/chrome"
)

// CookieHeaderForCosine gathers non-expired cookies for cosine.sh from Chrome and formats
// them as a Cookie header string. It searches common Chrome profiles on the system.
func CookieHeaderForCosine() (string, error) {
	stores := []kooky.CookieStore{
		chrome.NewCookieStore(),
	}
	var cookies []*kooky.Cookie
	for _, store := range stores {
		if store == nil {
			continue
		}
		defer store.Close()
		cs, err := kooky.ReadCookies(store, kooky.DomainHasSuffix(".cosine.sh"), kooky.Valid, kooky.NameNotIn("__Host-CSRF"))
		if err != nil {
			// try next store; aggregate later if needed
			continue
		}
		cookies = append(cookies, cs...)
	}

	// Fallback for exact domain match if suffix yielded none
	if len(cookies) == 0 {
		for _, store := range stores {
			if store == nil {
				continue
			}
			cs, err := kooky.ReadCookies(store, kooky.DomainHasSuffix("cosine.sh"), kooky.Valid)
			if err == nil {
				cookies = append(cookies, cs...)
			}
		}
	}

	if len(cookies) == 0 {
		return "", fmt.Errorf("no cosine.sh cookies found in Chrome; ensure you're logged in via Chrome")
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