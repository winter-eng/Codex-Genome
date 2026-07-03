package scanner

import "time"

// Project is the top-level model returned by a completed scan.
// It represents everything the scanner discovered about a directory tree.
type Project struct {
	RootPath       string
	Files          []File
	Directories    []Directory
	ScanStartedAt  time.Time
	ScanFinishedAt time.Time
	Duration       time.Duration
}

// File represents a single file discovered during a scan.
type File struct {
	Name         string
	AbsolutePath string
	RelativePath string
	Extension    string // includes the leading dot, e.g. ".go"; empty if none
	SizeBytes    int64
	LastModified time.Time
	Depth        int // 0 = directly inside the scanned root
}

// Directory represents a directory discovered during a scan.
// The root itself is not included; only its descendants are.
type Directory struct {
	Name         string
	AbsolutePath string
	RelativePath string
	Depth        int // 1 = directly inside the scanned root
}
