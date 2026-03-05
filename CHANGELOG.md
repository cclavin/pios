# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.3.0] - 2026-03-05
### Added
- Go Standard Project Layout: `go.mod` initialized at repository root.
- Template Embedding: Added `templates/embed.go` to natively package markdown artifacts into the CLI binary using `//go:embed`.
- `cmd/pios/`: Re-architected the PIOS validation CLI as a cross-platform Golang binary.
  - `pios init`: Recursively copies embedded PIOS templates into a user's working directory.
  - `pios status`: Parses `STATUS.md` utilizing robust YAML AST parsing via `gopkg.in/yaml.v3` to output JSON payloads.
  - `pios validate`: Programmatically scans `templates/tasks.md` ensuring strict phase gate completion.

### Removed
- Permanently deprecated and removed the bash-based `bin/pios` prototype script.

## [v0.2.0] - 2026-03-05
### Added
- Machine-readable YAML frontmatter to `STATUS.md` and `templates/status-template.md`.
- Strict machine-parseable markdown checkbox format to `TASKS.md`.
- Minimal Bash validation script (`bin/pios`) allowing `status` and `validate`.
- `workflows/autopilot-loop.md` to define background agent operation within PIOS.

## [v0.1.0] - 2026-03-05
### Added
- Initial framework scaffold.
- Root repository files: `README.md`, `LICENSE`, `AGENTS.md`, `ROADMAP.md`, `STATUS.md`.
- Core Templates: `min-spec.md`, `spec-lock.md`, `plan-lock.md`, `tasks.md`, `status-template.md`, `decision-log.md`.
- Agents: `architect`, `scaffold`, `auditor`.
- Tool Adapters: `claude.md`, `cursor.md`, `windsurf.md`, `gpt.md`, `continue.md`.
- Profiles: `base` standards and `stacks` (Go, Python, React).
- Workflows: Consolidated `guide.md`.
- Runs: Backtesting documentation and template.
