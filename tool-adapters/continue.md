# Continue Adapter (PIOS)

Continue supports rules and tool customization.

Recommended:
- Keep standards in `profiles/base/`
- Export selective rules to Continue rules folder (optional future exporter)
- Use tasks with acceptance criteria to drive incremental implementation

## Core PIOS Instructions
1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
