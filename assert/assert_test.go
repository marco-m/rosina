package assert_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/marco-m/rosina/assert"
)

func TestEqualStringsOneLine(t *testing.T) {
	assertPass(t, "IdenticalStrings", func(t testing.TB) {
		assert.Equal(t, "hello", "hello", "ciccio")
	})

	want := `
ciccio mismatch:
--- want
+++ have
@@ -1 +1 @@
-goodbye
\ No newline at end of file
+hello
\ No newline at end of file
`
	assertFail(t, "DifferentStrings", want, func(t testing.TB) {
		assert.Equal(t, "hello", "goodbye", "ciccio")
	})
}

func TestEqualStringsMultiLine(t *testing.T) {
	assertPass(t, "IdenticalStrings", func(t testing.TB) {
		assert.Equal(t, "hello\nciccio", "hello\nciccio", "ciccio")
	})

	want := `
ciccio mismatch:
--- want
+++ have
@@ -1,3 +1,3 @@
 line 1
-foo
+line 2
 line 3
`

	innerHave := `line 1
line 2
line 3
`

	innerWant := `line 1
foo
line 3
`

	assertFail(t, "DifferentStrings", want, func(t testing.TB) {
		assert.Equal(t, innerHave, innerWant, "ciccio")
	})
}

func TestEqualNumbers(t *testing.T) {
	assertPass(t, "IdenticalNumbers", func(t testing.TB) {
		assert.Equal(t, 42, 42, "ciccio")
	})

	want := `
ciccio mismatch:
have: 42
want: 3.14`
	assertFail(t, "DifferentNumbers", want, func(t testing.TB) {
		assert.Equal(t, 42, 3.14, "ciccio")
	})
}

func TestDeepEqual(t *testing.T) {
	type Zoo struct {
		X int
		Y string
		Z string
	}

	assertPass(t, "IdenticalStructs", func(t testing.TB) {
		assert.DeepEqual(t, Zoo{}, Zoo{}, "zoo")
	})

	outerWant := `
zoo mismatch: +have -want:
  assert_test.Zoo{
-     X: 2,
+     X: 1,
      Y: "same",
-     Z: "want",
+     Z: "have",
  }
`
	innerHave := Zoo{
		X: 1,
		Y: "same",
		Z: "have",
	}
	innerWant := Zoo{
		X: 2,
		Y: "same",
		Z: "want",
	}
	assertFail(t, "DifferentStructs", outerWant, func(t testing.TB) {
		assert.DeepEqual(t, innerHave, innerWant, "zoo")
	})
}

func TestAssertContains(t *testing.T) {
	haystack := "Nel mezzo del cammin di nostra vita"

	assertPass(t, "HaystackContains", func(t testing.TB) {
		assert.Contains(t, haystack, "mezzo del cammin")
	})

	want := `
substring not found in string:
substring: "una selva oscura"
string:    "Nel mezzo del cammin di nostra vita"`
	assertFail(t, "HaystackDoesNotContain", want, func(t testing.TB) {
		assert.Contains(t, haystack, "una selva oscura")
	})
}

//
//
//

// Trick (original name: testTester) taken from
// https://github.com/alecthomas/assert/blob/master/assert_test.go
type testSpy struct {
	*testing.T
	failed  bool
	failMsg string
}

func (t *testSpy) Fatalf(format string, args ...any) {
	t.failed = true
	t.failMsg = fmt.Sprintf(format, args...)
}

func (t *testSpy) Fatal(args ...any) {
	t.failed = true
	t.failMsg = fmt.Sprint(args...)
}

// Trick (original name: testTester) taken from
// https://github.com/alecthomas/assert/blob/master/assert_test.go
func assertFail(t *testing.T, name string, wantMsg string, fn func(t testing.TB)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Helper()
		spy := &testSpy{T: t}
		fn(spy)
		if !spy.failed {
			t.Fatalf("\nassertFail:\nwant: <fail>\nhave: <pass>")
		}
		spy.failMsg = fixCmpDiff(spy.failMsg)
		wantMsg = fixCmpDiff(wantMsg)
		if spy.failMsg != wantMsg {
			t.Fatalf("\nassertFail:\nhave:\n%s\nwant:\n%s\n",
				quote(spy.failMsg), quote(wantMsg))
		}
	})
}

// Trick (original name: testTester) taken from
// https://github.com/alecthomas/assert/blob/master/assert_test.go
func assertPass(t *testing.T, name string, fn func(t testing.TB)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Helper()
		spy := &testSpy{T: t}
		fn(spy)
		if spy.failed {
			t.Fatalf("\nassertPass:\nwant: <pass>\nhave:\n%s", quote(spy.failMsg))
		}
	})
}

func quote(msg string) string {
	var bld strings.Builder
	verbose := true
	for _, line := range strings.Split(msg, "\n") {
		if !verbose {
			fmt.Fprintf(&bld, "|%s\n", line)
			continue
		}
		fmt.Fprintf(&bld, "|%-40s|", line)
		for _, ru := range line {
			fmt.Fprintf(&bld, " %2x", ru)
		}
		fmt.Fprintln(&bld)
	}
	return bld.String()
}

func fixCmpDiff(s string) string {
	// randomly cmd.Diff returns NBS (non-breaking space) instead of space :-(
	sigh := func(r rune) rune {
		if r == 0xa0 {
			return ' '
		}
		return r
	}
	s = strings.Map(sigh, s)
	return strings.ReplaceAll(s, "\t", "    ")
}

func TestFixCmpDiff(t *testing.T) {
	have := fixCmpDiff("	")
	assert.Equal(t, have, "    ", "tabs")
	t.Log(quote("\t"))
}
