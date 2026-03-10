# Claude Adapter (PIOS)

Best for:
- long-form reasoning
- structured Markdown artifacts
- large blueprint creation

Recommended pattern:
- Produce Plan Lock + Tasks first.
- Generate artifacts as clean Markdown “file blocks” (path + content).
- Use strict phase gates to reduce drift.

Avoid:
- micro-prompts that fragment context

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Always check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.

> **MCP Server Architecture:** PIOS features a native MCP (Model Context Protocol) server. If you are operating within an MCP-capable environment (like Claude Code), you can invoke `pios_validate()`, `pios_status()`, and `pios_next()` natively via JSON-RPC without relying on terminal execution.
