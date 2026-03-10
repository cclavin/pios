# PIOS Scope (v1.0.0)

This document locks what is in scope for v1.0.0.

## In Scope

- Positioning reset to AI Project Execution Contract language
- Boundary documentation separating PIOS from runtime platforms
- Contract specification lock for `STATUS.md` and `templates/tasks.md`
- Validator and test hardening for deterministic gate checks
- Cross-platform workflow documentation without shell-specific assumptions

## Deferred

- Interop exporters (for example `.cursorrules` or `.windsurfrules` generation)
- Automation hooks triggered by phase transitions
- Broad adapter expansion beyond current maintained set

## Out of Scope

- Model runtime orchestration or API wrapping
- Memory backends, RAG, or vector store ownership
- Agent sandbox execution
- Hosted SaaS control plane or telemetry dashboards
- New CLI commands beyond `init`, `status`, and `validate`
