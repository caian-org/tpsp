set shell := ["/bin/bash", "-c"]

BIN := "bin/tpsp"

# --------------------------------------------------------------------------------------------------

_help:
    @just --list

# --------------------------------------------------------------------------------------------------

# build tpsp into bin/tpsp
build:
    @mkdir -p bin
    @go build -o {{ BIN }} ./cmd/tpsp
    @echo "built tpsp"

# run go test ./...
test:
    go test ./...

# run the test suite with the race detector
test-race:
    go test -race ./...

# coverage profile + per-function totals
cover:
    go test -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out | tail -20

# go vet (CI also runs golangci-lint)
lint:
    go vet ./...

# go mod tidy
tidy:
    go mod tidy

# remove build outputs
clean:
    rm -rf bin coverage.out

# build then run bin/tpsp with the given args
run *ARGS:
    @just build
    @{{ BIN }} {{ ARGS }}
