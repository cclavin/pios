# PIOS Next Phase Strategy

**Title:** PIOS: The AI Project Execution Contract
**Document Type:** Canonical Strategy Memo
**Target Version:** v0.4.0 (Positioning & Contract Hardening)

---

## Executive Summary
PIOS is entering its next maturity phase. Following an external review and our v0.3.1 hardening release, it is clear that PIOS's core value is its role as an **artifact-first project execution contract**, not an AI agent runtime platform. This document outlines the strategic pivot necessary to clarify this positioning, protect the core value of deterministic phase gates, and cleanly decouple PIOS from the runtime orchestration platforms (like PAI) that it complements. The next 30-60 days focus entirely on positioning clarity and contract hardening.

## Current Diagnosis
The initial "Operating System" framing carried overlap risk with full-stack agent infrastructure. PIOS has a distinct, proven mechanism—file-based phase gating and CLI validation—that prevents context drift and ensures project completion. However, ambiguous messaging permitted readers to misunderstand PIOS as an alternative to execution environments. In reality, PIOS is a complementary workflow contract. Our recent v0.3.1 release proved that our CLI and templates can be rigorously tested and safely executed across tools. Now, our documentation and strategic framing must catch up to that reality.

## Approved Direction
**PIOS will be repositioned explicitly as an AI Project Execution Contract.**
We will retain the PIOS acronym but retire the primary messaging of "Project Intelligence Operating System" in favor of framing it as an "Operating Standard" or an "Execution Contract." The framework exists to optimize four outcomes: finishability, portability, measurability, and trust.

## Non-Negotiable Principles
The locked principles guiding PIOS remain accurate and are reaffirmed here:
1. **PIOS remains tool-agnostic:** It must work equally well with Cursor, Claude Code, Windsurf, or any future agent. It owns the *contract*, not the *client*.
2. **PIOS is an execution contract layer:** It defines what "done" means via structured markdown artifacts.
3. **CLI gate checks are non-negotiable:** The `pios validate` tool is the engine of state transitions. Agents cannot move forward without satisfying deterministic testing requirements.
4. **Runtime orchestration is out of scope:** We do not host models, manage memory backends, run sandboxed code, or orchestrate API calls.
5. **Exporter work is deferred unless proven necessary:** Focus on stabilizing the core contract before building `.cursorrules` or `.windsurfrules` compilers. The current lightweight adapter model works.
6. **Next milestone is positioning + contract hardening:** No feature sprawl until the core messaging is perfectly documented and resilient.

## Layering Model
Understanding where PIOS sits in the AI development stack is critical:

| Layer | Component | PIOS Relationship |
|---|---|---|
| **Agent Runtime Layer** | PAI, AutoGPT, Claude Code | External. Executes the work loop. |
| **Coding Assistant Layer** | Cursor, Copilot, Windsurf | External. IDE integration point. PIOS adapters guide their behavior. |
| **Project Execution Contract** | **PIOS Core** | **Source of Truth.** The gatekeeper for phase transitions and artifacts. |
| **Memory / Context** | Vector DBs, Conversation History | External. PIOS outputs become portable context, but PIOS does not manage memory storage. |
| **Project Template Layer** | Repo Scaffolding | Provided by PIOS via `pios init`. Minimal, opinionated starting points. |
| **Repo Convention Layer** | Code formatting, linters | Enforced by the PIOS phase gates, executed by the runtime. |

## PIOS vs Runtime Infrastructure
PIOS is highly complementary to personal AI infrastructure and runtime systems like PAI. 
* **PAI** handles capability, persistence, tool execution, memory, and orchestration. It is the *engine*.
* **PIOS** handles the definition of done, phase gating, and artifact validation. It is the *track*.

A user relies on their runtime infrastructure to *do the work*, but layers PIOS into the repository to *prove the work is finished*. They do not compete; an agent orchestrated by PAI can be instructed to obey a PIOS contract in a given repository, ensuring the agent actually ships instead of looping endlessly.

## In-Scope for Next Phase
* **Positioning Reset:** Fully rewrite the README, adding explicit boundary statements and comparison documents.
* **Contract Hardening:** Formalize the expected syntax for `STATUS.md` and `templates/tasks.md` into a documented standard. Complete the test suite coverage for edge cases in `cmd/pios/main.go`.
* **Workflow Refinement:** Ensure all workflows (like the autopilot loop) are entirely tool and shell agnostic.

## Deferred Items
* **Interop Exporters:** We will postpone transacting `.rules` files programmatically (`pios export`) until the community demands it, as the markdown artifacts currently suffice.
* **Automation Hooks:** While triggering CI steps upon phase completion is valuable, it is deferred to avoid diluting the core contract work.

## Explicitly Out of Scope
* Model runtime orchestration or API wrapping.
* Memory backends, RAG implementations, or vector stores.
* Agent execution sandboxes.
* Hosted SaaS control planes or telemetry dashboards.

## Repo & Documentation Implications
The repository must reflect the new positioning immediately. 
We must scaffold a `docs/` folder structured around these boundaries:
* `docs/positioning.md`
* `docs/scope.md`
* `docs/layering.md`
* `docs/contracts.md` (defining the artifact grammar)
This structure replaces vague roadmap assertions with concrete architecture boundaries.

## Technical Contract-Hardening Implications
From a technical standpoint, the `pios` CLI must evolve from a "script replacement" into a rigid parser. Artifact versions should be checked (e.g., rejecting an outdated `tasks.md` format), validation errors must be highly specific, and the CLI must remain statically compiled, cross-platform Go with no obscure OS-level dependencies. The recent regex improvements in v0.3.1 established the baseline; the next step is formal schema documentation for that logic.

## Risks to Avoid
1. **The "Platform" Trap:** Bleeding into runtime orchestration out of excitement, abandoning our unique moat of finishability.
2. **Over-Correction:** Stripping PIOS down so far that it becomes a generic README template rather than an enforced workflow. The CLI must remain the anchor.
3. **Adapter Sprawl:** Adding adapters for every obscure tool before the core phase loop is bulletproof.

## Next 10 Actions
1. Review and finalize this strategy document (`docs/strategy/pios-next-phase.md`).
2. Rewrite `README.md` intro, quickstart, and scope statement to enforce the Execution Contract positioning.
3. Create `docs/scope.md` detailing explicit in/out boundaries.
4. Create `docs/layering.md` formalizing the stack placement.
5. Create `docs/contracts.md` explaining the regex expectations for tasks and status checks.
6. Scrub all existing documentation for legacy "Operating System" references that imply infrastructure capabilities.
7. Verify all tool adapters (`claude.md`, `cursor.md`, etc.) are actively aligned with the new contract-first messaging.
8. Establish a `main_test.go` pattern for all future task parsing changes (done in v0.3.1).
9. Finalize the `v0.4.0` roadmap to exclusively target Doc/Positioning completion.
10. Update the github repository description to read: "Artifact-first execution contract for AI-assisted software delivery."

## Final Recommendation
I fully endorse the strategic pivot to position PIOS as an AI Project Execution Contract. The locked principles provided are sound and necessary. We should proceed with the proposed documentation overhaul immediately to secure this positioning before pursuing further technical features.
