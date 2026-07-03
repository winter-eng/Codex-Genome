package language

import "strings"

// extMap maps file extensions (lowercase, including the leading dot) to language names.
// A single language may map from multiple extensions (e.g. ".yaml" and ".yml" → "YAML").
var extMap = map[string]string{
	// Go
	".go": "Go",

	// TypeScript / JavaScript
	".ts":  "TypeScript",
	".tsx": "TSX",
	".js":  "JavaScript",
	".jsx": "JSX",

	// Data / Config
	".json":     "JSON",
	".yaml":     "YAML",
	".yml":      "YAML",
	".toml":     "TOML",
	".xml":      "XML",

	// Markup
	".md":       "Markdown",
	".markdown": "Markdown",
	".html":     "HTML",
	".htm":      "HTML",

	// Styles
	".css":  "CSS",
	".scss": "SCSS",
	".sass": "SASS",

	// Shell
	".sh":   "Shell",
	".bash": "Shell",
	".ps1":  "PowerShell",

	// Systems
	".py":  "Python",
	".rs":  "Rust",
	".c":   "C",
	".h":   "C",
	".cpp": "C++",
	".cc":  "C++",
	".cxx": "C++",
	".hpp": "C++",

	// JVM / CLR
	".cs":   "C#",
	".java": "Java",
	".kt":   "Kotlin",
	".kts":  "Kotlin",

	// Mobile / Scripting
	".swift": "Swift",
	".php":   "PHP",
	".rb":    "Ruby",

	// Containers
	".dockerfile": "Dockerfile",
}

// Detect returns the language name for the given file extension.
// ext must include the leading dot (e.g. ".go"). Matching is case-insensitive.
// Returns "Unknown" for unrecognized or empty extensions.
func Detect(ext string) string {
	if lang, ok := extMap[strings.ToLower(ext)]; ok {
		return lang
	}
	return "Unknown"
}
