package diff

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/marco-m/rosina/diff/internal/diff"
)

// AnyDiff returns a textual representation of the differences between 'have'
// and 'want'.
func AnyDiff[T any](have, want T) string {
	// I observed cmp.Diff to be unstable: it randomly returns either tabs or
	// spaces on the exact same inputs. This is actually done on purpose. On the
	// other hand, we do want the output to be as stable as possible. We attempt
	// to make it stable by replacing tabs with spaces.
	return strings.ReplaceAll(cmp.Diff(want, have), "\t", "    ")
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
