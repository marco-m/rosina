package assert

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/marco-m/rosina/diff"
	"github.com/marco-m/rosina/internal"
)

func Equal[T comparable](t testing.TB, have T, want T, desc string) {
	t.Helper()
	internal.Equal(t.Fatalf, t, have, want, desc)
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
		t.Fatalf("%s:\nhave: %v\nwant: true", desc, pred)
	}
}

func False(t testing.TB, pred bool, desc string) {
	t.Helper()
	if pred {
		t.Fatalf("%s:\nhave: %v\nwant: false", desc, pred)
	}
}

func Contains(t testing.TB, haystack, needle string, desc string) {
	t.Helper()
	internal.Contains(t.Fatalf, t, haystack, needle, desc)
}

func NoError(t testing.TB, err error, desc string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s:\nhave: %v (%T)\nwant: <no error>", desc, err, err)
	}
}

func ErrorContains(t testing.TB, err error, want string, desc string) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s:\nhave: <no error>\nwant: <an error>", desc)
	}
	Contains(t, err.Error(), want, desc)
}

func ErrorMatches(t testing.TB, have error, wantPattern string, desc string) {
	t.Helper()
	if have == nil {
		t.Fatalf("%s:\nhave: <no error>\nwant: error matching: %s", desc, wantPattern)
	}
	ok, mErr := regexp.MatchString(wantPattern, have.Error())
	if mErr != nil {
		t.Fatalf("%s:\nregexp build error: %v", desc, mErr)
	}
	if !ok {
		t.Fatalf("%s:\nerror message: %s\ndoes not match pattern: %s", desc, have, wantPattern)
	}
}

func ErrorIs(t testing.TB, have, want error, desc string) {
	t.Helper()
	if have == nil {
		t.Fatalf("%s:\nhave: <no error>\nwant: %v (%T)", desc, want, want)
	}
	if !errors.Is(have, want) {
		t.Fatalf("%s:\nhave: %v (%T)\nwant: %v (%T)", desc, have, have, want, want)
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
