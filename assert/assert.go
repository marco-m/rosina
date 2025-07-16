package assert

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/marco-m/rosina/diff"
)

func Equal[T comparable](t testing.TB, have T, want T, desc string) {
	t.Helper()
	// https://stackoverflow.com/a/71588125/561422
	switch any(have).(type) {
	case string:
		textEqual(t, any(have).(string), any(want).(string), desc)
	default:
		if have != want {
			t.Fatalf("\n%s mismatch:\nhave: %v\nwant: %v", desc, have, want)
		}
	}
}

func textEqual(t testing.TB, have string, want string, desc string) {
	t.Helper()
	delta := diff.TextDiff("want", "have", want, have)
	if delta != "" {
		t.Fatalf("\n%s mismatch:\n%s", desc, delta)
	}
}

func DeepEqual[T any](t testing.TB, have T, want T, desc string) {
	t.Helper()
	if delta := diff.AnyDiff(have, want); delta != "" {
		t.Fatalf("\n%s mismatch: +have -want:\n%s", desc, delta)
	}
}

func True(t testing.TB, pred bool, desc string) {
	t.Helper()
	if !pred {
		t.Fatalf("\n%s predicate mismatch:have: %v\nwant: true", desc, pred)
	}
}

func False(t testing.TB, pred bool, desc string) {
	t.Helper()
	if pred {
		t.Fatalf("\n%s predicate mismatch:have: %v\nwant: false", desc, pred)
	}
}

func Contains(t testing.TB, haystack, needle string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Fatalf("\nsubstring not found in string:\nsubstring: %q\nstring:    %q",
			needle, haystack)
	}
}

func NoError(t testing.TB, err error, desc string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s:\nhave: %v (%T)\nwant: <no error>", desc, err, err)
	}
}

func ErrorContains(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("\nhave: <no error>\nwant: <an error>")
	}
	Contains(t, err.Error(), want)
}

func ErrorIs(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatalf("\nhave: <no error>\nwant: %q (%T)", want, want)
	}
	if !errors.Is(err, want) {
		t.Fatalf("\nhave: %s (%T)\nwant: %q (%T)", err, err, want, want)
	}
}

func PanicTextContains(t testing.TB, fn func(), want string) {
	t.Helper()

	var recovered any
	// This function wrapper is needed to make t.Helper() report the
	// correct file after the panic is recovered.
	func() {
		defer func() {
			recovered = recover()
		}()
		fn()
	}()

	if recovered == nil {
		t.Fatalf("\nhave: <no panic>\nwant panic: %s", want)
	}
	msg := fmt.Sprint(recovered)
	Equal(t, msg, want, "panic message")
}
