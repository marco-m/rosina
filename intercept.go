package rosina

import (
	"io"
	"os"
	"testing"
)

// InterceptOutput replaces what 'file' points to (normally os.Stdout) with a
// pipe and returns the 'readReset' function. The 'readReset' function returns
// what has been written to 'file' and resets 'file' to point to the original.
//
// Example:
//
//	// Replace what is pointed to by os.Stdout (note the &):
//	read := rosina.InterceptOutput(t, &os.Stdout)
//	// Here call the SUT, something that writes to os.Stdout.
//	// Once done, call read() and validate the output:
//	out := read()
func InterceptOutput(t *testing.T, file **os.File) func() string {
	t.Helper()
	rd, wr, err := os.Pipe()
	if err != nil {
		t.Fatalf("creating pipe: %s", err)
	}
	orig := *file
	*file = wr

	return func() string {
		t.Helper()
		wr.Close()
		out, err := io.ReadAll(rd)
		if err != nil {
			t.Fatalf("reading from pipe: %s", err)
		}
		*file = orig
		return string(out)
	}
}
