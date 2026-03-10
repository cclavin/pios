# Cursor Adapter (PIOS)

Cursor supports rule files and structured context.

Recommended:
- Keep PIOS as source-of-truth.
- Export concise rules to `.cursorrules` (or `.cursor/rules/`).
- Pin `AGENTS.md` in context.

> **MCP Server Integration:** PIOS features a live MCP server. In Cursor (Settings > Features > MCP), add a new shell command server named `PIOS` with the command `pios mcp`. This gives Cursor native access to `pios_status`, `pios_validate`, and `pios_next`.

Workflow:
- Architect phase in chat
- Scaffold phase as repo edits
- Auditor phase as review + PR-style notes

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
