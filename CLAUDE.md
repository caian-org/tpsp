# Claude Notes

Pointer doc for Claude Code agents working in `tpsp`.

**Read `AGENTS.md` first**. It is the canonical instruction layer for project
structure, commands, tests, release packaging, and agent-specific rules. This
file only covers Claude Code orientation.

## Bootstrap

1. `AGENTS.md` - canonical repository rules.
2. `README.md` and `docs/README.md` - user-facing CLI behavior.
3. Relevant `.codex/skills/*/SKILL.md` files for the subsystem being changed.

If this file conflicts with `AGENTS.md`, follow `AGENTS.md` and reconcile this
file in the same change set.

## Claude Code specifics

- Use Claude Code's native subagent feature when the session policy allows it.
  Specialists live under `.claude/agents/` and mirror `.codex/agents/`.
- Skills are not duplicated under `.claude/skills/`; the canonical home is
  `.codex/skills/`.
- Claude Code settings live in `.claude/settings.json`. Hooks are currently
  empty; keep future hooks fail-open unless the repository explicitly requires
  hard blocking behavior.
- Use the local checkout for repository work. Do not use GitHub API reads as a
  substitute for inspecting files here unless it is a narrow metadata check.
- Serialize shared-checkout mutations: edits, formatting, patch application,
  staging, committing, branch switching, rebasing, and pushing.
- Do not edit generated output by hand, especially `bin/tpsp`.

## Quick Reference

- `just build` - build `bin/tpsp`.
- `just clean` - remove generated `bin/tpsp`.
- `go test ./...` - run package tests and vet.
- `go vet ./...` - run vet explicitly.
- `go build -o /tmp/tpsp-check ./cmd/tpsp` - disposable build check.
- `go run ./cmd/tpsp --help` - inspect CLI help.
- `go run ./cmd/tpsp --version` - inspect version output.
- `go run ./cmd/tpsp --copyright` - inspect copyright output.

High-traffic skills:

- `.codex/skills/tpsp-dev-workflow/SKILL.md`
- `.codex/skills/tpsp-cli-http/SKILL.md`
- `.codex/skills/repo-commit-and-push/SKILL.md`
