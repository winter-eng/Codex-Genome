# Contributing to Codex Genome

Thank you for taking the time to contribute.

This document covers everything you need to get started: branching, coding standards, commit conventions, and the PR process.

---

## Prerequisites

- Go 1.23 or later
- Git

---

## Getting started

```bash
git clone https://github.com/codexgenome/codex-genome
cd codex-genome
go mod tidy
go build ./...
go test ./...
```

---

## Branching strategy

`main` is always stable. Never commit directly to `main`.

Branch names follow this pattern:

| Prefix | Use |
|---|---|
| `feat/` | New features |
| `fix/` | Bug fixes |
| `docs/` | Documentation only |
| `refactor/` | Code changes with no behavior change |
| `ci/` | CI and tooling changes |
| `chore/` | Maintenance (deps, cleanup) |

Examples:

```
feat/json-output
fix/scanner-permission-error
docs/contributing-guide
```

---

## Commit convention

This project follows [Conventional Commits](https://www.conventionalcommits.org/).

```
<type>: <short description>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `refactor`, `test`, `ci`, `chore`

Examples:

```
feat: add JSON output format to report package
fix: skip unreadable entries during filesystem walk
docs: document language detection architecture
refactor: extract extension computation to helper function
ci: add Windows to test matrix
```

Breaking changes add `!` after the type:

```
feat!: change Scan return type from Result to Project
```

---

## Coding standards

- Run `gofmt -w .` before every commit. Unformatted code will fail CI.
- Run `go vet ./...` and resolve all warnings.
- Write idiomatic Go. Follow the conventions at [go.dev/doc/effective_go](https://go.dev/doc/effective_go).
- Keep packages focused on a single responsibility.
- No interfaces unless multiple concrete types already exist that need them.
- No global mutable state.
- Comments on exported symbols only. Comments should explain *why*, not *what*.

---

## Architecture rules

The pipeline flows in one direction:

```
filesystem → scanner → tree / analyzer → report
```

- The **scanner** only discovers files. It does not analyze content.
- The **analyzer** only aggregates scanner output. It does not touch the filesystem.
- The **report** only renders. It does not compute metrics.
- The **language** package is stateless. It maps extensions to names and nothing else.
- New features go in new packages when possible. Extending existing packages should require a clear reason.

---

## Running the full check locally

```bash
gofmt -l .          # should print nothing
go vet ./...        # should print nothing
go test ./...
go build ./...
```

---

## Submitting a pull request

1. Fork the repository and create a branch from `main`.
2. Make your changes following the standards above.
3. Verify the full check passes locally.
4. Open a PR against `main` using the PR template.
5. A maintainer will review within a few days.

Keep PRs small and focused. One purpose per PR.
If you are unsure whether a feature fits the project, open an issue first.

---

## Reporting issues

Use the GitHub issue templates:

- **Bug Report** — unexpected behavior, crashes, wrong output
- **Feature Request** — new analysis engines, output formats, CLI flags

Please search existing issues before opening a new one.
