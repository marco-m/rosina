package check

import (
	"testing"

	"github.com/marco-m/rosina/internal"
)

func Equal[T comparable](t testing.TB, have T, want T, desc string) {
	t.Helper()
	internal.Equal(t.Errorf, t, have, want, desc)
}

func Contains(t testing.TB, haystack, needle string, desc string) {
	t.Helper()
	internal.Contains(t.Errorf, t, haystack, needle, desc)
}
