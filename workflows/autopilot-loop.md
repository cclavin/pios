# Autopilot Loop

This document outlines the standard operation flow for a background AI agent (e.g., Claude Code, native Cursor loop, Windsurf Cascade) operating entirely inside a PIOS repository.

## The Loop
An agent should repeatedly loop through these steps until the `current_gate` allows them to move forward, or until `pending_tasks` hits 0.

### 1. Check Context & Status
The agent reads the overarching state:
```bash
pios status
```
*The JSON output defines exactly what the agent should currently care about.*

### 2. Pick the Next Task
The agent searches `TASKS.md` for the first sequence matching `### [ ] TASK-`. 
It updates the file to mark it in progress:
The agent should use its native file-editing capabilities to locate the target task in `templates/tasks.md`, change the `[ ]` checkbox to `[/]`, and update its status to `In Progress`.

### 3. Implement & Test
The agent performs its native coding cycle (viewing files, editing files, running local commands). It must enforce that the acceptance criteria defined in the task block are met.

### 4. Conclude the Task
Once the task is verified complete, the agent marks it done.
The agent should use its native file-editing capabilities to update `templates/tasks.md`, changing `[/]` to `[x]` for the completed task and updating its status to `Done`.

### 5. Validate Phase Gate
The agent checks if the gate is complete.
```bash
pios validate
```
If it succeeds, the agent can manually update `STATUS.md` to point to the next gate, and loop back to step 1. If it fails, the CLI will output what specific criteria is missing.
