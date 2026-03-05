# Coding Standards (Base)

- Prefer readability and simplicity.
- DRY where it reduces duplication; avoid premature abstraction.
- Small modules, clear names, explicit types where helpful.
- Defensive error handling with useful messages.
- Add tests for critical behavior.
- Prefer modern language idioms and current best practices.
- **Secrets Rule**: NEVER store, hardcode, or commit secrets, tokens, or API keys in plain text in this repository. Secrets must always be dynamically requested from the OS credential store (macOS Keychain, Linux libsecret, pass, or Windows Credential Manager).
