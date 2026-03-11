# PIOS Model Context Protocol (MCP) Server

PIOS ships with a native MCP server inside the main CLI binary.

Instead of asking an agent to shell out, parse text, and infer state, MCP lets the agent call the PIOS contract tools directly over JSON-RPC.

## Why Use MCP

Without MCP, an agent typically has to:
- run `pios status`
- read the terminal output
- run `pios validate`
- decide when it is safe to move forward

That still works, but MCP is cleaner for tools that support it. The agent gets structured access to the same contract logic without depending on shell parsing.

## Start the Server

Run:

```bash
pios mcp
```

This starts the stdio MCP server exposed by the main `pios` binary.

## Available Tools

### `pios_status`
Reads the repository state from `STATUS.md` and `templates/tasks.md` and returns the current gate plus task counts.

### `pios_validate`
Validates the current milestone against the PIOS contract. It fails if required planning artifacts are missing or if unfinished tasks remain.

### `pios_init`
Seeds the current directory with the PIOS contract files. It also supports the optional `ide` argument for `cursor`, `windsurf`, or `claude` scaffolding.

### `pios_next`
Transitions the repo to the next milestone by archiving completed state, sweeping completed tasks, and resetting `STATUS.md` to planning mode.

`pios_next` does not create the next spec, plan, or task list by itself. After the transition, either the human or the agent still needs to draft those artifacts.

## Recommended Use

### Human-Gated Default

Use MCP when you want the agent to stay tightly aligned with the repo contract while you still review each milestone.

Typical loop:
1. The agent calls `pios_status` before coding.
2. It works only on tasks marked `[ ]`.
3. It calls `pios_validate` when the milestone looks complete.
4. You review the result.
5. You or the agent calls `pios_next` only when the milestone is truly done.

### Experimental Autonomous Loop

For bounded experiments, the agent can own the transition too.

```text
When `pios_validate` passes:
1. Call `pios_next`.
2. Draft the next `templates/spec-lock.md`, `templates/plan-lock.md`, and `templates/tasks.md`.
3. Update `STATUS.md` for the next milestone.
4. Do not resume implementation until those artifacts are written.
5. Stop if the next milestone requires a product-direction decision or human risk review.
```

## Setup Examples

### Claude Code

```bash
claude mcp add pios-mcp -- pios mcp
```

### Cursor

Add a shell command MCP server in Cursor with the command:

```bash
pios mcp
```

### Windsurf

If your Windsurf environment supports MCP servers, attach PIOS with the command:

```bash
pios mcp
```

## Fallback Path

If your tool does not support MCP, use the normal CLI path instead. The contract behavior is the same. MCP just removes the need for shell-output parsing.
