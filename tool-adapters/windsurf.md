# Windsurf Adapter (PIOS)

Windsurf supports project rules (commonly `.windsurfrules`) depending on setup.

Recommended:
- Keep `AGENTS.md` + `profiles/base/*` as the core guidance.
- Use PIOS phases explicitly to reduce back-and-forth.
- Generate scaffolds file-by-file for reliability.

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
