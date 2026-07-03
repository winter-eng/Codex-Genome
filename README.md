# Codex Genome

A CLI tool for analyzing software project structure and composition.

## Installation

```bash
git clone https://github.com/codexgenome/codex-genome
cd codex-genome
go mod tidy
go build -o codex-genome .
```

## Usage

```bash
# Analyze current directory
codex-genome analyze .

# Analyze a specific path
codex-genome analyze C:\Projects\MyApp
codex-genome analyze /home/user/myproject
```

## Output

The `analyze` command produces a structured report showing:

- Total file and directory count
- File extension breakdown with a visual bar chart
- Execution time

Ignored directories: `.git`, `node_modules`, `vendor`, `dist`, `build`

## Requirements

- Go 1.23+

## License

MIT
# Codex-Genome
