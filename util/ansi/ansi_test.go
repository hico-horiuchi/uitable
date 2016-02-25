package ansi

import (
	"testing"
)

func TestMatch(t *testing.T) {
	s := "\x1b[1mbold\x1b[0m"
	got := Match(s)
	want := []code{
		code{
			index:    0,
			submatch: "\x1b[1m",
		},
		code{
			index:    8,
			submatch: "\x1b[0m",
		},
	}
	if len(got) != len(want) {
		t.Fatal("want", want, "got", got)
	}
	for k, v := range got {
		if v.index != want[k].index || v.submatch != want[k].submatch {
			t.Fatal("want", want, "got", got)
			break
		}
	}

	s = "\x1b[31mforeground\x1b[0m"
	got = Match(s)
	want = []code{
		code{
			index:    0,
			submatch: "\x1b[31m",
		},
		code{
			index:    15,
			submatch: "\x1b[0m",
		},
	}
	if len(got) != len(want) {
		t.Fatal("want", want, "got", got)
	}
	for k, v := range got {
		if v.index != want[k].index || v.submatch != want[k].submatch {
			t.Fatal("want", want, "got", got)
			break
		}
	}

	s = "\x1b[1m\x1b[31mbold,foreground\x1b[0m"
	got = Match(s)
	want = []code{
		code{
			index:    0,
			submatch: "\x1b[1m",
		},
		code{
			index:    4,
			submatch: "\x1b[31m",
		},
		code{
			index:    24,
			submatch: "\x1b[0m",
		},
	}
	if len(got) != len(want) {
		t.Fatal("want", want, "got", got)
	}
	for k, v := range got {
		if v.index != want[k].index || v.submatch != want[k].submatch {
			t.Fatal("want", want, "got", got)
			break
		}
	}
}

func TestPatch(t *testing.T) {
	s := "bold"
	a := []code{
		code{
			index:    0,
			submatch: "\x1b[1m",
		},
		code{
			index:    8,
			submatch: "\x1b[0m",
		},
	}
	got := Patch(s, a)
	if got != "\x1b[1mbold\x1b[0m" {
		t.Fatal("want", "\x1b[1mbold\x1b[0m", "got", got)
	}

	s = "foreground"
	a = []code{
		code{
			index:    0,
			submatch: "\x1b[31m",
		},
		code{
			index:    15,
			submatch: "\x1b[0m",
		},
	}
	got = Patch(s, a)
	if got != "\x1b[31mforeground\x1b[0m" {
		t.Fatal("want", "\x1b[31mforeground\x1b[0m", "got", got)
	}

	s = "bold,foreground"
	a = []code{
		code{
			index:    0,
			submatch: "\x1b[1m",
		},
		code{
			index:    4,
			submatch: "\x1b[31m",
		},
		code{
			index:    24,
			submatch: "\x1b[0m",
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
