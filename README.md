# π PIOS - AI Project Execution Contract

[![Version](https://img.shields.io/github/v/release/cclavin/PIOS?color=orange&label=version)](https://github.com/cclavin/PIOS/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/cclavin/pios)](https://goreportcard.com/report/github.com/cclavin/PIOS)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<p align="center">
  <img src="pios-banner.gif" alt="PIOS - Contracts over vibes. Build useful things, faster." />
</p>

## Table of Contents
- [Mission](#mission)
- [Core Positioning](#core-positioning)
- [Repository Layout](#repository-layout)
- [Quick Start](#quick-start)
- [Example Prompts](#example-prompts)
- [Backtesting PIOS](#backtesting-pios)
- [Continuing the Loop](#continuing-the-loop-post-completion)
- [Related Ecosystems](#related-ecosystems--methodologies)
- [Roadmap](#roadmap)

## Mission

PIOS is built on the belief that human attention is our most valuable resource. The current generation of AI tools often leads to endless chat sessions that waste time, compute, and energy without producing tangible results. PIOS exists to return **agency and focus** to developers. By enforcing deterministic boundaries and execution contracts, PIOS ensures that AI serves as a direct catalyst for human creativity rather than a conversational sinkhole, helping you build useful things, faster, and with less frustration.

---

## Core Positioning

PIOS is **artifact-first**: it produces repo files, phase gates, and repeatable workflows, not just chat.

- **Completion-first:** PIOS operates on **phase gates** with explicit exit criteria. The CLI strictly enforces both artifact existence (specs must exist before coding) and checklist completion (Minimum Spec → Spec Lock → Plan Lock → Task Lock → Scaffold Done).
- **Contract-first:** PIOS is not a competing "rules format." It establishes a source-of-truth contract. Maintain core guidance in PIOS via `AGENTS.md` (tool-agnostic baseline) and tool-specific adapters. This keeps your repo consistent even as tools change.
- **Artifact-first:** Focuses strictly on files, diffs, commands, and checklists. Minimal prose.
- **Framework & Runtime Agnostic:** PIOS doesn't care if you use Cursor, Copilot, Windsurf, or a background Claude/OpenAI agent. It serves as the immutable data layer that any AI can read to understand exactly what to do next.

---

## Core Philosophy

PIOS exists to convert AI assistance into **repeatable outcomes**.

1. **Contracts over vibes:** Start with structured inputs (specs) that can be refined without restarting.
2. **Artifacts over chat:** Produce files, diffs, commands, and checklists.
3. **Phase gates:** Each phase ends only when exit criteria are met.
4. **Tool-agnostic by default:** Build a source of truth and export to tool-specific formats.
5. **Completion-first:** PIOS optimizes for finishing, not perfect planning.
6. **Autopilot loops:** Default workflow is implement → test → fix → document → summarize.
7. **Decisions are logged:** Avoid re-litigating choices.

---

## Repository Layout

- `cmd/pios/` — the Golang PIOS validator CLI
- `templates/` — reusable project artifacts (specs, plans, tasks, decision logs) safely embedded into the CLI
- `agents/` — role-specific agent instructions
- `tool-adapters/` — guidance for using PIOS with specific tools
- `profiles/` — shared standards + stack templates
- `workflows/` — consolidated phase-by-phase operating flows and commands

---

## Tutorials & Workflows

PIOS can be used in five distinct ways, depending on how deeply you want to integrate it into your AI developer environment. Choose the workflow that best fits your style.

### 1. The Agent-Native Path (MCP Server) - *Highly Recommended*

This is the most powerful way to use PIOS. The CLI incorporates a fully native Model Context Protocol (MCP) server. Once connected, your AI agent reads the current tasks and validates its own phase gates over a background JSON-RPC connection—you never have to type `pios validate` in the terminal!

**The Setup:**
*   **For Claude Code:** Run `claude mcp add pios-mcp -- pios mcp`
*   **For Cursor:** Open Settings > Features > MCP. Click **+ Add new MCP server**. Set Name to `pios`, Type to `command`, and Command to `pios mcp`.

**The Daily Workflow:**
1.  Run `pios init` in a new folder to drop the project templates.
2.  Fill out your `templates/min-spec.md`.
3.  Prompt your connected agent (e.g., *"Review my PIOS spec and begin compiling the plan-lock."*).
4.  As the agent writes code, it will automatically call the `pios_validate` tool to check its own work before moving to the next milestone!

### 2. The Command-Line Path (Human-in-the-loop)

If you prefer to manually control the phase gates while the AI writes the code, the native Golang CLI instantly injects templates and tracks the repository state via strict terminal commands.

**Install the CLI:**
```bash
go install github.com/cclavin/pios/cmd/pios@latest
```

**Initialize a new project:**
```bash
pios init --ide=cursor
```
This drops the templates and scaffolds the IDE context rules inline. Your daily workflow simply consists of letting the AI build, and you manually typing `pios validate` to ensure the contract is met before checking off the milestone.

### 3. The Zero-to-Hero Path (Fully Autonomous)

If you have a powerful agent (like a strong Windsurf cascade) and don't want to touch the terminal at all, you can give your AI a single "Zero-to-Hero" prompt that commands it to install the CLI locally, initialize the context, and start building in one shot.

**Example Prompt:**
> "First, check if Go is installed on my system. If not, best-effort install it silently for my OS. Once Go is installed, install the PIOS cli globally via `go install github.com/cclavin/pios/cmd/pios@latest`. 
> 
> Next, create a new directory for this project, enter it, and run `pios init`. After initialization, read the `AGENTS.md` file to understand the contract. Finally, proceed through the PIOS phases to build me a python script that scrapes hacker news."

### 4. The Clone & Run Path (CLI/IDE Native)

If you just want the rule framework without installing Go or running a global binary on your machine, simply clone the repository into your new project folder. This gives your AI the `AGENTS.md` context without any system dependencies.

```bash
# Clone the repository
git clone https://github.com/cclavin/pios.git my-new-project
cd my-new-project

# Remove the .git folder to start fresh
rm -rf .git
```

### 5. Open-Source Agents (OpenHands, Cline, OpenClaw)

If you are running local open-source agents inside sandboxed containers or local environments, PIOS provides the perfect structure.
1. Command the agent to install PIOS: `go install github.com/cclavin/pios/cmd/pios@latest`
2. Command them to run `pios init` in the workspace.
3. Because these agents possess shell execution capabilities, they will naturally use `pios validate` and `pios next` to govern their own loops, guaranteeing they don't hallucinate out of bounds.

### 6. The Manual / Creative Path (Framework Agnostic)

If you are just having a chat on the ChatGPT web interface, PIOS is still highly effective.
1. Manually copy `STATUS.md` and the Markdown files in the `/templates/` directory into your project's `/docs` folder.
2. Fill out the specs.
3. Pass the completed `tasks.md` to ChatGPT along with the strict prompt: *"You are operating under the PIOS execution contract. Read `AGENTS.md`. Only work on tasks marked `[ ]`."*
4. Without the CLI, *you* are the manual phase gate validator! Ensure the AI respects the checklist.

---

## Example Prompts

A good PIOS prompt focuses on the *what* and the *constraints*, leaving the *how* to the AI within the contract framework.

**Good Prompt (Concise & Constrained):**
> "Initialize a new project using the PIOS templates. Create a vanilla JS widget that fetches the current weather for a user's location. It must be a single file, styled with inline CSS, and handle permissions gracefully."

**Good Prompt (Next Milestone Update):**
> "We've completed Milestone 1. Please update the `plan-lock.md` to include an auth layer using Supabase, then wipe the current `TASKS.md` and generate Milestone 2 tasks for the login and registration flows."

---

## Backtesting PIOS

PIOS includes a methodology to ensure it doesn't become a "vibe framework." 
Compare your baseline workflow vs. PIOS on 2–3 small projects to validate efficiency.

Track the following metrics:
- Time-to-scaffold (minutes)
- Clarification turns
- Rework rate / context resets
- Finish rate (0 or 1)

### v1.0 Release Benchmarks

To validate the PIOS execution contract model outside of simple HTML environments, the following benchmarks were designed as strict "Zero Human Intervention" tests targeting multi-file architectures and package managers. Each project went from a blank directory to a validated, running build under total agent autonomy.

> **Environment:** Claude Sonnet 4.6 (Claude Code) · `pios` CLI · Zero Human Clarifications

| Benchmark | Time ⏱️ | Clarification Turns | Rework Events | Context Resets | Finish Rate | Artifacts |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: |
| **Robust Backend Constraint (Go + SQLite)** | 15m | 0 | 0 | 0 | 100% | [view full report](runs/v1-backend/) |
| **User-Like Frontend (React/NextJS)** | 10m | 0 | 2 (Self-Resolved) | 0 | 100% | [view full report](runs/v1-frontend/) |

> **Scope note:** Pre-release v0.4 frontend benchmarks are available in the [`runs/archived-tests/`](runs/archived-tests/) directory.

#### 1. Robust Backend Constraint (Go 1.22 + SQLite)
- **Prompt:** A strict, highly constrained prompt demanding a Go 1.22 `net/http` REST API with persistent SQLite bindings via `modernc.org/sqlite` (no CGO), covered entirely by tests and packaged in a multi-stage `Dockerfile`.
- **Outcome:** Flawless semantic execution. The PIOS contract prevented the agent from defaulting to popular but prohibited frameworks (like Gin or GORM). It correctly scoped the architecture, resolved internal state bugs via test suites, and pushed a production-grade multi-stage container. **10/10 passing tests. 0 human interventions.**
- **Anomaly Report (The Power of Precision):** The agent generated dockerfiles, tests, and source code perfectly, but *failed* to write a `README.md` for the generated project. Why? Because the prompt did not explicitly ask for one, and `AGENTS.md` does not strictly mandate one to pass Phase 4. Rather than a downside, this is the core value proposition of PIOS: **It prevents LLM hallucinations.** The agent builds *exactly* what is in the spec-lock, maintaining total discipline. Nothing more, nothing less. If you want a README, you spec a README.

#### 2. User-Like Frontend (React + Next.js App Router)
- **Prompt:** A loose, generalized prompt simply asking for a "nice looking dashboard... React, Next.js, Tailwind, dark mode, glassmorphism."
- **Outcome:** A stress-test proving PIOS prevents code vomit on unstructured prompts. The project locked the layout and component structure in Phase 3 *before* building. During scaffolding, the agent encountered directory conflicts with `create-next-app` and validator read errors on template checkboxes. Because PIOS execution defines specific checkpoints, the agent recognized the defect, debugged its own path, and resolved it in-flight without breaking contract. **Turbopack build succeeded. 0 human interventions.**

---

## The Continuous Building Loop (`pios next`)

Once your AI agent finishes the initial milestone and all tasks are marked `[x]`, the project is not dead. In fact, this is where PIOS shines.

There are two primary ways to operate between milestones depending on your comfort level:

**1. Normal Sprints (Human-in-the-Loop)**
For standard Agile-style development, you act as the product manager. When a milestone finishes, you pause the AI, review the code, and mentally verify the business logic. Then, you manually run `pios next` to execute the system snapshot, after which you write the specs for the next sprint and hand it back to the AI.

**2. Experimental Autonomy**
If you want to see how far an agent can go on its own, PIOS enables continuous, unprompted building. An agent connected to the MCP Server (or a powerful open-source agent with CLI access) can finish its tasks and instantly invoke `pios_next` by itself. It will archive its own work, generate its own roadmap for the next milestone, and start coding again—indefinitely.

When you are ready for the next feature (or when your agent decides it's ready), simply run:
```bash
pios next
```

This instantly executes the PIOS Continuous Loop protocol:
1. **Snapshot:** It creates a timestamped archive (`templates/archive/YYYY-MM.../`) of your completed `TASKS.md` and `STATUS.md` so you never lose the history of your technical decisions.
2. **Sweep:** It scrubs all `[x]` checked tasks from your active tasks file, leaving behind a clean board containing any rolled-over pending items.
3. **Reset:** It resets your `STATUS.md` phase gates back to planning mode.

If you are connected via the **MCP Server**, your AI agent can even run `pios_next` by itself! You can simply prompt: *"Great job on the auth layer. Run the next command, then let's draft Milestone 2 focusing on the database adapter."*

---

## Related Ecosystems & Methodologies

PIOS pairs exceptionally well with system-level instruction formatting. For users looking to standardize their global AI behavior, tools, and customized context, I highly recommend exploring [Daniel Miessler's Personal AI Infrastructure](https://github.com/danielmiessler/Personal_AI_Infrastructure). 

By combining a robust personal AI infrastructure (to define your developer identity and global rules) with PIOS (to enforce project-specific execution and finishing), you create a highly deterministic, end-to-end AI development capability.

---

## Documentation Website

PIOS includes a dedicated documentation website built with VitePress (located in the `docs` folder). To run the documentation locally:

```bash
cd docs
npm install
npm run docs:dev
```
Then open `http://localhost:5173` in your browser.

---

## Roadmap

- [x] v0.1 - templates, agents, adapters, workflows, backtest harness
- [x] v0.2 - machine-readable state
- [x] v0.3 - Golang CLI: `pios init / validate / status`
- [x] v0.4 - Contract Hardening & Backtesting
- [x] v0.5 - Universal Context Scaffolding (`--ide`) & ASCII Easter Eggs
- [x] v0.6 - Goreleaser Native Distribution & VitePress Documentation Site
- [x] v1.0 - Model Context Protocol (MCP) Server Integration & Stable Release Matrix
