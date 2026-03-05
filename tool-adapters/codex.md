# Codex Adapter (PIOS)

Best for:
- direct code changes
- incremental commits
- tight implement-test loops

Recommended pattern:
- Feed Codex the Plan Lock + Tasks.
- Execute one task at a time.
- Require tests or smoke checks per task.
- Use autopilot loop: implement → test → fix → doc → summarize

Avoid:
- expecting deep architecture without Plan Lock
