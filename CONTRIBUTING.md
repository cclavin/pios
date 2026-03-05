# Contributing to PIOS

Welcome! PIOS is an open-source framework designed to ensure AI-assisted software projects actually cross the finish line. We openly welcome community contributions to expand our tool integrations, stack profiles, and core tooling.

## The Golden Rule (Artifact-First)
Before submitting a Pull Request, ensure your changes align with the core philosophy of PIOS: **completion-first, artifact-first**. 
We do not accept PRs that introduce philosophical bloat, lengthy prose, or highly opinionated "prompt engineering" masquerading as system rules.

## Contributing to Profiles & Adapters

### Adding a Stack Profile (`profiles/stacks/`)
When adding a new stack (e.g., `rust-actix.md` or `python-django.md`), ensure that the profile only contains *universal architectural standards* for that language. Do not enforce hyper-specific linters or formatting tools unless they are the definitive community standard (like `gofmt` for Go).

### Adding an AI Tool Adapter (`tool-adapters/`)
An AI tool adapter (e.g., `windsurf.md`, `cursor.md`) translates the PIOS baseline for a specific AI IDE or Agent.
- **Required**: The adapter *must* instruct the AI to utilize `pios status` to check the current project phase.
- **Required**: The adapter *must* instruct the AI to only check out tasks marked `[ ]` in `templates/tasks.md`.
- **Required**: The adapter *must* instruct the AI to run `pios validate` before concluding a phase gate.

## Contributing to the Core CLI (`cmd/pios/`)
The native validation CLI is written in Golang. 

### Local Development
1. Ensure you have Go `1.22+` installed.
2. Clone the repository and navigate to the root directory.
3. If changing templates, note that they are natively bundled via `//go:embed` in `templates/embed.go`.
4. Compile and test the binary locally:
   ```bash
   go build -o pios.exe ./cmd/pios
   ./pios.exe init
   ./pios.exe status
   ```

### Code Standards
- Adhere strictly to the Go Standard Project Layout.
- Ensure cross-platform compatibility (macOS, Linux, Windows).
- Do not introduce heavy third-party dependencies unless strictly required for parsing algorithms (e.g., `yaml.v3`).

## Submitting a Pull Request
1. Fork the repository and create your feature branch (`feat/your-feature`).
2. Adhere to our `git-standards.md`—make atomic commits using imperative mood.
3. **Mandatory**: If your change adds a feature, you MUST update `CHANGELOG.md` under the `[Unreleased]` block in the very same commit.
4. Ensure your PR description clearly states the *Why* behind the change.
