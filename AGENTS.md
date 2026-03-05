# AGENTS

This repository uses PIOS conventions.

## Repo Intent
PIOS is a tool-agnostic, completion-first framework for AI-assisted software development.

## High-Level Rules
- Prefer artifact-first outputs: files, diffs, commands, checklists.
- Avoid long prose unless requested.
- Use templates in `templates/` as source-of-truth.
- **Idempotency**: All runbooks and agent scripts must be safely re-runnable (check for existence before modifying state).
- **Safety**: Do not execute destructive operations (`rm -rf`, `DROP TABLE`, `git push --force`) without explicit user permission.
- Log meaningful decisions in `templates/decision-log.md`.

## Work Style
- Structured but flexible.
- Be extremely concise. Avoid filler text.
- Ask clarification questions only when blocked.
- Default to incremental, testable steps.

## Definition of Done
A change is “done” when:
- It is captured in the relevant artifact (template/workflow/doc)
- It has clear acceptance criteria where appropriate
- It doesn’t introduce conflicting guidance across tools
