package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// Abs returns the absolute representation of path.
func Abs(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolving path %q: %w", path, err)
	}
	return abs, nil
}

// IsDir reports whether path exists and is a directory.
func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("path %q does not exist", path)
		}
		return false, fmt.Errorf("accessing %q: %w", path, err)
	}
	return info.IsDir(), nil
}
