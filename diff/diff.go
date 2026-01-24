package diff

import (
	"strings"

	"github.com/alecthomas/repr"
	"github.com/google/go-cmp/cmp"
	"github.com/marco-m/rosina/diff/internal/diff"
	diffp "github.com/marco-m/rosina/diff/internal/diffp"
)

// ReprDiff returns a highly human-readable unified diff of the string representation
// of (have, want). This is probably the function you are looking for.
func ReprDiff[T any](have, want T) string {
	options := []repr.Option{repr.Indent("  "), repr.OmitEmpty(false), repr.OmitZero(false)}
	haveStr := repr.String(have, options...)
	wantStr := repr.String(want, options...)
	// TextDiff can have spurious diffs. TextDiffPatient seems better.
	return string(TextDiffPatient("want", []byte(wantStr), "have", []byte(haveStr)))
}

// TextDiff returns a unified diff of the strings (old, new).
// Based on a copy of x/tools/internal/diff
// https://github.com/golang/tools/tree/master/internal/diff
func TextDiff(oldLabel string, newLabel string, old string, new string) string {
	return diff.Unified(oldLabel, newLabel, old, new)
}

// TextDiffPatient returns a unified diff of the strings (old,new), using the patient
// diff algorithm.
// Note that "patient" comes from "patient sorting": it is actually _faster_ than a
// standard diff algorithm.
// Based on a copy of x/tools/internal/diffp
// https://github.com/golang/go/tree/master/src/internal/diffp
func TextDiffPatient(oldName string, old []byte, newName string, new []byte) []byte {
	return diffp.Diff(oldName, old, newName, new)
}

// AnyDiff returns a textual representation of the differences between 'have'
// and 'want'.
// Consider using [ReprDiff] instead, which has a better human-readable output.
func AnyDiff[T any](have, want T) string {
	// I observed cmp.Diff to be unstable: it randomly returns either tabs or
	// spaces on the exact same inputs. This is actually done on purpose. On the
	// other hand, we do want the output to be as stable as possible. We attempt
	// to make it stable by replacing tabs with spaces.
	return strings.ReplaceAll(cmp.Diff(want, have), "\t", "    ")
}
