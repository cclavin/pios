# Cursor Adapter (PIOS)

Cursor supports rule files and structured context.

Recommended:
- Keep PIOS as source-of-truth.
- Export concise rules to `.cursorrules` (or `.cursor/rules/`).
- Pin `AGENTS.md` in context.

Workflow:
- Architect phase in chat
- Scaffold phase as repo edits
- Auditor phase as review + PR-style notes

## Core PIOS Instructions
1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
