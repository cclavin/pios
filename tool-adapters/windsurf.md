# Windsurf Adapter (PIOS)

Windsurf supports project rules (commonly `.windsurfrules`) depending on setup.

Recommended:
- Keep `AGENTS.md` + `profiles/base/*` as the core guidance.
- Use PIOS phases explicitly to reduce back-and-forth.
- Generate scaffolds file-by-file for reliability.

> **MCP Server Integration:** PIOS features a live MCP server. If your Windsurf environment supports the Model Context Protocol, you can attach PIOS by setting its command to `pios mcp`. This grants the agent access to validation and continuous loop tools.

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
