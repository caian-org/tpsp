---
name: tpsp-cli-http
description: CLI, HTTP fetch, service filtering, terminal table output, and JSON output guidance for tpsp. Use when modifying cmd/tpsp/main.go behavior or tests.
---

# TPSP CLI HTTP

Use this skill when changing runtime behavior in `cmd/tpsp/main.go`.

## Public CLI Contract

The command shape is:

```text
tpsp [service] [flags]
```

Supported services are `metro`, `cptm`, `viamobilidade`, and `viaquatro`.
Supported flags are `-j/--json`, `-v/--version`, `--copyright`, and
`-h/--help`. Unknown flags and invalid services should fail with a clear message
on stderr and a non-zero exit.

## HTTP And API Rules

- The current source endpoint is `https://www.tictrens.com.br/helper/line-statuses`.
- Keep a finite HTTP timeout.
- Treat non-200 HTTP responses, JSON decode failures, and `status: false` API
  responses as errors.
- Model external JSON with explicit structs. Do not flatten or discard fields
  before the code that needs them.
- For tests, avoid real network calls. Prefer `httptest` and inject the endpoint
  or HTTP client if fetch behavior is refactored.

## Output Rules

- Table output is for humans and may use ANSI color.
- JSON output is for scripts and should stay stable:

```json
{
  "code": 200,
  "data": [
    {
      "line": "Azul",
      "status": "Opera\u00e7\u00e3o Normal"
    }
  ],
  "message": "success"
}
```

- Preserve the normalized Portuguese status strings used by the current CLI.
- Preserve line-name extraction from the API `line` field unless fixtures prove
  the upstream format changed.
- Do not add a leading blank line to JSON-only output in future refactors unless
  the user intentionally keeps existing output byte-for-byte.

## Test Expectations

When behavior changes, cover at least the closest pure helpers first:

- service validation is case-insensitive;
- filtering with empty service returns all lines;
- filtering by service is case-insensitive;
- status normalization handles plural upstream strings;
- line-name formatting handles expected `prefix - name` input;
- JSON output keeps stable field names.

Add HTTP tests when changing fetch behavior, timeout handling, or decode/error
paths.
