# Repository Guidelines

## Project Structure & Module Organization

`tpsp` is a small Go CLI that prints the current line status for Sao Paulo
public transport services: Metro, CPTM, ViaMobilidade, and ViaQuatro. The
runtime entrypoint is `cmd/tpsp/main.go`. The built binary belongs in `bin/`;
do not edit generated binaries by hand. User-facing documentation is in
`README.md` and `docs/README.md`. Release packaging lives in `.goreleaser.yaml`,
`Dockerfile`, and `.github/workflows/release.yml`.

## Build, Test, and Development Commands

- `just build` builds `bin/tpsp`.
- `go build -o bin/tpsp ./cmd/tpsp` builds the CLI directly.
- `go run ./cmd/tpsp --help` checks the help surface.
- `go run ./cmd/tpsp --version` checks the version surface.
- `go run ./cmd/tpsp --copyright` checks the CC0 copyright text.
- `go test ./...` runs tests and `go vet` checks for the package.
- `go vet ./...` runs vet explicitly.
- `just clean` removes the generated `bin/tpsp` binary.
- `gofmt -w cmd/tpsp/main.go` formats the Go source after edits.

Avoid making tests depend on the live `https://www.tictrens.com.br` endpoint.
For behavior tests, prefer `httptest` or a replaceable HTTP client/URL.

## Coding Style & Naming Conventions

Use standard Go style: `gofmt`, small functions, clear error wrapping, and
explicit structs for external JSON. Keep exported names only when they are part
of a real package API; this CLI currently lives in `package main`, so most names
should stay unexported unless a test or future package split requires otherwise.

Preserve the public CLI contract:

- positional service filters: `metro`, `cptm`, `viamobilidade`, `viaquatro`
- flags: `-j/--json`, `-v/--version`, `--copyright`, `-h/--help`
- JSON shape: `{ "code": 200, "data": [...], "message": "success" }`
- no service argument means all services

Keep terminal table output readable and script output stable. Do not change
normalization of line names or status strings without updating docs and tests.

## Testing Guidelines

Add focused tests next to `cmd/tpsp/main.go` when changing parsing, formatting,
filtering, JSON output, or HTTP error handling. Favor table-driven tests for
pure helpers such as service validation, status normalization, line formatting,
and filtering. For fetch behavior, inject an HTTP server/client rather than
calling the production service.

Before finishing meaningful changes, run:

```sh
gofmt -w cmd/tpsp/main.go
go test ./...
go build -o /tmp/tpsp-check ./cmd/tpsp
git diff --check
```

## Release And Packaging

GoReleaser builds Linux, macOS, and Windows artifacts for amd64 and arm64 from
`./cmd/tpsp`. The GitHub Actions release workflow runs only for `v*` tags. The
Docker image uses a multi-stage build and a `scratch` runtime with CA
certificates, so preserve static binary behavior (`CGO_ENABLED=0`) unless the
packaging is updated with matching evidence.

## Commit & Pull Request Guidelines

Recent history uses short imperative messages, with conventional prefixes for
some changes such as `fix:` and `build:`. Keep commits focused. Prefer
conventional subjects when the scope is obvious, for example `fix: handle API
decode errors` or `build: update release workflow`. Pull requests should include
what changed, why it matters for CLI users, and the exact validation commands
run.

## Agent-Specific Instructions

Respect the disclaimer: this project is not affiliated with the Estado de Sao
Paulo, CPTM, Metro, ViaMobilidade, ViaQuatro, or any government agency. Do not
make wording or UX imply official status.

Do not edit generated output by hand: `bin/tpsp`, future coverage files, and
temporary build artifacts are generated. Do not commit credentials, real API
secrets, or machine-local files.

Skills live under `.codex/skills/`. Codex agents live under `.codex/agents/`;
Claude Code agents live under `.claude/agents/`. Keep `AGENTS.md` authoritative
for shared repository rules.
