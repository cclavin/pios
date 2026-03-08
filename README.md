# π PIOS — AI Project Execution Contract

[![Version](https://img.shields.io/badge/version-0.4.0-orange)](https://github.com/cclavin/PIOS/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/cclavin/PIOS)](https://github.com/cclavin/PIOS/commits/main)
[![GitHub issues](https://img.shields.io/github/issues/cclavin/PIOS)](https://github.com/cclavin/PIOS/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/cclavin/PIOS)](https://github.com/cclavin/PIOS/pulls)

<p align="center">
  <img src="pios-banner.gif" alt="PIOS — Contracts over vibes. Build useful things, faster." />
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

PIOS is built on the belief that human attention is our most valuable resource. The current generation of AI tools often leads to endless chat sessions that waste time, compute, and energy without producing tangible results. PIOS exists to return **agency and focus** to developers. By enforcing deterministic boundaries and execution contracts, PIOS ensures that AI serves as a direct catalyst for human creativity rather than a conversational sinkhole—helping you build useful things, faster, and with less frustration.

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
4. **The System Prompt:** Instruct your AI: *"You are operating under the PIOS execution contract. Read `AGENTS.md`. Only work on tasks marked `[ ]`. When you finish a task, check it off `[x]` and update `STATUS.md` before writing more code."*
5. **Human Validation:** Without the CLI, *you* are the phase gate validator! Review the agent's work and check its `STATUS.md` discipline before allowing it to proceed to the next milestone.

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

### v0.4.0 Pre-Release Benchmarks

To validate the PIOS execution contract model, the following three mini-projects were executed autonomously using the `pios` CLI strict phase gates. Each project went from a blank directory to validated completion in under 20 minutes.

> **Environment:** Gemini 3.1 (high reasoning) · Antigravity IDE · `pios` autopilot loop

| Benchmark | Time ⏱️ | Clarification Turns | Rework Events | Context Resets | Finish Rate | Artifacts |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: |
| **Dual-Game Menu Showcase** | 15m | 0 | 0 | 0 | 100% | [view run](runs/game-menu/) |
| **Physics Destruction Sim** | 5m | 0 | 0 | 0 | 100% | [view run](runs/destruction-sim/) |
| **Excuse Generator Widget** | 10m | 0 | 0 | 0 | 100% | [view run](runs/excuse-widget/) |

> **Scope note:** All three pre-release benchmarks are frontend/HTML5 projects — the fastest class for validating scaffolding speed and phase gate discipline. v1.0 benchmarks will include backend services, CLIs, and multi-file architectures.

#### 1. Dual-Game Menu & Mechanics Showcase
- **Prompt:** *"a game menu with two games built in and able to choose from where one the cat is just sleeping and it slightly reacts with user click or keyboard interaction. Another game in the menu that's brighter and more active, 2d platformer with slick ice physics and tuck down on keyboard (up down left right w a s d normal inputs) to build up speed on drop in altitude or downhill with test map of big hill and jump please, this must all be included in first run so do anything that's needed to test and complete"*
- **Outcome:** A cohesive HTML5 Canvas experience with a menu linking to a serene interactive cat toy and a high-speed momentum-based 2D platformer. Demonstrates the framework's ability to scaffold non-trivial logic accurately on the first run.

#### 2. Physics Destruction Simulator *(unstructured input stress test)*
- **Prompt:** *"Come up with a plan for a more interesting test, something more advanced and impressive to someone as a demo (maybe gif if it's a simple physics engine or something)... make the project manager decision and come at me with a ready to build plan for the mini project benchmark."*
- **Note:** This benchmark was intentionally run with a loose, open-ended prompt to stress-test whether PIOS phase gates would still enforce a structured output even without a structured input. They did.
- **Outcome:** A polished HTML5 Canvas & Matter.js implementation featuring a dark theme, neon glass blocks, and mouse interaction. Zero compilation required.

#### 3. Embeddable Excuse Generator Widget
- **Prompt:** *"For the third and final pre-release test, random excuse generator for a button that loads a different excuse each time, designed as a widget that can be embedded and uploaded to common online stores for secure efficient CWV friendly embed on websites."*
- **Outcome:** A pristine Vanilla JS IIFE injecting scoped CSS and HTML, ensuring zero layout shifts (CWV-friendly) and a lightweight footprint.

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

- [x] v0.1 — templates, agents, adapters, workflows, backtest harness
- [x] v0.2 — machine-readable state
- [x] v0.3 — Golang CLI: `pios init / validate / status`
- [x] v0.4 — Contract Hardening & Positioning Reset
- [ ] v1.0 — Stable Release Matrix: 
  - Implementation of **Model Context Protocol (MCP)** server embedding. This will allow Claude Code and Cursor users to natively integrate PIOS phase validations as direct tools (e.g. `call pios_validate()`) without needing to invoke shell commands.
  - Native distribution via Homebrew (`brew install pios`) and Winget.
  - Comprehensive VitePress/Nextra documentation site.
