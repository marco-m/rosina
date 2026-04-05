package golden

import (
	"io"
	"os"
	"testing"

	"github.com/marco-m/rosina/assert"
	"github.com/marco-m/rosina/diff"
)

// DiffTextAndFile returns a text diff between the contents of the file at 'goldenPath'
// and the string 'have' (if 'update' is false).
// If 'update' is true, DiffTextAndFile overwrites the file at 'goldenPath' with string
// 'have' and returns an empty diff.
func DiffTextAndFile(t *testing.T, have string, goldenPath string, update bool) string {
	t.Helper()
	haveStr := have
	wantStr := ReadOrUpdate(t, haveStr, goldenPath, update)
	return string(diff.TextDiffPatient("want", []byte(wantStr), "have", []byte(haveStr)))
}

// DiffFiles returns a text diff between the contents of the file at 'goldenpath' and the
// contents of the file at 'havePath' (if 'update' is false).
// If 'update' is true, DiffFiles overwrites the file at 'goldenPath' with file at
// 'havePath' and returns an empty diff.
func DiffFiles(t *testing.T, havePath string, goldenPath string, update bool) string {
	t.Helper()
	have, err := os.ReadFile(havePath)
	assert.NoError(t, err, "reading file")
	haveStr := string(have)
	wantStr := ReadOrUpdate(t, haveStr, goldenPath, update)
	return string(diff.TextDiffPatient("want", []byte(wantStr), "have", []byte(haveStr)))
}

// ReadOrUpdate returns the contents of the file at 'goldenPath' (if 'update' is false).
// If 'update' is true, ReadOrUpdate overwrites the file at 'goldenPath' with string
// 'have' and returns 'have'.
func ReadOrUpdate(t *testing.T, have string, goldenPath string, update bool) string {
	t.Helper()
	fi, err := os.OpenFile(goldenPath, os.O_RDWR, 0o644)
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			t.Fatalf("closing file: %s", err)
		}
	}()

	if update {
		t.Logf("updating golden file: %s", goldenPath)
		if _, err := fi.WriteString(have); err != nil {
			t.Fatalf("writing file %s: %s", goldenPath, err)
		}
		return have
	}

	content, err := io.ReadAll(fi)
	if err != nil {
		t.Fatalf("reading file %s: %s", goldenPath, err)
	}
	return string(content)
}
