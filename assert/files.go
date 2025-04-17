package assert

import (
	"os"
	"strings"
	"testing"
)

// FileEqualsString asserts that the contents of file 'fPath' are equal to
// string 'want'.
func FileEqualsString(t testing.TB, fPath string, want string) {
	t.Helper()
	have, err := os.ReadFile(fPath)
	if err != nil {
		t.Fatalf("reading fPath: %s", err)
	}
	Equal(t, string(have), want, fPath)
}

// FileEqualsFile asserts that the contents of file 'fpath' are equal to
// the contents of file 'wantPath'.
func FileEqualsFile(t testing.TB, fPath string, wantPath string) {
	t.Helper()
	have, err := os.ReadFile(fPath)
	if err != nil {
		t.Fatalf("reading fPath: %s", err)
	}
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("reading wantPath: %s", err)
	}
	Equal(t, string(have), string(want), "fPath and wantPath")
}

// FileContains asserts that string 'want' is contained in file 'fPath'.
func FileContains(t testing.TB, fPath string, want string) {
	t.Helper()
	have, err := os.ReadFile(fPath)
	if err != nil {
		t.Fatalf("reading fPath: %s", err)
	}
	haystack := string(have)

	if !strings.Contains(haystack, want) {
		// Heuristic to avoid printing a potentially huge file.
		if len(haystack) < 1024 {
			t.Fatalf("\nsubstring not found in file %s:\nsubstring: %q\nstring:    %q",
				fPath, want, haystack)
		}
		t.Fatalf("\nsubstring not found in file %s:\nsubstring: %q",
			fPath, want)
	}
}
