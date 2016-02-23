package ansi

import (
	"testing"
)

func TestMatch(t *testing.T) {
	s := "\x1b[1mbold\x1b[0m"
	got := Match(s)
	if len(got) != 2 {
		t.Fatal("want", []CodeStruct{
			CodeStruct{
				Index:    0,
				Submatch: "\x1b[1m",
			},
			CodeStruct{
				Index:    8,
				Submatch: "\x1b[0m",
			},
		}, "got", got)
	}

	s = "\x1b[31mforeground\x1b[0m"
	got = Match(s)
	if len(got) != 2 {
		t.Fatal("want", []CodeStruct{
			CodeStruct{
				Index:    0,
				Submatch: "\x1b[31m",
			},
			CodeStruct{
				Index:    15,
				Submatch: "\x1b[0m",
			},
		}, "got", got)
	}

	s = "\x1b[1m\x1b[31mbold,foreground\x1b[0m"
	got = Match(s)
	if len(got) != 3 {
		t.Fatal("want", []CodeStruct{
			CodeStruct{
				Index:    0,
				Submatch: "\x1b[1m",
			},
			CodeStruct{
				Index:    4,
				Submatch: "\x1b[31m",
			},
			CodeStruct{
				Index:    24,
				Submatch: "\x1b[0m",
			},
		}, "got", got)
	}
}

func TestPatch(t *testing.T) {
	s := "bold"
	a := []CodeStruct{
		CodeStruct{
			Index:    0,
			Submatch: "\x1b[1m",
		},
		CodeStruct{
			Index:    8,
			Submatch: "\x1b[0m",
		},
	}
	got := Patch(s, a)
	if got != "\x1b[1mbold\x1b[0m" {
		t.Fatal("want", "\x1b[1mbold\x1b[0m", "got", got)
	}

	s = "foreground"
	a = []CodeStruct{
		CodeStruct{
			Index:    0,
			Submatch: "\x1b[31m",
		},
		CodeStruct{
			Index:    15,
			Submatch: "\x1b[0m",
		},
	}
	got = Patch(s, a)
	if got != "\x1b[31mforeground\x1b[0m" {
		t.Fatal("want", "\x1b[31mforeground\x1b[0m", "got", got)
	}

	s = "bold,foreground"
	a = []CodeStruct{
		CodeStruct{
			Index:    0,
			Submatch: "\x1b[1m",
		},
		CodeStruct{
			Index:    4,
			Submatch: "\x1b[31m",
		},
		CodeStruct{
			Index:    24,
			Submatch: "\x1b[0m",
		},
	}
	got = Patch(s, a)
	if got != "\x1b[1m\x1b[31mbold,foreground\x1b[0m" {
		t.Fatal("want", "\x1b[1m\x1b[31mbold,foreground\x1b[0m", "got", got)
	}
}

func TestRemove(t *testing.T) {
	s := "\x1b[1mbold\x1b[0m"
	got := Remove(s)
	if got != "bold" {
		t.Fatal("want", "bold", "got", got)
	}

	s = "\x1b[31mforeground\x1b[0m"
	got = Remove(s)
	if got != "foreground" {
		t.Fatal("want", "foreground", "got", got)
	}

	s = "\x1b[1m\x1b[31mbold,foreground\x1b[0m"
	got = Remove(s)
	if got != "bold,foreground" {
		t.Fatal("want", "bold,foreground", "got", got)
	}
}
