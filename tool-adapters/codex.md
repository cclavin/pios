# Codex Adapter (PIOS)

Best for:
- direct code changes
- incremental commits
- tight implement-test loops

Recommended pattern:
- Feed Codex the Plan Lock + Tasks.
- Execute one task at a time.
- Require tests or smoke checks per task.
- Use autopilot loop: implement → test → fix → doc → summarize

Avoid:
- expecting deep architecture without Plan Lock

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
