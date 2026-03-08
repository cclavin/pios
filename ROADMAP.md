# ROADMAP

## v0.1 (Completed)
- Core templates (min-spec/spec-lock/plan-lock/tasks/decision-log)
- 3 agent roles (architect/scaffold/auditor)
- Tool adapters (claude/cursor/windsurf)
- Workflows docs (scaffold/audit)
- Backtesting guidance

## v0.2 (Completed)
- Gate validation checklist implementation
- Machine-readable YAML state tracking
- Autopilot loop logic

## v0.3 (Completed)
- Golang Validation CLI (`pios init`, `pios validate`, `pios status`)
- Embed `templates/` natively into the `go` binary

## v0.5.0 (Current)
- **Universal Context Scaffolding:** Add `--ide=<cursor|windsurf|claude>` flags to `pios init` to auto-generate agent rules files (`.cursorrules`, `CLAUDE.md`, etc) directly into repositories.
- **CLI Easter Eggs:** Add aesthetic ASCII art to reward CLI usage (secret, unannounced).

## v0.6.0 (Upcoming)
- **VitePress/Nextra Documentation Site:** Comprehensive web documentation hosting guides, workflows, and adapter setups.
- **Native Distribution:** Availability via package managers (`brew install pios`, Winget) to remove the `go install` dependency for non-Go developers.

## v1.0.0 (Release Candidate)
- **Model Context Protocol (MCP) Server Integration:** Native tool integration allowing AI agents (Claude Code, Cursor) to invoke `pios_validate()` via JSON-RPC, bypassing terminal bash execution.
- **Advanced Backtesting Suite:** Introduce multi-paradigm, non-HTML5 benchmark tests (e.g. Go CLI tools, fullstack React/Python web apps) to validate PIOS agent discipline on complex file structures.
- **Stable Spec Schema:** Lock the `pios_contract_version` structure in YAML headers.
- **Repeatable Release Process:** Fully automated semantic releases.

## Deferred / Later
- Interop exporters to transpile `profiles/` into Cursor/Continue/Windsurf `.rules` formats.
- Integrations: hooks, commit templates
