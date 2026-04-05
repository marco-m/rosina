package internal

import (
	"strings"
	"testing"

	"github.com/marco-m/rosina/diff"
)

type FatalFunc func(format string, args ...any)

func Equal[T comparable](fn FatalFunc, t testing.TB, have T, want T, desc string) {
	t.Helper()
	// https://stackoverflow.com/a/71588125/561422
	switch any(have).(type) {
	case string:
		textEqual(fn, t, any(have).(string), any(want).(string), desc)
	default:
		if have != want {
			fn("%s:\nhave: %v\nwant: %v", desc, have, want)
		}
	}
}

func textEqual(fn FatalFunc, t testing.TB, have string, want string, desc string) {
	t.Helper()
	if !strings.Contains(have, "\n") && !strings.Contains(want, "\n") {
		if have != want {
			fn("%s:\nhave: %v\nwant: %v", desc, have, want)
		}
		return
	}
	if delta := diff.TextDiff("want", "have", want, have); delta != "" {
		fn("%s:\n%s", desc, delta)
	}
}
