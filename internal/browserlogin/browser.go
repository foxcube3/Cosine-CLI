package browserlogin

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/all"
)

// CookieHeaderForCosine gathers cookies for cosine.sh from installed browsers
// via kooky and formats them into a Cookie header string.
func CookieHeaderForCosine() (string, error) {
	return cookieHeaderForCosineChromeSafe()
}

func excludeCSRFFilter() kooky.Filter {
	return kooky.FilterFunc(func(c *kooky.Cookie) bool {
		return c != nil && c.Name != "__Host-CSRF"
	})
}

// cookieHeaderForCosineChromeSafe wraps the cookie collection with a recover guard
// to avoid crashing if the underlying library panics during decryption on some systems.
func cookieHeaderForCosineChromeSafe() (header string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Chrome cookie decryption failed (library panic): %v. Consider exporting cookies manually or using another browser.", r)
		}
	}()
	return cookieHeaderForCosineWithFilters(kooky.Valid, excludeCSRFFilter())
}

// cookieHeaderForCosineWithFilters is the internal logic to aggregate and format cookies.
func cookieHeaderForCosineWithFilters(filters ...kooky.Filter) (string, error) {
	read := kooky.ReadCookies(filters...)

	// Select cookies for cosine.sh and exclude CSRF cookie
	var cookies []*kooky.Cookie
	for _, c := range read {
		if c == nil {
			continue
		}
		d := strings.ToLower(c.Domain)
		if strings.Contains(d, "cosine.sh") && c.Name != "__Host-CSRF" {
			cookies = append(cookies, c)
		}
	}

	// If still none, try without filters for maximum compatibility.
	if len(cookies) == 0 {
		read2 := kooky.ReadCookies()
		for _, c := range read2 {
			if c == nil {
				continue
			}
			d := strings.ToLower(c.Domain)
			if strings.Contains(d, "cosine.sh") && c.Name != "__Host-CSRF" {
				cookies = append(cookies, c)
			}
		}
	}

	if len(cookies) == 0 {
		return "", fmt.Errorf("no cosine.sh cookies found; ensure you're logged in to cosine.sh in your browser")
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
</old_code>
<old_code>
// cookieHeaderForCosineWithFilters is the internal logic to aggregate and format cookies.
func cookieHeaderForCosineWithFilters(filters ...kooky.Filter) (string, error) {
	// kooky v0.2.2 API exposes ReadCookies instead of TraverseCookies
	read := kooky.ReadCookies(filters...)

	var cookies []*kooky.Cookie
	for _, c := range read {
		if c != nil {
			cookies = append(cookies, c)
		}
	}

	// Fallback for exact domain match if suffix yielded none
	if len(cookies) == 0 {
		read2 := kooky.ReadCookies(
			kooky.Valid,
			kooky.DomainHasSuffix("cosine.sh"),
		)
		for _, c := range read2 {
			if c != nil {
				cookies = append(cookies, c)
			}
		}
	}
</old_code>
<new_code>
// cookieHeaderForCosineWithFilters is the internal logic to aggregate and format cookies.
func cookieHeaderForCosineWithFilters(filters ...kooky.Filter) (string, error) {
	// kooky v0.2.2 API exposes ReadCookies instead of TraverseCookies
	read := kooky.ReadCookies(filters...)

	// Manually select cosine.sh cookies to avoid suffix semantics differing across versions.
	var cookies []*kooky.Cookie
	for _, c := range read {
		if c == nil {
			continue
		}
		d := strings.ToLower(c.Domain)
		if strings.Contains(d, "cosine.sh") && c.Name != "__Host-CSRF" {
			cookies = append(cookies, c)
		}
	}

	// If still none, try without filters for maximum compatibility.
	if len(cookies) == 0 {
		read2 := kooky.ReadCookies()
		for _, c := range read2 {
			if c == nil {
				continue
			}
			d := strings.ToLower(c.Domain)
			if strings.Contains(d, "cosine.sh") && c.Name != "__Host-CSRF" {
				cookies = append(cookies, c)
			}
		}
	}

// cookieHeaderForCosineWithFilters is the internal logic to aggregate and format cookies.
func cookieHeaderForCosineWithFilters(filters ...kooky.Filter) (string, error) {
	// kooky v0.2.2 API exposes ReadCookies instead of TraverseCookies
	read := kooky.ReadCookies(filters...)

	var cookies []*kooky.Cookie
	for _, c := range read {
		if c != nil {
			cookies = append(cookies, c)
		}
	}

	// Fallback for exact domain match if suffix yielded none
	if len(cookies) == 0 {
		read2 := kooky.ReadCookies(
			kooky.Valid,
			kooky.DomainHasSuffix("cosine.sh"),
		)
		for _, c := range read2 {
			if c != nil {
				cookies = append(cookies, c)
			}
		}
	}

	if len(cookies) == 0 {
		return "", fmt.Errorf("no cosine.sh cookies found; ensure you're logged in to cosine.sh in Chrome")
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