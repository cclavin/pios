# GPT Adapter (PIOS)

Best for:
- balanced architecture + implementation
- structured artifacts
- iterative refinement

Recommended pattern:
- Start with Spec Lock + Plan Lock generation.
- Then generate files in batches by directory.
- Keep a running STATUS.md and decision log.

Avoid:
- generating entire repo in one huge message if tool cannot write files directly

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

1. **Status Check:** Run `pios status` to check the current project phase before starting work.
2. **Task Scope:** Only check out and work on tasks marked `[ ]` in `templates/tasks.md`. Update them to `[/]` while in progress, and `[x]` when completed.
3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.
