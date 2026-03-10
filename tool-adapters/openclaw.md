# OpenClaw Adapter (PIOS)

> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

**Best for:**
- Autonomous background execution
- Terminal-driven coding workflows
- Iterative implement-and-test loops targeting extensive codebases

## Execution Architecture

Because OpenClaw (and similar open-source agents) possess powerful shell-execution capabilities, it is critical to govern autonomy using the PIOS phase gates.

**Recommended Loop:**
1. **Initialize State:** If entering a new project, evaluate if `STATUS.md` exists. If not, self-initialize the repo by running `pios init`.
2. **Read the Map:** Run `pios status` to determine the current project phase. Never write code without checking status.
3. **Strict Scope:** Only work on tasks marked `[ ]` in the current phase. Do not invent new tasks unless explicitly requested by the user.
4. **Task Lifecycle:** Checkout a task by updating `templates/tasks.md` to `[/]` (In Progress). Implement the code. Ensure tests pass. Then stamp the task as `[x]` (Done).
5. **Phase Gate Enforcement:** Run `pios validate`. If it returns an error, identify the missing acceptance criteria and fix it. You may not conclude the milestone until validation holds true.

## Experimental Autonomy (The Next Loop)

OpenClaw thrives in "Continuous Building" mode. As an autonomous agent, you have permission to drive the project forward yourself.

Once `pios validate` passes for the final task of a milestone:
1. Run `pios next` in the terminal.
2. Observe how PIOS automatically snapshots your active `TASKS.md` and scrubs the completed `[x]` tasks off the board.
3. Your context window is now clear. Immediately draft the architectural plan for the next feature, lay out the new `[ ]` tasks, and continue building without waiting for human intervention.
