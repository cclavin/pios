# Getting Started with PIOS

PIOS (Project Input/Output System) is not a wrapper, an LLM, or a runtime. It is a strict execution contract designed to force modern AI agents to behave deterministically.

## The Problem

When using powerful agents like Claude Code, Cursor, or Codex, it is easy for the agent to:
1. Start coding before the architecture is actually thought out.
2. Fall into loops trying to fix one bug while breaking another.
3. Hallucinate files, dependencies, or milestones that you did not ask for.

## The Solution

PIOS forces the workflow into explicit phase gates.

Before the agent is permitted to write application code, it should:
1. Lock the project scope in `templates/spec-lock.md`.
2. Lock the implementation plan in `templates/plan-lock.md`.
3. Lock the task board in `templates/tasks.md`.

Once implementation starts, the agent works only from the active tasks, validates the gate, and moves forward only when the contract is satisfied.

## Installation

### Preferred Install Path

PIOS is distributed via native package managers:

**macOS / Linux (Homebrew):**
```bash
brew tap cclavin/tap
brew install pios
```

**Windows (Winget):**
```powershell
winget install cclavin.pios
```

**Any OS (Go Fallback):**
```bash
go install github.com/cclavin/pios/cmd/pios@latest
```

### Initialize a Project

```bash
mkdir my-app && cd my-app
pios init --ide=cursor
```

Supported `--ide` values are `cursor`, `windsurf`, and `claude`.

### MCP Setup

If your tool supports MCP, connect PIOS through:

```bash
pios mcp
```

That gives the agent native access to `pios_status`, `pios_validate`, `pios_init`, and `pios_next`.

## Recommended Usage

### Default Path: MCP or CLI With Human-Gated Milestones

This is the safest and most reliable workflow.

1. Create or retrofit the repo with `pios init`.
2. Fill out `templates/min-spec.md`, then lock the spec, plan, and tasks.
3. Let the agent implement against the active tasks only.
4. Use `pios status` or `pios_status` to check context.
5. Use `pios validate` or `pios_validate` to close the milestone.
6. Run `pios next` only after you decide the milestone is complete.

### Zero-To-Hero Bootstrap Prompt

If you want a shell-capable agent to bootstrap everything in one shot, use a prompt like this:

> "First, install PIOS using the best available method for my OS. Prefer a native package-manager install using `brew tap cclavin/tap && brew install pios` for macOS/Linux, or `winget install cclavin.pios` for Windows. If neither is available, fallback to `go install github.com/cclavin/pios/cmd/pios@latest`. Next, create a new directory for this project, enter it, and run `pios init`. After initialization, read `AGENTS.md` to understand the contract, then proceed through the PIOS phases."

Treat this as a bootstrap recipe, not the default daily workflow.

### Experimental Autonomous Continuation

If you want to see how far an agent can go on its own, keep the project bounded and use explicit loop rules.

```text
You are operating under the PIOS execution contract.

Goal:
<one bounded project goal>

Loop rules:
1. Call `pios_status` or run `pios status` at the start of each cycle.
2. Only work on tasks marked `[ ]` in `templates/tasks.md`.
3. Mark the active task `[/]`, complete it, verify it, then mark it `[x]`.
4. When a milestone is complete, run `pios_validate`.
5. If validation passes, run `pios_next`.
6. Immediately draft the next `templates/spec-lock.md`, `templates/plan-lock.md`, and `templates/tasks.md`, and update `STATUS.md` for the next milestone.
7. Do not start implementation for the new milestone until those artifacts are updated.
8. Continue automatically only if the next milestone still fits the original project goal.
9. Stop and summarize if blocked, if product direction is unclear, or if security or deployment risk needs human review.
```

## Tool-Specific Adapters

For tool-specific setup, see:
- [Claude](https://github.com/cclavin/PIOS/blob/main/tool-adapters/claude.md)
- [Cursor](https://github.com/cclavin/PIOS/blob/main/tool-adapters/cursor.md)
- [Windsurf](https://github.com/cclavin/PIOS/blob/main/tool-adapters/windsurf.md)
- [Codex](https://github.com/cclavin/PIOS/blob/main/tool-adapters/codex.md)
- [Continue](https://github.com/cclavin/PIOS/blob/main/tool-adapters/continue.md)
- [OpenClaw](https://github.com/cclavin/PIOS/blob/main/tool-adapters/openclaw.md)
