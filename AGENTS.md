# AGENTS

This repository uses PIOS conventions.

## Repo Intent
PIOS is a tool-agnostic, completion-first execution contract layer for AI-assisted software development.

## High-Level Rules
- Prefer artifact-first outputs: files, diffs, commands, checklists.
- Avoid long prose unless requested.
- Use templates in `templates/` as source-of-truth.
- **Idempotency**: All runbooks and agent scripts must be safely re-runnable (check for existence before modifying state).
- **Safety**: Do not execute destructive operations (`rm -rf`, `DROP TABLE`, `git push --force`) without explicit user permission.
- Log meaningful decisions in `templates/decision-log.md`.

## Lifecycle Maintenance
Agents must proactively keep project state documents synchronized:
- **STATUS.md**: Update the YAML frontmatter (`current_phase`, `current_gate`, `status`) immediately when a phase gate is crossed or a milestone is achieved.
- **CHANGELOG.md**: Document any notable features, fixes, or architectural pivots under an `[Unreleased]` header as soon as the code is committed.
- **ROADMAP.md**: Strike through or check off `[x]` roadmap items when the underlying EPIC or release version is completed.

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
