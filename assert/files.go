package assert

import (
	"os"
	"strings"
	"testing"
)

// AssertFileEqualsString asserts that the contents of file 'fPath' are equal to
// string 'want'.
func AssertFileEqualsString(t testing.TB, fPath string, want string) {
	t.Helper()
	have, err := os.ReadFile(fPath)
	if err != nil {
		t.Fatalf("reading fPath: %s", err)
	}
	Equal(t, string(have), want, fPath)
}

// AssertFileEqualsFile asserts that the contents of file 'fpath' are equal to
// the contents of file 'wantPath'.
func AssertFileEqualsFile(t testing.TB, fPath string, wantPath string) {
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

// AssertFileContains asserts that string 'want' is contained in file 'fPath'.
func AssertFileContains(t testing.TB, fPath string, want string) {
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
