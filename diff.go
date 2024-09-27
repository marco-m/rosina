package rosina

import (
	"github.com/google/go-cmp/cmp"
	"github.com/marco-m/rosina/internal/diff"
)

// AnyDiff returns a textual representation of the differences between 'have' and
// 'want'. Usage:
//
//	if delta := diff(body, tc.wantBody); delta != "" {
//		t.Fatalf("get %s: body: mismatch. +have -want:\n%s",
//		tc.urlPath, delta)
//	}
//
// Unfortunately I observed cmp.AnyDiff to be unstable: it randomly returns either
// tabs or spaces on the exact same inputs. This is normally OK but makes flaky
// tests that compare the output of cmp.AnyDiff :-(
func AnyDiff[T any](have, want T) string {
	return cmp.Diff(want, have)
}

// DiffLenient is based solely on the untyped output of fmt.Print, but should
// be OK since we wrap it with generics to force the same type. Introduced to
// see if we can solve the problems we have with diffUnstable.
// func DiffLenient[T any](have, want T) string {
// 	return pretty.Compare(want, have)
// }

// TextDiff returns a unified diff.
// Based on a copy of internal/diff from the Go stdlib
// https://github.com/golang/go/tree/master/src/internal/diff
func TextDiff(oldName string, old []byte, newName string, new []byte) []byte {
	return diff.Diff(oldName, old, newName, new)
}
