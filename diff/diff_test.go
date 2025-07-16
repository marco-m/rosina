package diff_test

import (
	"testing"

	"github.com/marco-m/rosina/diff"
)

func TestTextDiff(t *testing.T) {
	test := func(old, new, want string) {
		t.Helper()
		have := diff.TextDiff("old", "new", old, new)
		if have != want {
			t.Errorf("\n===have===\n%s\n===want===\n%s", have, want)
		}
	}

	test("", "", "")
	test("ciccio", "ciccio", "")
	test(
		"ciccio",
		"bello",
		`--- old
+++ new
@@ -1 +1 @@
-ciccio
\ No newline at end of file
+bello
\ No newline at end of file
`)
	test(
		"ciccio\n",
		"bello\n",
		`--- old
+++ new
@@ -1 +1 @@
-ciccio
+bello
`)
	test(
		"1 aaa\n2 bbb\n3 ccc\n",
		"1 aaa\n2 xxx\n3 ccc\n",
		`--- old
+++ new
@@ -1,3 +1,3 @@
 1 aaa
-2 bbb
+2 xxx
 3 ccc
`)
}
