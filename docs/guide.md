# Getting Started with PIOS

PIOS (Project Input/Output System) is not a wrapper, an LLM, or a runtime. It is a **strict execution contract** designed to force modern AI agents to behave deterministically.

## The Problem
When using powerful agents like Claude Code or Cursor, it is incredibly easy for the agent to:
1. Start coding before the architecture is actually thought out.
2. Fall into endless loops trying to fix one bug while breaking another.
3. Hallucinate new files or dependencies that you did not ask for.

## The Solution
PIOS forces the AI into a rigid waterfall structure known as **Phase Gates**. 

Before the AI is permitted to write a single line of application code, it must:
1. Formally lock the project scope (`spec-lock.md`).
2. Formally lock the architectural plan (`plan-lock.md`).
3. Formally list out every single task required to build the plan (`tasks.md`).

Once it begins coding, it must check off every single task in `tasks.md`. It cannot move to a new phase until all tasks are complete, and you (the human) validate the gate.

## Installation

### The Zero-To-Hero Agent Prompt
If you use a background agent (like Claude Code) and don't want to type anything, run the agent in an empty folder and paste this prompt:

> "First, check if Go is installed on my system. If not, figure out the best way to install it silently for my OS. Once Go is installed, install the PIOS cli globally via `go install github.com/cclavin/pios/cmd/pios@latest`. Next, create a new directory for this project, enter it, and run `pios init`. After initialization, read the `AGENTS.md` file to understand the contract."

### Manual CLI Installation
If you prefer to install it yourself, ensure Go 1.22+ is installed, then run:

```bash
go install github.com/cclavin/pios/cmd/pios@latest
```

Then, initialize a new project folder:
```bash
mkdir my-app && cd my-app
pios init
```

*Note: If you use Cursor or Windsurf, run `pios init --ide=cursor` to instantly import the native PIOS rules into your IDE.*
