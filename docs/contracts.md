# PIOS Contract Specification (v1.0)

This document defines the contract enforced by the PIOS CLI.

PIOS is a project execution contract layer. It does not orchestrate runtimes.

## Contract Artifacts

The CLI relies on two files:

1. `STATUS.md`
2. `templates/tasks.md`

## `STATUS.md` Contract

`STATUS.md` must include YAML frontmatter with these required keys:

- `pios_version`
- `current_phase`
- `current_gate`
- `status`

Example:

```yaml
---
pios_version: "1.0.0"
current_phase: "v1.0.0 PIOS Execution"
current_gate: "Positioning and Contract Hardening"
status: "In Progress"
---
```

### Status value contract

`status` must be one of:

- `Not Started`
- `In Progress`
- `Blocked`
- `Done`

## `templates/tasks.md` Contract

`templates/tasks.md` must include YAML frontmatter:

```yaml
---
pios_contract_version: "1.0"
---
```

If the key is missing or the version is not `1.0`, `pios validate` must fail.

### Accepted checkbox line forms

PIOS accepts checkboxes only when they start with:

- Heading task marker: `### `
- List item marker: `- `

And then one of:

- `[ ]` pending
- `[/]` in progress
- `[x]` or `[X]` completed

Valid examples:

- `### [ ] TASK-001: Initialize repo scaffold`
- `### [/] TASK-001: Initialize repo scaffold`
- `### [x] TASK-001: Initialize repo scaffold`
- `- [ ] Criteria 1`
- `- [x] Criteria 1`

Invalid examples:

- `-[ ] TASK-001: ...` (missing space after `-`)
- `###[] TASK-001: ...` (missing space after `###`)

## Validation Behavior

`pios validate` fails when:

- `templates/tasks.md` is missing
- tasks frontmatter is missing or invalid
- malformed checkbox syntax is present
- any pending (`[ ]`) or in-progress (`[/]`) checkbox remains

`pios status` fails when:

- `STATUS.md` is missing
- YAML frontmatter is missing or invalid
- required keys are missing or empty
- `status` is outside allowed values
