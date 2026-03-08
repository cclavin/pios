# OpenClaw Adapter (PIOS)

> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

Best for:
- Autonomous background execution
- Terminal-driven coding workflows
- Iterative implement-and-test loops

Recommended pattern:
1. **Always check state before executing:** Run `pios status` to determine the current project phase and remaining tasks.
2. **Strict Scope:** Only work on tasks marked `[ ]` in the current phase. Do not invent new tasks unless explicitly requested by the user.
3. **Task Checkout:** When beginning a task, update `templates/tasks.md` to mark it as `[/]` (In Progress).
4. **Task Completion:** When finishing a task, satisfy all acceptance criteria, then update `templates/tasks.md` to mark it as `[x]` (Done).
5. **Phase Validation:** Before assuming the phase is complete, run `pios validate`. If it fails, fix the missing criteria.
