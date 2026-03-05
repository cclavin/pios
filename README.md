# π PIOS — Project Intelligence Operating System

![Version](https://img.shields.io/badge/version-0.3.0-orange) ![License](https://img.shields.io/badge/license-MIT-orange)

PIOS is a structured, tool-agnostic framework for **starting and finishing** AI-assisted software projects.

It’s designed for builders who use AI tools (Claude, GPT, Codex, Cursor, Continue, Windsurf, OpenClaw, etc.) and want to reduce:
- endless back-and-forth prompting
- context drift / "context rot"
- vague planning that never turns into a repo
- stalled projects with no clear "done" state

---

## Core Positioning

PIOS is **artifact-first**: it produces repo files, phase gates, and repeatable workflows — not just chat.

- **Completion-first:** PIOS operates on **phase gates** with explicit exit criteria (Minimum Spec → Spec Lock → Plan Lock → Task Lock → Scaffold Done) and an **autopilot loop** (implement → test → fix → document → summarize).
- **Interop-first:** PIOS is not a competing "rules format." It is a source-of-truth export target. Maintain core guidance in PIOS and export to `AGENTS.md` (tool-agnostic baseline) and tool-specific adapters. This keeps your repo consistent even as tools change.
- **Artifact-first:** Focuses strictly on files, diffs, commands, and checklists. Minimal prose.

---

## Core Philosophy

PIOS exists to convert AI assistance into **repeatable outcomes**.

1. **Contracts over vibes:** Start with structured inputs (specs) that can be refined without restarting.
2. **Artifacts over chat:** Produce files, diffs, commands, and checklists.
3. **Phase gates:** Each phase ends only when exit criteria is met.
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

---

## Roadmap

- [x] v0.1 — templates, agents, adapters, workflows, backtest harness
- [x] v0.2 — machine-readable state
- [x] v0.3 — Golang CLI: `pios init / validate / export`
- [ ] v0.4 — interop exporters for Cursor / Continue / Windsurf formats
- [ ] v1.0 — stable “project OS” release
