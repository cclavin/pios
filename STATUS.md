---
pios_version: "0.2"
current_phase: "v0.1 scaffolding"
current_gate: "Plan Lock"
status: "In Progress"
---
# STATUS

## Current Phase
- Phase: v0.1 scaffolding

## Next 3 Actions
1) Create repo files (templates/agents/adapters/workflows)
2) Add one example project
3) Run one backtest run (baseline vs PIOS) and document results

## Open Questions
- What should PIOS export to by default beyond `AGENTS.md`?
- Should we ship a minimal CLI in v0.1 or wait until v0.3?
- Should stack templates be opinionated (“strict mode”) or optional?

## Risks
- Becoming “yet another prompt framework” (mitigate via interop + completion gating)
- Over-scoping too early (mitigate via v0.1 minimalism)

## Definition of Done (v0.1)
- Templates exist and are usable
- Agents exist and are coherent
- Tool adapters exist for 3+ tools
- Backtest harness exists
- One example exists
