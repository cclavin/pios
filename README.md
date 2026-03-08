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

PIOS is built on the belief that human attention is our most valuable resource. The current generation of AI tools often leads to endless chat sessions that waste time, compute, and energy without producing tangible results. PIOS exists to return **agency and focus** to developers. By enforcing deterministic boundaries and execution contracts, PIOS ensures that AI serves as a direct catalyst for human creativity rather than a conversational sinkhole - helping you build useful things, faster, and with less frustration.

---

## Core Positioning

PIOS is **artifact-first**: it produces repo files, phase gates, and repeatable workflows — not just chat.

- **Completion-first:** PIOS operates on **phase gates** with explicit exit criteria (Minimum Spec → Spec Lock → Plan Lock → Task Lock → Scaffold Done) and an **autopilot loop** (implement → test → fix → document → summarize).
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

## Quick Start

PIOS can be used programmatically via the built-in CLI, or manually by simply copying the contract templates and feeding them to your AI.

### The Automated Path (CLI)

The native Golang CLI instantly injects templates and tracks your AI agent's progress via strict validation gates.

**Install the CLI:**
```bash
go install github.com/cclavin/pios/cmd/pios@latest
```

**Initialize a new project:**
```bash
pios init
```
This will eject the templates and `STATUS.md` into your current directory. Follow the phase gates:

1. Fill out **Minimum Spec**: Keep it short. Avoid premature details.
2. Run **Spec Lock**: Resolve only the highest-impact unknowns.
3. Generate **Plan Lock**: Architecture, data flow, constraints, risks, and test strategy.
4. Convert plan → **TASKS.md**: Small tasks, each testable, each with acceptance criteria.
5. **Autopilot Loop**: Point an AI Agent at the repository. The agent can use `pios validate` and `pios status` to autonomously burn down `TASKS.md`.

### The Manual / Creative Path (No CLI)

If you don't use Go or prefer a lighter touch, PIOS is still highly effective as purely text-based prompting architecture.

1. **Copy the Templates:** Manually copy `STATUS.md` and the Markdown files in the `/templates/` directory into your project's root or a `/docs` folder.
2. **Set the Contract:** Fill out the specs just as you would with the CLI.
3. **Initialize the Agent:** Pass the completed `tasks.md` and `status-template.md` to Claude, ChatGPT, Cursor, or Windsurf. 
4. **The System Prompt:** The easiest way to kick off the agent is to give it the following explicit directive: *"You are operating under the PIOS execution contract. Read `AGENTS.md`. Only work on tasks marked `[ ]`. When you finish a task, check it off `[x]` and update `STATUS.md` before writing more code."*
5. **Human Validation:** Without the CLI, *you* are the phase gate validator! Review the agent's work and check its `STATUS.md` discipline before allowing it to proceed to the next milestone.

> **Coming Soon (v1.0 MCP):** In upcoming releases, the PIOS Go CLI will expose a native Model Context Protocol (MCP) server, allowing AI agents to invoke the validation tools natively via JSON-RPC instead of spawning sub-shells.

### The Fully Autonomous Path (Agent-Driven)

If you have a powerful agent (like Claude Code or a strong Windsurf cascade) and don't want to touch the terminal at all, you can give your AI a single "Zero-to-Hero" prompt that commands it to install PIOS, initialize the context, and start building in one shot.

**Example Prompt:**
> "First, check if Go is installed on my system. If not, figure out the best way to install it silently for my OS. Once Go is installed, install the PIOS cli globally via `go install github.com/cclavin/pios/cmd/pios@latest`. 
> 
> Next, create a new directory for this project, enter it, and run `pios init`. After initialization, read the `AGENTS.md` file to understand the contract. Finally, proceed through the PIOS phases to build me a python script that scrapes hacker news."

### The Clone & Run Path (CLI/IDE Native)

If you prefer not to install the PIOS Go CLI globally, you can simply clone the repository and run your AI agent directly inside it. This is highly recommended for **Cursor**, **Windsurf**, or **Claude Code** users.

```bash
# Clone the repository
git clone https://github.com/cclavin/pios.git my-new-project
cd my-new-project

# Remove the .git folder to start fresh
rm -rf .git
```

Once cloned, open the folder in your AI IDE. The included `AGENTS.md` file acts as the project's brain, immediately instructing your agent on how to use the contract system.

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

## Continuing the Loop (Post-Completion)

Once your AI agent finishes the initial sequence and all Phase 4 tasks are marked `[x]`, the project is not dead—the contract simply resets for the next milestone.

To seamlessly continue development:
1. **Archive the Log:** Move completed tasks to the bottom of `TASKS.md` or archive them to clear the Active deck.
2. **Draft the Next Milestone:** Use your agent to draft a highly detailed roadmap for the next feature set based on your minimal input (e.g., *"Great job on the MVP. Now, let's draft Milestone 2 focusing on user authentication."*). 
3. **Reset the Gate:** Ask the agent to establish a new `plan-lock.md` and generate a fresh task list with empty `[ ]` checkboxes.
4. **Use as a Skill:** If you are using Claude Code, Cursor, or a Personal AI Infrastructure (PAI), you can integrate PIOS workflows as a permanent "Skill." Direct your system prompt to *always* initialize the PIOS contract format when asked to "start a new module" or "plan a new feature," guaranteeing that your AI never writes code without a boundaried plan again.

---

## Related Ecosystems & Methodologies

PIOS pairs exceptionally well with system-level instruction formatting. For users looking to standardize their global AI behavior, tools, and customized context, I highly recommend exploring [Daniel Miessler's Personal AI Infrastructure](https://github.com/danielmiessler/Personal_AI_Infrastructure). 

By combining a robust personal AI infrastructure (to define your developer identity and global rules) with PIOS (to enforce project-specific execution and finishing), you create a highly deterministic, end-to-end AI development capability.

---

## Roadmap

- [x] v0.1 - templates, agents, adapters, workflows, backtest harness
- [x] v0.2 - machine-readable state
- [x] v0.3 - Golang CLI: `pios init / validate / status`
- [x] v0.4 - Contract Hardening & Backtesting
- [x] v0.5 - Universal Context Scaffolding (`--ide`) & ASCII Easter Eggs
- [x] v0.6 - Goreleaser Native Distribution & VitePress Documentation Site
- [ ] v1.0 - Model Context Protocol (MCP) Server Integration & Stable Release Matrix
