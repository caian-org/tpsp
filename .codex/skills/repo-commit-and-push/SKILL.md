---
name: repo-commit-and-push
description: Use when committing or pushing changes in this repository. Create focused commits matching tpsp history, push only when requested, and push tags whenever pushing release work.
---

# Repo Commit And Push

Use this skill when the task includes committing or pushing changes.

## Policy

- Push only when the user explicitly asks.
- Keep commits focused on one coherent purpose.
- Do not stage unrelated user changes.
- Do not edit or commit generated output such as `bin/tpsp`.
- If pushing a release branch or tag-related work, push tags in the same work
  unit when appropriate.

## Commit Format

Recent history mixes short imperative subjects with conventional prefixes such
as `fix:` and `build:`. Prefer conventional commits when the type is clear:

```text
fix: handle API decode errors
build: update GoReleaser config
docs: document CLI examples
```

Use an imperative summary and add a short body when the change is not obvious.

## Workflow

1. Inspect `git status --short` and recent history.
2. Group files by purpose.
3. Stage explicit paths for one group at a time.
4. Run the validation relevant to that group before committing.
5. Commit with a focused subject.
6. Push only if requested by the user.

Never bypass hooks unless the user explicitly requests it.
