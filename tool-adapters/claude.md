# Claude Adapter (PIOS)

Best for:
- long-form reasoning
- structured Markdown artifacts
- large blueprint creation

Recommended pattern:
- Produce Plan Lock + Tasks first.
- Generate artifacts as clean Markdown “file blocks” (path + content).
- Use strict phase gates to reduce drift.

Avoid:
- micro-prompts that fragment context

## Core PIOS Instructions
> **Note:** PIOS is an execution contract layer, not a runtime platform. Follow these strict validations explicitly.

3. **Phase Validation:** You must run `pios validate` and ensure it passes before concluding a phase gate.

> **Coming soon in v1.0:** PIOS is developing a native MCP (Model Context Protocol) server. Soon, tools like Claude Code will be able to invoke `pios_validate()` and `pios_status()` natively via JSON-RPC, rather than relying on terminal execution. For now, strictly use the shell commands.
