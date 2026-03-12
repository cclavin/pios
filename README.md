# π PIOS - AI Project Execution Contract

[![Version](https://img.shields.io/github/v/release/cclavin/PIOS?color=orange&label=version)](https://github.com/cclavin/PIOS/releases) [![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go&logoColor=white)](https://go.dev/) [![CI](https://github.com/cclavin/pios/actions/workflows/ci.yml/badge.svg)](https://github.com/cclavin/pios/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/cclavin/pios)](https://goreportcard.com/report/github.com/cclavin/PIOS) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

<p align="center">
  <img src="pios-banner.gif" alt="PIOS - Contracts over vibes. Build useful things, faster." />
</p>

**Works with:** Claude Code · GPT · Cursor · Windsurf · Codex · Continue · OpenClaw · OpenHands · Cline · any tool that reads Markdown

## Table of Contents
- [Mission](#mission)
- [How It Works](#how-it-works)
- [Repository Layout](#repository-layout)
- [Tutorials & Workflows](#tutorials--workflows)
- [Example Prompts](#example-prompts)
- [Backtesting PIOS](#backtesting-pios)
- [Continuous Building Loop](#the-continuous-building-loop-pios-next)
- [Related Ecosystems](#related-ecosystems--methodologies)
- [Roadmap](#roadmap)

## Mission

I've struggled with effectively completing projects, and I'm sure we all have a graveyard of folders and old projects that had something there, but the momentum did not continue. PIOS lets you carry momentum.

I built it because too many AI workflows turn into endless chat instead of forward motion. Spend less time back-and-forth in chat, and more time deploying, dreaming new features, and rapidly prototyping.

PIOS gives that momentum a structure. It uses deterministic artifacts, phase gates, and explicit validation so the repo stays grounded in what is actually done, not what was vaguely discussed.

I wanted to make effective AI-assisted development more approachable to beginners, but highly powerful for advanced users, too.

---

## How It Works

PIOS is **artifact-first**: it produces repo files, phase gates, and repeatable workflows, not just chat.

- **Completion-first:** PIOS operates on **phase gates** with explicit exit criteria. The CLI strictly enforces both artifact existence (specs must exist before coding) and checklist completion (Minimum Spec → Spec Lock → Plan Lock → Task Lock → Scaffold Done).
- **Contract-first:** PIOS is not a competing "rules format." It establishes a source-of-truth contract. Maintain core guidance in PIOS via `AGENTS.md` (tool-agnostic baseline) and tool-specific adapters. This keeps your repo consistent even as tools change.
- **Artifact-first:** Focuses strictly on files, diffs, commands, and checklists. Minimal prose.
- **Framework & Runtime Agnostic:** PIOS doesn't care if you use Cursor, Copilot, Windsurf, or a background Claude/OpenAI agent. It serves as the immutable data layer that any AI can read to understand exactly what to do next.

---

## Repository Layout

- `cmd/pios/` — the Golang PIOS CLI + MCP server
- `templates/` — reusable project artifacts (specs, plans, tasks, decision logs) safely embedded into the CLI
- `agents/` — role-specific agent instructions
- `tool-adapters/` — guidance for using PIOS with specific tools
- `profiles/` — shared standards + stack templates
- `workflows/` — consolidated phase-by-phase operating flows and commands

---

## Tutorials & Workflows

PIOS works best when you make two decisions up front: how your tool connects to the contract, and who advances the milestone. The recommended starting point is MCP or CLI with a human reviewing each completed milestone. Autonomous looping is available, but is still experimental.

### 1. The MCP-Native Path (Recommended)

This is the best fit for Claude Code, Cursor, and Windsurf setups that support MCP. Instead of parsing terminal output, the agent can call PIOS tools directly over JSON-RPC.

**The Setup:**
*   **For Claude Code:** Run `claude mcp add pios-mcp -- pios mcp`
*   **For Cursor:** Open Settings > Features > MCP and add a shell command server with the command `pios mcp`.
*   **For Windsurf:** If your setup supports MCP servers, attach PIOS with the command `pios mcp`.

**The Daily Workflow:**
1.  Run `pios init` in a new folder, or let the agent call `pios_init` in an empty project.
2.  Fill out `templates/min-spec.md`, then lock the scope and planning artifacts.
3.  Let the agent call `pios_status` before coding, `pios_validate` before closing a milestone, and `pios_next` only when the current milestone is complete and you want to continue.

### 2. The Command-Line Path (Default Human-Gated Loop)

If your tool can edit files and run shell commands but does not have MCP attached, use the CLI directly. This is the safest default for Codex, Continue, local open-source agents, and terminal-first workflows.

**Install PIOS:**
Native package-manager installs via Homebrew and Winget are intended to be the primary path. Until those packages are published, use the Go fallback:
```bash
go install github.com/cclavin/pios/cmd/pios@latest
```

**Initialize a new project:**
```bash
pios init --ide=cursor
```

This seeds the contract into the repo and can also scaffold IDE-specific rule files. Supported `--ide` values are `cursor`, `windsurf`, and `claude`. From there, the standard loop is simple: let the agent work one task at a time, use `pios status` to check context, use `pios validate` to close the gate, and run `pios next` only after you decide the milestone is done.

### 3. The Existing Repo / Retrofit Path

PIOS is not only for greenfield projects. You can adopt it in an active repo without restructuring your app.

1.  Run `pios init` at the repo root.
2.  Backfill `templates/min-spec.md`, `templates/spec-lock.md`, `templates/plan-lock.md`, and `templates/tasks.md` from the current state of the project.
3.  Update `STATUS.md` so it reflects the milestone and gate you are actually in.
4.  Resume work under normal PIOS validation.

If your repo already uses a root `templates/` directory for app assets or generator output, decide how you want to resolve that conflict before adopting PIOS so the contract layout stays stable.

### 4. The Vendored Contract Path (No Go Install Yet)

If you do not want to install Go yet, you can still use PIOS as a Markdown contract. Copy the core files into your repo and keep the canonical layout so you can adopt the CLI later without moving anything around.

- Keep `AGENTS.md` at the repo root
- Keep `STATUS.md` at the repo root
- Keep `templates/` at the repo root

### 5. The Chat-Only Path (Manual Phase Gates)

If you are working in a chat tool that cannot write files or run commands, PIOS is still usable.

1.  Keep the canonical PIOS files in the repo root.
2.  Fill out the specs and task list yourself.
3.  Give the chat tool the active artifacts it needs, especially `AGENTS.md`, `STATUS.md`, and `templates/tasks.md`.
4.  You become the validator: update checkboxes, review output, and decide when the next gate or milestone is allowed.

### Optional Bootstrap: The Zero-to-Hero Prompt

If you have a shell-capable agent in an empty directory and want a one-shot bootstrap, you can give it a prompt that installs PIOS using the best available method for the host OS, runs `pios init`, reads `AGENTS.md`, and starts building. Treat this as a bootstrap recipe, not the default daily workflow.

**Example Prompt:**
> "First, install PIOS using the best available method for my OS. Prefer a native package-manager install if PIOS is available through Homebrew or Winget. If not, install Go 1.22+ and then run `go install github.com/cclavin/pios/cmd/pios@latest`.
>
> Next, create a new directory for this project, enter it, and run `pios init`. After initialization, read the `AGENTS.md` file to understand the contract. Finally, proceed through the PIOS phases to build me a Python script that scrapes Hacker News."

**Where Different Tools Fit:**
- Claude Code, Cursor, and Windsurf are best on the MCP-native path when MCP is available.
- Codex, Continue, and shell-capable open-source agents are best on the CLI path.
- ChatGPT and similar chat-only tools fit the manual path.
- OpenClaw-style autonomous agents fit the CLI path today, then layer on the experimental loop below.

For tool-specific setup details, see [Claude](tool-adapters/claude.md), [Cursor](tool-adapters/cursor.md), [Windsurf](tool-adapters/windsurf.md), [Codex](tool-adapters/codex.md), [Continue](tool-adapters/continue.md), and [OpenClaw](tool-adapters/openclaw.md).

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

Once a milestone is complete and all active tasks are marked `[x]`, `pios next` prepares the repo for the next milestone. It is a transition command, not a planning engine.

There are two solid ways to use it:

**1. Recommended: Human-Gated Continuation**
Review the completed milestone first. When you are satisfied, run `pios next`. Then write the next spec, plan, and task artifacts before asking the agent to continue.

**2. Experimental: Agent-Driven Continuation**
If your agent is connected through MCP or has reliable shell access, it can run `pios_next` or `pios next` itself after the milestone is complete. This works best on bounded projects where the agent is already following the task contract cleanly.

When you are ready for the next feature, run:
```bash
pios next
```

This executes the PIOS transition loop:
1. **Snapshot:** It creates a timestamped archive (`templates/archive/YYYY-MM.../`) of your completed `TASKS.md` and `STATUS.md` so you never lose the history of your technical decisions.
2. **Sweep:** It scrubs all `[x]` checked tasks from your active tasks file, leaving behind a clean board containing any rolled-over pending items.
3. **Reset:** It resets your `STATUS.md` phase gates back to planning mode.

What it does **not** do by itself is invent the next milestone, rewrite the spec, or choose the next product direction. After the reset, either you or the agent still needs to draft the next spec, plan, and task list before implementation resumes.

If you want to experiment with autonomous continuation, use these guardrails:
1. Only call `pios_next` after `pios_validate` passes and the current milestone is genuinely complete.
2. Keep milestones small and concrete.
3. Review the archive and the next task plan at each loop.
4. Prefer human gating for production, security-sensitive, or multi-developer repos.

**Experimental Full-Autonomy Prompt:**
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
9. Stop and summarize if blocked, if a product-direction decision is needed, or if deployment or security review needs a human.
```

If you are connected via the **MCP Server**, your AI agent can run `pios_next` itself. A good short prompt is: *"Great job on the auth layer. Run the next command, then draft Milestone 2 spec, plan, and tasks focused on the database adapter. Do not start implementation until those artifacts are updated."*

---

## Related Ecosystems & Methodologies

PIOS is designed to be modular. It doesn't replace your favorite methodologies; it gives them a machine-readable execution layer. PIOS pairs exceptionally well with:

- **[Daniel Miessler's Personal AI Infrastructure](https://github.com/danielmiessler/Personal_AI_Infrastructure):** For users looking to standardize their global AI behavior and customized context. Combine a robust personal AI infrastructure (to define your developer identity) with PIOS (to enforce project-specific finishing) for a highly deterministic, end-to-end AI capability.
- **Get Shit Done (GSD):** If you already subscribe to fast-iteration, momentum-focused workflows, PIOS serves as the strict, artifact-driven enforcer of those philosophies. It stops AI chat loops so you actually get the thing built.
- **System Prompt Libraries:** Whether you use custom `CLAUDE.md` files or specialized Cursor profiles, PIOS operates flawlessly underneath them as the source-of-truth project state.

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
