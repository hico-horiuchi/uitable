// Package ansi provides methods for manipulating ANSI escape code
package ansi

import (
	"regexp"
	"strings"
)

type CodeStruct struct {
	Index    int
	Submatch string
}

const code = `\x1B\[\d{1,3}[mK]`

// Match returns ANSI escape codes in the given string.
func Match(s string) []CodeStruct {
	r := regexp.MustCompile(code)
	is := r.FindAllStringIndex(s, -1)
	a := make([]CodeStruct, len(is))
	for k, v := range r.FindAllStringSubmatch(s, -1) {
		a[k] = CodeStruct{
			Index:    is[k][0],
			Submatch: v[0],
		}
	}
	return a
}

// Patch joins ANSI escape codes to the given string.
func Patch(s string, a []CodeStruct) string {
	var b, c, i int
	for _, v := range a {
		i = v.Index + b
		if i > len(s) || strings.TrimSpace(s[i:]) == "..." {
			return s + "\x1b[0m"
		}
		if c = strings.Count(s[i:], "\n"); c == 0 {
			s = s[:i] + v.Submatch + s[i:]
			continue
		}
		s = s[:i] + v.Submatch + strings.Replace(s[i:], "\n", "\x1b[0m\n"+v.Submatch, -1)
		b += len("\x1b[0m\n"+v.Submatch) * c
	}
	return s
}

// Remove removes ANSI escape codes from the given string.
func Remove(s string) string {
	return regexp.MustCompile(code).ReplaceAllString(s, "")
}
