package rosina

import (
	"fmt"
	"os"

	"github.com/go-quicktest/qt"
)

// FileEqualsString returns a qt.Checker checking equality of the contents of
// file fpath with string want.
func FileEqualsString(fpath string, want string) qt.Checker {
	return &fileEqualsStringChecker{fpath, want}
}

type fileEqualsStringChecker struct {
	fpath string
	want  string
}

func (fc *fileEqualsStringChecker) Args() []qt.Arg {
	return []qt.Arg{{"fpath", fc.fpath}, {"want", fc.want}}
}

func (fc *fileEqualsStringChecker) Check(note func(key string, value any)) error {
	got, err := os.ReadFile(fc.fpath)
	if err != nil {
		return err
	}
	return qt.Equals(string(got), fc.want).Check(note)
}

// FileEqualsFile returns a qt.Checker checking equality of the contents of
// file gotPath with file wantPath.
func FileEqualsFile(gotPath string, wantPath string) qt.Checker {
	return &fileEqualsFileChecker{gotPath, wantPath}
}

type fileEqualsFileChecker struct {
	gotPath  string
	wantPath string
}

func (fc *fileEqualsFileChecker) Args() []qt.Arg {
	return []qt.Arg{{"gotPath", fc.gotPath}, {"wantPath", fc.wantPath}}
}

func (fc *fileEqualsFileChecker) Check(note func(key string, value any)) error {
	got, err := os.ReadFile(fc.gotPath)
	if err != nil {
		return fmt.Errorf("reading gotPath: %w", err)
	}
	want, err := os.ReadFile(fc.wantPath)
	if err != nil {
		return fmt.Errorf("reading wantPath: %w", err)
	}

	return qt.Equals(string(got), string(want)).Check(note)
}

// FileContains returns a qt.Checker checking whether string want is contained
// in the file at fpath.
func FileContains(fpath string, want string) qt.Checker {
	return &fileContainsChecker{fpath, want}
}

type fileContainsChecker struct {
	fpath string
	want  string
}

func (fc *fileContainsChecker) Args() []qt.Arg {
	return []qt.Arg{{"fpath", fc.fpath}, {"want", fc.want}}
}

func (fc *fileContainsChecker) Check(note func(key string, value any)) error {
	got, err := os.ReadFile(fc.fpath)
	if err != nil {
		return err
	}
	return qt.StringContains(string(got), fc.want).Check(note)
}
