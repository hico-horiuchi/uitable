// Package ansi provides methods for manipulating ANSI escape code
package ansi

import (
	"regexp"
	"strings"
)

type code struct {
	index    int
	submatch string
}

const pattern = `\x1B\[\d{1,3}[mK]`

// Match returns ANSI escape codes in the given string.
func Match(s string) []code {
	r := regexp.MustCompile(pattern)
	is := r.FindAllStringIndex(s, -1)
	a := make([]code, len(is))
	for k, v := range r.FindAllStringSubmatch(s, -1) {
		a[k] = code{
			index:    is[k][0],
			submatch: v[0],
		}
	}
	return a
}

// Patch joins ANSI escape codes to the given string.
func Patch(s string, a []code) string {
	var b, c, i int
	for _, v := range a {
		i = v.index + b
		if i > len(s) {
			return s + "\x1b[0m"
		} else if strings.TrimSpace(s[i:]) == "..." {
			return strings.TrimRight(s, "...") + "\x1b[0m..."
		}
		b += strings.Count(s[c:i], " \n")
		c = v.index + b
		s = s[:c] + v.submatch + s[c:]
	}
	return s
}

// Remove removes ANSI escape codes from the given string.
func Remove(s string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(s, "")
}
