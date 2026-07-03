# ◈ Codex Genome

A command-line tool for analyzing the structure and composition of software projects.

---

## What it does

Codex Genome scans a directory and produces a structured report: file counts, directory layout, extension breakdown, and language distribution. The analysis pipeline is designed to grow incrementally — each engine is independent and composable without touching existing code.

## Why it exists

Most tools that inspect a project either do too much (full IDE-level analysis) or too little (a few shell one-liners). Codex Genome occupies the middle ground: a focused, inspectable pipeline that starts with filesystem discovery and can be extended engine by engine.

## Features

- Recursive filesystem scanning with configurable ignore rules
- File extension breakdown with proportional bar chart
- Language detection for 30+ languages and formats
- In-memory project tree representation
- Professional terminal output
- Cross-platform: Linux, macOS, Windows

## Installation

**Build from source** (requires Go 1.23+)

```bash
git clone https://github.com/codexgenome/codex-genome
cd codex-genome
go build -o codex-genome .
```

**Using go install**

```bash
go install github.com/codexgenome/codex-genome@latest
```

## Usage

```bash
# Analyze the current directory
codex-genome analyze .

# Analyze a specific path
codex-genome analyze /path/to/project
codex-genome analyze C:\Projects\MyApp
```

## Example output

```
 ◈ Codex Genome  /home/user/projects/myapp
──────────────────────────────────────────────────────────────

  Total Files              62
  Total Directories         8

──────────────────────────────────────────────────────────────

  File Extensions

  .go                   42  █████████████░░░░░░░
  .ts                    8  ██░░░░░░░░░░░░░░░░░░
  .json                  5  █░░░░░░░░░░░░░░░░░░░
  .md                    4  █░░░░░░░░░░░░░░░░░░░
  .yaml                  3  ░░░░░░░░░░░░░░░░░░░░

──────────────────────────────────────────────────────────────

  Language Profile

  Primary Language      Go

  Go                    42  █████████████░░░░░░░  67.7%
  TypeScript             8  ██░░░░░░░░░░░░░░░░░░  12.9%
  JSON                   5  █░░░░░░░░░░░░░░░░░░░   8.1%
  Markdown               4  █░░░░░░░░░░░░░░░░░░░   6.5%
  YAML                   3  ░░░░░░░░░░░░░░░░░░░░   4.8%

  Total Languages        5

──────────────────────────────────────────────────────────────
  Completed in    4ms
```

## Architecture

The codebase is organized as a pipeline of single-responsibility packages:

```
cmd/
  root.go          Cobra root command
  analyze.go       analyze subcommand and orchestration

internal/
  filesystem/      Path resolution and validation
  scanner/         Recursive filesystem discovery, domain models
  tree/            In-memory project tree construction
  language/        Extension → language name mapping
  analyzer/        Metric aggregation from scan results
  report/          Terminal output rendering

pkg/               Reserved for future exportable libraries
docs/              Reserved for extended documentation
examples/          Reserved for example integrations
```

Data flows in one direction:

```
scanner → tree        (independent, not yet wired to render)
scanner → analyzer → report
```

No package imports from a layer above it. No circular dependencies.

## Ignored directories

The scanner skips the following by default:

`.git` · `node_modules` · `vendor` · `dist` · `build`

## Roadmap

- [ ] Gitignore-aware scanning
- [ ] Lines of code per language
- [ ] Dependency graph extraction (imports, modules)
- [ ] Tree view rendering
- [ ] JSON and CSV output formats
- [ ] Project health scoring
- [ ] Watch mode

## Development philosophy

- Each PR delivers one complete, working thing.
- The scanner only discovers. Analyzers only aggregate. The renderer only renders.
- No speculative code. No TODO comments in source files.
- Standard library preferred over third-party dependencies.
- If something needs an interface to justify it, it probably does not need an interface.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE).
