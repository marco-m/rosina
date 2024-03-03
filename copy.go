package rosina

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies srcPath to dstPath.
func CopyFile(dstPath string, srcPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("copyfile: src: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("copyfile: dst: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("copyfile: copy: %w", err)
	}
	return nil
}
