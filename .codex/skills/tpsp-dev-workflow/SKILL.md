---
name: tpsp-dev-workflow
description: Local development workflow for the tpsp Go CLI. Use when orienting in the workspace, choosing build/test commands, validating CLI behavior, or changing release/package files.
---

# TPSP Dev Workflow

Use this skill for repository orientation, command selection, and everyday
implementation hygiene across `tpsp`.

## Repository Shape

- `cmd/tpsp/main.go` - single Go CLI entrypoint and most runtime behavior.
- `README.md` and `docs/README.md` - user-facing behavior and examples.
- `Makefile` - local build command that writes `bin/tpsp`.
- `.goreleaser.yaml` - release artifacts for linux, darwin, and windows.
- `Dockerfile` - static binary runtime image using `scratch`.
- `.github/workflows/release.yml` - tag-driven GoReleaser workflow.
- `.github/dependabot.yml` - dependency update configuration.

Generated output belongs in `bin/` or temporary paths outside the repo. Do not
edit `bin/tpsp` by hand.

## Commands

```sh
make
go build -o bin/tpsp ./cmd/tpsp
go build -o /tmp/tpsp-check ./cmd/tpsp
go test ./...
go vet ./...
gofmt -w cmd/tpsp/main.go
go run ./cmd/tpsp --help
go run ./cmd/tpsp --version
go run ./cmd/tpsp --copyright
```

Use `/tmp/tpsp-check` for disposable validation builds to avoid changing tracked
or ignored build output.

## Implementation Rules

- Preserve the existing public CLI contract unless the user explicitly asks for
  a breaking change.
- Keep the JSON output stable: `code`, `data`, and `message`.
- Keep service names stable: `metro`, `cptm`, `viamobilidade`, `viaquatro`.
- Keep the project disclaimer accurate; never imply official affiliation.
- Prefer small, table-driven tests for parsing, filtering, formatting, and error
  paths.
- Do not make tests depend on the live TicTrens endpoint. Use `httptest` or
  injectable dependencies when fetch behavior needs coverage.
- Run `gofmt` after Go edits.

## Validation

For meaningful Go changes, run:

```sh
go test ./...
go build -o /tmp/tpsp-check ./cmd/tpsp
git diff --check
```

For CLI-surface changes, also run the relevant help/version/copyright commands.
For release or Docker changes, inspect `.goreleaser.yaml`, `Dockerfile`, and the
GitHub Actions workflow together because they share the same binary entrypoint.
