# Git Standards (Base)

## Version Control
- Write commit messages using the imperative mood (e.g. "Add feature" not "Added feature").
- Never use `--no-verify` or bypass git hooks without explicit permission.
- Make commits atomic and focused on a single logical change.
- Never force push (`--force`) without checking with the user first.
- **Documentation Sync**: When committing a notable feature or fix, the commit MUST include the corresponding update to `CHANGELOG.md` in the same logical commit.
