# GPT Adapter (PIOS)

Best for:
- balanced architecture + implementation
- structured artifacts
- iterative refinement

Recommended pattern:
- Start with Spec Lock + Plan Lock generation.
- Then generate files in batches by directory.
- Keep a running STATUS.md and decision log.

Avoid:
- generating entire repo in one huge message if tool cannot write files directly
