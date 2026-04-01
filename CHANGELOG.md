# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.0] - 2026-03-31
### Added
- **MCP Server Integration:** Native JSON-RPC support for Claude Code, Cursor, and Windsurf through the `pios mcp` command.
- **Package Manager Distribution:** Official formulas for Homebrew (macOS/Linux) and Winget (Windows) via Goreleaser.
- **VitePress Documentation:** Launched comprehensive [documentation site](https://cclavin.github.io/pios/) covering workflows and tool adapters.
- **Validated Backtesting:** Hardened the contract against complex Go and React multi-file architectures.

## [v0.5.0] - 2026-03-08
### Added
- `--ide` context scaffolding logic to `pios init` (supports `cursor`, `windsurf`, `claude`).
- Secret ASCII easter eggs (PIOS color-gradient banner and cat command).

## [v0.4.0] - 2026-03-08
### Added
- Completed formal v1.0 backtesting (Go Backend and React/NextJS Frontend).
- Archived pre-release HTML5 benchmarks.
- Overhauled README with hard metrics, badges, and automated Zero-to-Hero onboarding prompts.
- Created positioning documents defining PIOS as an execution contract (`docs/positioning.md`, `docs/scope.md`).
- Locked contract grammar definitions (`docs/contracts.md`).

### Changed
- Hardened CLI validator matching logic in `cmd/pios/main.go`.
- Relocated exporters from v0.4 Roadmap to Deferred.
- Updated tool adapters and workflows to remove runtime orchestration framing.

## [v0.3.1] - 2026-03-05
### Fixed
- Fixed directory traversal logic in `pios status` and `pios validate` by implementing `findProjectRoot()`.
- Replaced brittle string matching with robust RegEx parsing for `TASKS.md` completion checks.
- Aligned `tool-adapters/*.md` with `CONTRIBUTING.md` by explicitly appending required `pios` instructions for CI loops.
- Added `main_test.go` test harness.
- Configured GitHub Actions CI pipeline.

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
