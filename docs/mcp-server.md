# PIOS Model Context Protocol (MCP) Server

> **Status: Planned for v1.0**

As part of the PIOS v1.0 roadmap, we are developing an [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) server wrapper for the PIOS Go CLI. 

By default, an AI agent interacting with PIOS must execute shell commands (e.g., `pios validate`) and parse the standard output to determine if a phase gate is passed. While effective, executing raw bash commands can be restricted in secure environments or cause parsing errors in certain AI tools.

## The MCP Solution

The `pios-mcp` server connects local LLM clients (like **Claude Code**, **Cursor**, or **Windsurf**) directly to the PIOS validation engine via a standardized JSON-RPC interface.

When PIOS is installed as an MCP server, the AI agent is given native "Tools" it can call deterministically, completely bypassing the shell.

### Proposed Tool Schema

1. `pios_status()`
   - **Description:** Reads the repository's `STATUS.md` and active `TASKS.md` to determine the current phase, gate, and next actions.
   - **Returns:** A structured JSON object containing the `current_phase`, `current_gate`, and a list of remaining tasks.

2. `pios_validate()`
   - **Description:** Runs the core Go-based validation logic against the current file tree.
   - **Returns:** A boolean `success` flag, along with detailed `errors` (e.g., "Found 2 unchecked tasks in Milestone 1").

3. `pios_init()`
   - **Description:** Ejects the PIOS templates into the current working directory.
   - **Returns:** Confirmation that the `STATUS.md` and `templates/` folder have been seeded.

## How it will look for Claude Code Users

Once v1.0 is released, adding PIOS to a Claude Code environment will be as simple as:

```bash
# Add PIOS as a local MCP server
claude mcp add pios --command "pios mcp-run"
```

Once added, Claude will automatically understand how to check project boundaries before writing code, drastically reducing context drift and hallucination.
