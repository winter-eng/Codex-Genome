package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/codexgenome/codex-genome/internal/filesystem"
)

// ignoredDirs lists directory names excluded from every scan.
var ignoredDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	"dist":         true,
	"build":        true,
}

// Scan walks the directory tree rooted at path and returns a populated Project.
// Entries are visited in deterministic lexicographic order (guaranteed by filepath.WalkDir).
// Permission errors on individual entries are skipped gracefully; the walk continues.
func Scan(path string) (*Project, error) {
	abs, err := filesystem.Abs(path)
	if err != nil {
		return nil, err
	}

	isDir, err := filesystem.IsDir(abs)
	if err != nil {
		return nil, err
	}
	if !isDir {
		return nil, fmt.Errorf("%q is not a directory", abs)
	}

	project := &Project{
		RootPath:      abs,
		ScanStartedAt: time.Now(),
	}

	err = filepath.WalkDir(abs, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // skip inaccessible entries; continue the walk
		}

		if d.IsDir() {
			if ignoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			if p == abs {
				return nil // root is the scan target, not a discovered directory
			}
			project.Directories = append(project.Directories, newDirectory(abs, p, d.Name()))
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil // skip files whose metadata is unreadable
		}

		project.Files = append(project.Files, newFile(abs, p, d.Name(), info))
		return nil
	})

	project.ScanFinishedAt = time.Now()
	project.Duration = project.ScanFinishedAt.Sub(project.ScanStartedAt)

	if err != nil {
		return nil, fmt.Errorf("scanning %q: %w", abs, err)
	}

	return project, nil
}

func newFile(root, abs, name string, info fs.FileInfo) File {
	rel, _ := filepath.Rel(root, abs)
	return File{
		Name:         name,
		AbsolutePath: abs,
		RelativePath: rel,
		Extension:    filepath.Ext(name),
		SizeBytes:    info.Size(),
		LastModified: info.ModTime(),
		Depth:        strings.Count(rel, string(filepath.Separator)),
	}
}

func newDirectory(root, abs, name string) Directory {
	rel, _ := filepath.Rel(root, abs)
	return Directory{
		Name:         name,
		AbsolutePath: abs,
		RelativePath: rel,
		Depth:        strings.Count(rel, string(filepath.Separator)) + 1,
	}
}
