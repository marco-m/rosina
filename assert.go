package rosina

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertEqual[T comparable](t *testing.T, have T, want T, desc string) {
	t.Helper()
	if have != want {
		t.Fatalf("\n%s mismatch:\nhave: %v\nwant: %v\n", desc, have, want)
	}
}

func AssertDeepEqual[T any](t *testing.T, have T, want T, desc string) {
	t.Helper()
	if delta := diff(have, want); delta != "" {
		t.Fatalf("\n%s mismatch: +have -want:\n%s", desc, delta)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("\nhave: %s\nwant: <no error>", err)
	}
}

func AssertErrorIs(t *testing.T, err error, want error) {
	t.Helper()
	if !errors.Is(err, want) {
		t.Fatalf("\nhave: %s (%T)\nwant: %s (%T)", err, err, want, want)
	}
}

func AssertErrorTextEq(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("\nhave: <no error>\nwant: %s", want)
	}
	if delta := diff(err.Error(), want); delta != "" {
		t.Fatalf("\n%s mismatch: +have -want:\n%s", "error text mismatch", delta)
	}
}

// diff returns a textual representation of the differences between 'have' and
// 'want'. Usage:
//
//	if delta := diff(body, tc.wantBody); delta != "" {
//		t.Fatalf("get %s: body: mismatch. +have -want:\n%s",
//		tc.urlPath, delta)
//	}
func diff[T any](have, want T) string {
	return cmp.Diff(want, have)
}
