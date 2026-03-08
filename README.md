# π PIOS — AI Project Execution Contract

![Version](https://img.shields.io/badge/version-0.4.0-orange) ![License](https://img.shields.io/badge/license-MIT-orange)

PIOS is a structured, tool-agnostic [execution contract](docs/positioning.md) for **starting and finishing** AI-assisted software projects. It is [strictly scoped](docs/scope.md) as a contract layer, **not a runtime orchestration platform**. It's designed to eliminate:
- endless back-and-forth prompting
- context drift / "context rot"
- vague planning that never turns into a repo
- stalled projects with no clear "done" state

---

## Core Positioning

PIOS is **artifact-first**: it produces repo files, phase gates, and repeatable workflows — not just chat.

- **Completion-first:** PIOS operates on **phase gates** with explicit exit criteria (Minimum Spec → Spec Lock → Plan Lock → Task Lock → Scaffold Done) and an **autopilot loop** (implement → test → fix → document → summarize).
- **Contract-first:** PIOS is not a competing "rules format." It establishes a source-of-truth contract. Maintain core guidance in PIOS via `AGENTS.md` (tool-agnostic baseline) and tool-specific adapters. This keeps your repo consistent even as tools change.
- **Artifact-first:** Focuses strictly on files, diffs, commands, and checklists. Minimal prose.

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
PIOS comes with a native CLI to instantly copy the templates into your new repository and programmatically track your AI agent's progress.

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
To validate the PIOS execution contract model, the following three mini-projects were executed autonomously using the `pios` CLI strict phase gates. Each project went from a blank directory to validated completion in **under 20 minutes**.

1. **Dual-Game Menu & Mechanics Showcase**
   - **Prompt:** *"a game menu with two games built in and able to choose from where one the cat is just sleeping and it slightly reacts with user click or keyboard interaction. Another game in the menu that's brighter and more active, 2d platformer with slick ice physics and tuck down on keyboard (up down left right w a s d normal inputs) to build up speed on drop in altitude or downhill with test map of big hill and jump please, this must all be included in first run so do anything that's needed to test and complete"*
   - **Outcome:** A cohesive HTML5 Canvas experience housing a menu that links to a serene interactive cat toy and a high-speed momentum-based 2D platformer. Demonstrates the framework's ability to scaffold non-trivial logic matrices accurately.
2. **Physics Destruction Simulator**
   - **Prompt:** *"Come up with a plan for a more interesting test, something more advanced and impressive to someone as a demo (maybe gif if it's a simple physics engine or something)... make the project manager decision and come at me with a ready to build plan for the mini project benchmark."*
   - **Outcome:** A polished HTML5 Canvas & Matter.js implementation featuring a dark theme, neon glass blocks, and mouse interaction. Zero compilation required.
3. **Embeddable Excuse Generator Widget**
   - **Prompt:** *"For the third and final pre-release test, random excuse generator for a button that loads a different excuse each time, designed as a widget that can be embedded and uploaded to common online stores for secure efficient CWV friendly embed on websites."*
   - **Outcome:** A pristine Vanilla JS IIFE injecting scoped CSS and HTML, ensuring zero layout shifts (CWV-friendly) and a lightweight footprint.

---

## Roadmap

- [x] v0.1 — templates, agents, adapters, workflows, backtest harness
- [x] v0.2 — machine-readable state
- [x] v0.3 — Golang CLI: `pios init / validate / status`
- [x] v0.4 — Contract Hardening & Positioning Reset
- [ ] v1.0 — stable “execution contract” release
