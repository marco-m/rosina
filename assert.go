package rosina

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, have T, want T, desc string) {
	t.Helper()
	if have != want {
		t.Fatalf("\n%s mismatch:\nhave: %v\nwant: %v", desc, have, want)
	}
}

func AssertTextEqual(t testing.TB, have string, want string, desc string) {
	t.Helper()
	delta := TextDiff("want", []byte(want), "have", []byte(have))
	if delta != nil {
		t.Fatalf("\n%s mismatch: +have -want:\n%s", desc, string(delta))
	}
}

func AssertDeepEqual[T any](t testing.TB, have T, want T, desc string) {
	t.Helper()
	if delta := AnyDiff(have, want); delta != "" {
		t.Fatalf("\n%s mismatch: +have -want:\n%s", desc, delta)
	}
}

func AssertTrue(t testing.TB, pred bool, desc string) {
	t.Helper()
	if !pred {
		t.Fatalf("\n%s predicate mismatch:have: %v\nwant: true", desc, pred)
	}
}

func AssertFalse(t testing.TB, pred bool, desc string) {
	t.Helper()
	if pred {
		t.Fatalf("\n%s predicate mismatch:have: %v\nwant: false", desc, pred)
	}
}

func AssertContains(t testing.TB, haystack, needle string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Fatalf("\nsubstring not found in string:\nsubstring: %q\nstring:    %q",
			needle, haystack)
	}
}

func AssertIsNil(t testing.TB, x any) {
	t.Helper()
	if x != nil {
		t.Fatalf("\nhave: %v (%T)\nwant: <nil>", x, x)
	}
}

func AssertIsNotNil(t testing.TB, x any) {
	t.Helper()
	if x == nil {
		t.Fatalf("\nhave: <nil>\nwant: not nil")
	}
}

func AssertErrorIs(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatalf("\nhave: <no error>\nwant: %q (%T)", want, want)
	}
	if !errors.Is(err, want) {
		t.Fatalf("\nhave: %s (%T)\nwant: %q (%T)", err, err, want, want)
	}
}

func AssertPanicTextContains(t testing.TB, fn func(), want string) {
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
	AssertEqual(t, msg, want, "panic message")
}
