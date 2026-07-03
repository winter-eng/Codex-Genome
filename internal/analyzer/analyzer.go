package analyzer

import (
	"sort"
	"time"

	"github.com/codexgenome/codex-genome/internal/language"
	"github.com/codexgenome/codex-genome/internal/scanner"
)

// Report contains the aggregated metrics derived from a scan.
type Report struct {
	RootPath        string
	TotalFiles      int
	TotalDirs       int
	Extensions      []ExtensionCount
	Languages       []LanguageStat
	PrimaryLanguage string
	TotalLanguages  int
	Duration        time.Duration
}

// ExtensionCount pairs a file extension with the number of files using it.
type ExtensionCount struct {
	Ext   string
	Count int
}

// LanguageStat describes a detected language and its share of the project.
type LanguageStat struct {
	Language   string
	FileCount  int
	Percentage float64 // share of total project files
}

// Analyze processes a scanner.Project and returns aggregated metrics.
func Analyze(project *scanner.Project) *Report {
	langs, primary, totalLangs := computeLanguages(project.Files)
	return &Report{
		RootPath:        project.RootPath,
		TotalFiles:      len(project.Files),
		TotalDirs:       len(project.Directories),
		Extensions:      computeExtensions(project.Files),
		Languages:       langs,
		PrimaryLanguage: primary,
		TotalLanguages:  totalLangs,
	}
}

func computeExtensions(files []scanner.File) []ExtensionCount {
	counts := make(map[string]int)
	for _, f := range files {
		ext := f.Extension
		if ext == "" {
			ext = "(no extension)"
		}
		counts[ext]++
	}

	exts := make([]ExtensionCount, 0, len(counts))
	for ext, count := range counts {
		exts = append(exts, ExtensionCount{Ext: ext, Count: count})
	}
	sort.Slice(exts, func(i, j int) bool {
		if exts[i].Count != exts[j].Count {
			return exts[i].Count > exts[j].Count
		}
		return exts[i].Ext < exts[j].Ext
	})
	return exts
}

func computeLanguages(files []scanner.File) (stats []LanguageStat, primary string, total int) {
	counts := make(map[string]int)
	for _, f := range files {
		counts[language.Detect(f.Extension)]++
	}

	totalFiles := len(files)
	stats = make([]LanguageStat, 0, len(counts))
	for lang, count := range counts {
		pct := 0.0
		if totalFiles > 0 {
			pct = float64(count) / float64(totalFiles) * 100
		}
		stats = append(stats, LanguageStat{
			Language:   lang,
			FileCount:  count,
			Percentage: pct,
		})
	}

	sort.Slice(stats, func(i, j int) bool {
		if stats[i].FileCount != stats[j].FileCount {
			return stats[i].FileCount > stats[j].FileCount
		}
		return stats[i].Language < stats[j].Language
	})

	// Primary language: first recognized entry after sorting by count.
	for _, s := range stats {
		if s.Language != "Unknown" {
			primary = s.Language
			break
		}
	}

	// Total: distinct recognized languages only.
	for lang := range counts {
		if lang != "Unknown" {
			total++
		}
	}

	return stats, primary, total
}
