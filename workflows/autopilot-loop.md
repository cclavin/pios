# Autopilot Loop

This document outlines the standard operation flow for a background AI agent (e.g., Claude Code, native Cursor loop, Windsurf Cascade) operating entirely inside a PIOS repository.

## The Loop
An agent should repeatedly loop through these steps until the `current_gate` allows them to move forward, or until `pending_tasks` hits 0.

### 1. Check Context & Status
The agent reads the overarching state:
```bash
./bin/pios status
```
*The JSON output defines exactly what the agent should currently care about.*

### 2. Pick the Next Task
The agent searches `TASKS.md` for the first sequence matching `### [ ] TASK-`. 
It updates the file to mark it in progress:
```bash
# Example: Marking TASK-001 in progress
sed -i 's/### \[ \] TASK-001/### \[\/\] TASK-001/' templates/tasks.md
sed -i 's/- \*\*Status:\*\* Open/- \*\*Status:\*\* In Progress/' templates/tasks.md
```

### 3. Implement & Test
The agent performs its native coding cycle (viewing files, editing files, running local commands). It must enforce that the acceptance criteria defined in the task block are met.

### 4. Conclude the Task
Once the task is verified complete, the agent marks it done.
```bash
sed -i 's/### \[\/\] TASK-001/### \[x\] TASK-001/' templates/tasks.md
sed -i 's/- \*\*Status:\*\* In Progress/- \*\*Status:\*\* Done/' templates/tasks.md
```

### 5. Validate Phase Gate
The agent checks if the gate is complete.
```bash
./bin/pios validate
```
If it succeeds, the agent can manually update `STATUS.md` to point to the next gate, and loop back to step 1. If it fails, the CLI will output what specific criteria is missing.
