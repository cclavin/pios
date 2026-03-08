# Archived Pre-Release Benchmarks (v0.4.0)

This directory contains the artifacts and metrics from the first three automated benchmark runs of the PIOS framework. 

All three runs were conducted using HTML5 Canvas / Vanilla JS without a heavy compilation step. They serve as the baseline proof-of-concept for the PIOS execution contract.

## Methodological Findings

### Where PIOS Succeeded
1. **Speed & Consistency:** Every project went from zero to a visually impressive, functioning application in under 15 minutes.
2. **Deterministic Outputs:** By forcing the AI Agent through the `pios validate` phase gates, we completely eliminated the "code vomit" problem. The agent accurately drafted the `min-spec.md`, locked the plan, generated isolated tasks in `tasks.md`, and burned them down sequence by sequence.
3. **No Conversational Drift:** The agent never lost context of the goal because the contract state (`STATUS.md`) anchored it during every step of the autopilot loop.
4. **Resilience to "Vibes":** Even when provided a highly unstructured, loose prompt (e.g., the *Physics Destruction Sim*), the constraints of the framework forced the AI to normalize the unstructured input into a highly structured execution plan before writing code.

### Where the Flow Needs Improving for V1.0
1. **Dependency Management:** The HTML5 tests succeeded because they had zero dependencies (other than Matter.js via CDN). V1.0 needs to prove PIOS can handle stateful package management (e.g., `go mod tidy`, `npm install`). 
2. **Multi-file Architectures:** We need to validate that an agent can correctly separate concerns across multiple files (routes, controllers, models) rather than stuffing everything into an `index.html` or `widget.js`.
3. **Shell Friction:** Currently, the agent relies heavily on bash environment execution (`pios validate`). This can occasionally cause shell escaping errors depending on the IDE. This highlights why the **native MCP Server integration** planned for v1.0 is the most critical feature on the roadmap.

---

*These benchmarks prove the underlying philosophy works. The next step is scaling complexity.*
